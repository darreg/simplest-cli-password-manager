package client

import (
	"context"
	"log"
	"testing"

	"github.com/alrund/yp-2-project/client/pkg/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGetAllTypes(t *testing.T) {
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

		entries, err := client.GetAllTypes(context.Background())
		require.NoError(t, err)
		assert.Equal(t, testResponseTypes[0].TypeId, entries[0].ID)
	})

	t.Run("fail with empty grpcClient", func(t *testing.T) {
		client := &Client{
			sessionKey: "test key",
		}

		_, err := client.GetAllTypes(context.Background())
		require.ErrorIs(t, err, ErrGRPCClient)
	})
}
