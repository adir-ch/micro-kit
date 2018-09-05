package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	sdetcd "github.com/go-kit/kit/sd/etcd"
)

var registry = "http://etcd:2379" // TODO: take from config

var ops = map[string]string{ // TODO: take from config
	"+": "services/add/",
	"-": "services/sub/",
	"*": "services/mul/",
	"/": "services/div/",
}

type svcreq struct {
	Numbers []float64 `json:"numbers,omitempty"`
}

type svcres struct {
	Rs  float64 `json:"rs"`
	Err error   `json:"err"`
}

func callOp(op string, numbers []float64) (float64, error) {
	svcName, found := ops[op]
	if found == false {
		return 0, fmt.Errorf("unable to find service for op: %s", op)
	}

	srvURI, err := discover(svcName)
	if err != nil {
		return 0, err
	}

	return send(svcreq{numbers}, fmt.Sprintf("http://%s/calc", srvURI))
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

func send(data interface{}, uri string) (float64, error) {
	log.Printf("sending op request to uri: %s", uri)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

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

	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("response status error from %s: %d", uri, res.StatusCode)
	}

	var result svcres
	if info, err := readAndParseJSON(res.Body, &result); err != nil {
		return 0, fmt.Errorf("unable to decode response from %s: %s (%s)", uri, info, err)
	}

	return result.Rs, result.Err
}

func readAndParseJSON(body io.ReadCloser, dest interface{}) (string, error) {
	if data, err := ioutil.ReadAll(body); err != nil {
		return "unable to read body structure", err
	} else if err = json.Unmarshal(data, dest); err != nil {
		return "unable to parse body structure", err
	}
	return "", nil
}
