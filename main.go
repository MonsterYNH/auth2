package main

import (
	"context"

	"github.com/MonsterYNH/api/v1/health"
	"github.com/MonsterYNH/api/v1/oauth2"
	"github.com/MonsterYNH/athena/config"
	"github.com/MonsterYNH/athena/server"
	"github.com/MonsterYNH/auth2/service"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	grpcServer, mux := server.New(nil, nil)

	ctx := context.Background()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	health.RegisterHealthServiceServer(grpcServer, new(service.HealthService))
	health.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)

	oauth2.RegisterAuth2SerivceServer(grpcServer, new(service.AuthService))
	oauth2.RegisterAuth2SerivceHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)

	engine, err := server.NewServer(grpcServer, mux)
	if err != nil {
		panic(err)
	}

	if err := engine.Run(func(sc *config.ServiceConfig) error {
		sc.EnableHTTP = conf.ServiceConfig.EnableHTTP
		sc.Host = conf.ServiceConfig.Host
		sc.Port = conf.ServiceConfig.Port
		sc.ServiceName = conf.ServiceConfig.ServiceName
		return nil
	}); err != nil {
		panic(err)
	}
}
