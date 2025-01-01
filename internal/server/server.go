package server

import (
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	"review-service/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegister)

func NewRegister(conf *conf.Registry) registry.Registrar {
	config := api.DefaultConfig()
	config.Address = conf.Consul.Address
	config.Scheme = conf.Consul.Scheme
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	reg := consul.New(client, consul.WithHealthCheck(true))
	return reg
}
