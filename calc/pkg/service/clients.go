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

	subpb "github.com/adir-ch/micro-kit/sub/pkg/grpc/pb"
	"google.golang.org/grpc"
)

type ClientFunc func(data interface{}, uri string) (float64, error)

func httpSend(data interface{}, uri string) (float64, error) {
	log.Printf("sending HTTP op request to uri: %s", uri)
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

func grpcSendSub(data interface{}, uri string) (float64, error) {
	log.Printf("sending gRPC op request to uri: %s", uri)

	req := data.([]float64)
	if len(req) < 2 {
		return 0, fmt.Errorf("illegal input data len received: %d", len(req))
	}

	conn, err := grpc.Dial(uri, grpc.WithInsecure())
	if err != nil {
		log.Printf("unable to connect to: %s, err: %s", uri, err.Error())
		return 0, err
	}

	client := subpb.NewSubClient(conn)
	rs, err := client.Sub(context.Background(), &subpb.SubReqest{Left: req[0], Right: req[1]})
	if err != nil {
		return 0, err
	}

	return rs.Result, nil
}
