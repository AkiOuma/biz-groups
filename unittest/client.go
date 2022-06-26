package unittest

import (
	"context"
	"log"

	v1 "github.com/AkiOuma/biz-groups/api/goapi/janus-biz-groups/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client v1.TransportClient
	ctx    = context.Background()
	addr   = "127.0.0.1:8080"
)

func Client() v1.TransportClient {
	if client != nil {
		return client
	}
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
	}
	client := v1.NewTransportClient(conn)
	return client
}
