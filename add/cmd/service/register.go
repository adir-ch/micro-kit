package service

import (
	"context"

	log "github.com/go-kit/kit/log"
	sdetcd "github.com/go-kit/kit/sd/etcd"
)

func etcdRegister(logger log.Logger) (*sdetcd.Registrar, error) {
	var (
		etcdServer = "http://etcd:2379"
		prefix     = "/services/add/"
		instance   = "add:8081"
		key        = prefix + instance
	)

	client, err := sdetcd.NewClient(context.Background(), []string{etcdServer}, sdetcd.ClientOptions{})
	if err != nil {
		return nil, err
	}

	registrar := sdetcd.NewRegistrar(client, sdetcd.Service{
		Key:   key,
		Value: instance,
	}, logger)

	registrar.Register()
	logger.Log("service add is registered in etcd")
	return registrar, nil
}
