package component

import (
	"context"
	"log"
	"net"

	v1 "github.com/AkiOuma/biz-groups/api/goapi/janus-biz-groups/service/v1"
	"github.com/AkiOuma/biz-groups/internal/infrastructure/config"
	"github.com/AkiOuma/biz-groups/internal/interface/transport"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener net.Listener
	server   *grpc.Server
}

func NewGrpcServer(
	trp *transport.Transport,
	c *config.Bootstrap,
) *GrpcServer {
	lis, err := net.Listen("tcp", c.GetServer().GetGrpc().GetAddr())
	if err != nil {
		log.Fatalf("[Error]infrastructure.component.grpc.NewGrpcServer: create listener failed because: [%v]", err)
	}
	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_recovery.UnaryServerInterceptor([]grpc_recovery.Option{
			grpc_recovery.WithRecoveryHandler(trp.RecoveryInterceptor),
		}...),
		grpc.UnaryServerInterceptor(trp.ErrInterceptor),
	))
	v1.RegisterTransportServer(grpcServer, trp)
	return &GrpcServer{
		listener: lis,
		server:   grpcServer,
	}
}

func (s *GrpcServer) Serve(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		s.server.Stop()
	}()
	log.Printf("Start Grpc Server on: %s", s.listener.Addr())
	return s.server.Serve(s.listener)
}
