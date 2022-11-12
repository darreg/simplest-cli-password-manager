package client

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	"github.com/alrund/yp-2-project/client/internal/application/usecase"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	bufSize = 1024 * 1024
)

var (
	testTime            = time.Now()
	testSessionKey      = "test key"
	testResponseEntries = []*proto.GetAllEntriesResponse_Entry{
		{
			EntryId:   "entry id 1",
			TypeId:    "type id 1",
			Name:      "entry name 1",
			CreatedAt: timestamppb.New(testTime),
			UpdatedAt: timestamppb.New(testTime),
		},
		{
			EntryId:   "entry id 2",
			TypeId:    "type id 2",
			Name:      "entry name 2",
			CreatedAt: timestamppb.New(testTime),
			UpdatedAt: timestamppb.New(testTime),
		},
	}
	testResponseTypes = []*proto.GetAllTypesResponse_Type{
		{
			TypeId:   "type id 1",
			Name:     "type name 1",
			IsBinary: false,
		},
		{
			TypeId:   "type id 2",
			Name:     "type name 2",
			IsBinary: false,
		},
		{
			TypeId:   "type id 3",
			Name:     "type name 3",
			IsBinary: false,
		},
	}
	testResponseUser = proto.GetUserResponse{
		UserId: "user id 1",
		Name:   "user name 1",
		Login:  "user login 1",
	}
)

type mockAppServer struct {
	proto.UnimplementedAppServer
}

func (c *mockAppServer) GetAllEntries(
	ctx context.Context,
	in *proto.GetAllEntriesRequest,
) (*proto.GetAllEntriesResponse, error) {
	return &proto.GetAllEntriesResponse{Entries: testResponseEntries}, nil
}

func (c *mockAppServer) GetAllTypes(
	ctx context.Context,
	in *proto.GetAllTypesRequest,
) (*proto.GetAllTypesResponse, error) {
	return &proto.GetAllTypesResponse{Types: testResponseTypes}, nil
}

func (c *mockAppServer) GetEntry(
	ctx context.Context,
	in *proto.GetEntryRequest,
) (*proto.GetEntryResponse, error) {
	return &proto.GetEntryResponse{
		EntryId:   testResponseEntries[0].EntryId,
		TypeId:    testResponseEntries[0].TypeId,
		Name:      testResponseEntries[0].Name,
		CreatedAt: testResponseEntries[0].CreatedAt,
		UpdatedAt: testResponseEntries[0].UpdatedAt,
	}, nil
}

func (c *mockAppServer) GetUser(
	ctx context.Context,
	in *proto.GetUserRequest,
) (*proto.GetUserResponse, error) {
	return &testResponseUser, nil
}

func (c *mockAppServer) Login(
	ctx context.Context,
	in *proto.LoginRequest,
) (*proto.LoginResponse, error) {
	return &proto.LoginResponse{EncryptedSessionKey: testSessionKey}, nil
}

func (c *mockAppServer) Registration(
	ctx context.Context,
	in *proto.RegistrationRequest,
) (*proto.RegistrationResponse, error) {
	return &proto.RegistrationResponse{EncryptedSessionKey: testSessionKey}, nil
}

func (c *mockAppServer) SetEntry(
	ctx context.Context,
	in *proto.SetEntryRequest,
) (*proto.SetEntryResponse, error) {
	return &proto.SetEntryResponse{}, nil
}

func testDialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(bufSize)
	server := grpc.NewServer()
	proto.RegisterAppServer(server, &mockAppServer{})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestSetGRPCClient(t *testing.T) {
	client := &Client{}

	conn, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(testDialer()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	testAppClient := proto.NewAppClient(conn)

	t.Run("success", func(t *testing.T) {
		err := client.SetGRPCClient(testAppClient)
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		err := client.SetGRPCClient(nil)
		assert.ErrorIs(t, err, usecase.ErrInvalidArgument)
	})
}

func TestSetSessionKey(t *testing.T) {
	client := &Client{}

	t.Run("success", func(t *testing.T) {
		err := client.SetSessionKey("session key")
		assert.NoError(t, err)
	})

	t.Run("fail", func(t *testing.T) {
		err := client.SetSessionKey("")
		assert.ErrorIs(t, err, usecase.ErrInvalidArgument)
	})
}

func TestIsEmptySessionKey(t *testing.T) {
	t.Run("false", func(t *testing.T) {
		client := &Client{sessionKey: "test key"}
		result := client.IsEmptySessionKey()
		assert.False(t, result)
	})

	t.Run("true", func(t *testing.T) {
		client := &Client{}
		result := client.IsEmptySessionKey()
		assert.True(t, result)
	})
}
