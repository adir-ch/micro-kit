package service

import (
	"context"
	"fmt"
	"log"

	sdetcd "github.com/go-kit/kit/sd/etcd"
)

type OpService struct {
	Name   string
	Client ClientFunc
}

var registry = "http://etcd:2379" // TODO: take from config

var ops = map[string]OpService{ // TODO: take from config
	"+": OpService{"services/add/", httpSend},
	"-": OpService{"services/sub/", grpcSendSub},
	"*": OpService{"services/mul/", httpSend},
	"/": OpService{"services/div/", httpSend},
}

type svcreq struct {
	Numbers []float64 `json:"numbers,omitempty"`
}

type svcres struct {
	Rs  float64 `json:"rs"`
	Err error   `json:"err"`
}

func callOp(op string, numbers []float64) (float64, error) {
	svc, found := ops[op]
	if found == false {
		return 0, fmt.Errorf("unable to find service for op: %s", op)
	}

	srvURI, err := discover(svc.Name)
	if err != nil {
		return 0, err
	}

	return svc.Client(svcreq{numbers}, srvURI)
}

func discover(name string) (string, error) {
	client, err := sdetcd.NewClient(context.Background(), []string{registry}, sdetcd.ClientOptions{})
	if err != nil {
		return "", fmt.Errorf("unable to connect to registry: %s", err.Error())
	}

	entries, err := client.GetEntries(name)
	if err != nil || len(entries) == 0 {
		return "", fmt.Errorf("unable to discover %s, err: %s", name, err.Error())
	}

	log.Printf("discovered %s: %v", name, entries)
	return entries[0], nil // TODO: implement some round robin or something similar
}
