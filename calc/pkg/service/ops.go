package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	sdetcd "github.com/go-kit/kit/sd/etcd"
)

var registry = "http://etcd:2379" // TODO: take from config

var ops = map[string]string{ // TODO: take from config
	"+": "service/add/",
	"-": "service/sub/",
	"*": "service/mul/",
	"/": "service/div/",
}

func callOp(op string, numbers []float64) (float64, error) {
	svcName, found := ops[op]
	if found == false {
		return fmt.Errorf("unable to find service for op: %s", op)
	}

	srvURI, err := discover(svcName)
	if err != nil {
		return 0, err
	}

	return send(numbers, srvURI)
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

	return entries[0], nil // TODO: implement some round robin or something similar
}

func send(data interface{}, uri string) (float64, error) {
	dataBytes := json.Marshal(data)
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(dataBytes))

	if err != nil {
		return 0, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	res, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	logger.Log("response status code from %s: %d", name, res.StatusCode)
	return 0, nil
}
