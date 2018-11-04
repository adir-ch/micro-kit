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

	mulep "github.com/adir-ch/micro-kit/mul/pkg/endpoint"
	subpb "github.com/adir-ch/micro-kit/sub/pkg/grpc/pb"
	nats "github.com/nats-io/go-nats"
	"google.golang.org/grpc"
)

type ClientFunc func(data interface{}, svcURL string) (float64, error)

func httpSend(data interface{}, svcURL string) (float64, error) {
	log.Printf("sending HTTP op request to svcURL: %s", svcURL)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	url := fmt.Sprintf("http://%s/calc", svcURL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(dataBytes))

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
		return 0, fmt.Errorf("response status error from %s: %d", svcURL, res.StatusCode)
	}

	var result svcres
	if info, err := readAndParseJSON(res.Body, &result); err != nil {
		return 0, fmt.Errorf("unable to decode response from %s: %s (%s)", svcURL, info, err)
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

func grpcSendSub(data interface{}, svcURL string) (float64, error) {
	log.Printf("sending gRPC op request to svcURL: %s", svcURL)

	req := data.(svcreq)
	if len(req.Numbers) < 2 {
		return 0, fmt.Errorf("illegal input data len received: %d", len(req.Numbers))
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s", svcURL), grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to: %s, err: %s", svcURL, err.Error())
		return 0, err
	}

	client := subpb.NewSubClient(conn)
	rs, err := client.Sub(context.Background(), &subpb.SubRequest{Left: req.Numbers[0], Right: req.Numbers[1]})
	if err != nil {
		return 0, err
	}

	return rs.Result, nil
}

func reqRepSendMul(data interface{}, svcURL string) (float64, error) {
	log.Printf("sending NATS req-rep op request to broker on: %s", nats.DefaultURL)

	req := data.(svcreq)
	if len(req.Numbers) < 2 {
		return 0, fmt.Errorf("illegal input data len received: %d", len(req.Numbers))
	}

	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		return 0, fmt.Errorf("unable to connect to NATS broker")
	}

	r, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("unable to marshal data into byte: %d", len(req.Numbers))
	}

	defer nc.Close()
	rs, err := nc.Request("mul", r, 500*time.Millisecond)
	if err != nil || nc.LastError() != nil {
		return 0, fmt.Errorf("Error in NATS Request: %s (last: %v)", err, nc.LastError())
	}

	var result mulep.MulResponse
	if err := json.Unmarshal(rs.Data, &result); err != nil {
		return 0, fmt.Errorf("unable to unmarshal service response: %s", err)
	}
	return result.Rs, result.Err
}
