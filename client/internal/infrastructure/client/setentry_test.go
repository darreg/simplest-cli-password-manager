package client

import (
	"context"
	"log"
	"testing"

	"github.com/alrund/yp-2-project/client/pkg/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSetEntry(t *testing.T) {
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
		client := &Client{
			grpcClient: testAppClient,
			sessionKey: "test key",
		}

		err := client.SetEntry(
			context.Background(),
			testResponseEntries[0].TypeId,
			testResponseEntries[0].Name,
			"",
			[]byte(""),
		)
		require.NoError(t, err)
	})

	t.Run("fail with empty sessionKey", func(t *testing.T) {
		client := &Client{
			grpcClient: testAppClient,
			sessionKey: "",
		}

		err := client.SetEntry(
			context.Background(),
			testResponseEntries[0].TypeId,
			testResponseEntries[0].Name,
			"",
			[]byte(""),
		)
		require.ErrorIs(t, err, ErrSessionKey)
	})

	t.Run("fail with empty grpcClient", func(t *testing.T) {
		client := &Client{
			sessionKey: "test key",
		}

		err := client.SetEntry(
			context.Background(),
			testResponseEntries[0].TypeId,
			testResponseEntries[0].Name,
			"",
			[]byte(""),
		)
		require.ErrorIs(t, err, ErrGRPCClient)
	})
}
