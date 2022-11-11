package adapter

import (
	"fmt"
	"net"

	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/handler"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/interceptor"
	"github.com/alrund/yp-2-project/server/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Server GRPC server.
type Server struct {
	sessionLifeTime, runAddress string
	certFile, keyFile           string
	decryptor                   port.Decryptor
	grpcServer                  *grpc.Server
	sessionRepository           port.SessionRepository
	logger                      port.Logger
}

func NewServer(
	sessionLifeTime, runAddress string,
	certFile, keyFile string,
	decryptor port.Decryptor,
	sessionRepository port.SessionRepository,
	logger port.Logger,
) *Server {
	return &Server{
		sessionLifeTime:   sessionLifeTime,
		runAddress:        runAddress,
		certFile:          certFile,
		keyFile:           keyFile,
		decryptor:         decryptor,
		sessionRepository: sessionRepository,
		logger:            logger,
	}
}

// Serve starts the GRPC server.
func (s *Server) Serve(handlerCollection any) error {
	collection, ok := handlerCollection.(*handler.Collection)
	if !ok {
		return fmt.Errorf("incorrect handlerCollection")
	}

	creds, err := credentials.NewServerTLSFromFile(s.certFile, s.keyFile)
	if err != nil {
		return err
	}

	s.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			interceptor.Auth(s.sessionLifeTime, s.sessionRepository, s.decryptor),
			interceptor.RequestLog(s.logger),
		),
	)

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

// Shutdown stops the GRPC server.
func (s *Server) Shutdown() error {
	if s.grpcServer == nil {
		return fmt.Errorf("the server is not running")
	}
	s.grpcServer.GracefulStop()
	return nil
}
