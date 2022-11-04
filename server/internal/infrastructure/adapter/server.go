package adapter

import (
	"fmt"
	"net"

	"github.com/alrund/yp-2-project/server/internal/infrastructure/handler"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	runAddress, certFile, keyFile string
	grpcServer                    *grpc.Server
}

func NewServer(runAddress, certFile, keyFile string) *Server {
	return &Server{
		runAddress: runAddress,
		certFile:   certFile,
		keyFile:    keyFile,
	}
}

func (s *Server) Serve(handlerCollection any) error {
	collection, ok := handlerCollection.(*handler.Collection)
	if !ok {
		return fmt.Errorf("incorrect handlerCollection")
	}

	creds, err := credentials.NewServerTLSFromFile(s.certFile, s.keyFile)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer(grpc.Creds(creds)) // grpc.UnaryInterceptor(collection.AuthInterceptor(a.Encryptor))

	proto.RegisterAppServer(s.grpcServer, collection)

	listen, err := net.Listen("tcp", s.runAddress)
	if err != nil {
		return err
	}

	err = s.grpcServer.Serve(listen)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Shutdown() error {
	if s.grpcServer == nil {
		return fmt.Errorf("the server is not running")
	}
	s.grpcServer.GracefulStop()
	return nil
}
