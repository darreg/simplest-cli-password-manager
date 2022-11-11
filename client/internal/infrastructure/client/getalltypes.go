package client

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/pkg/proto"
)

// GetAllTypes calls the same GRPC method.
func (c *Client) GetAllTypes(ctx context.Context) ([]*model.Type, error) {
	if c.grpcClient == nil {
		return nil, ErrGRPCClient
	}

	if c.types == nil {
		response, err := c.grpcClient.GetAllTypes(ctx, &proto.GetAllTypesRequest{})
		if err != nil {
			return nil, err
		}

		types := make([]*model.Type, len(response.Types))
		for i, tp := range response.Types {
			types[i] = &model.Type{
				ID:       tp.TypeId,
				Name:     tp.Name,
				IsBinary: tp.IsBinary,
			}
		}

		c.types = types
	}

	return c.types, nil
}
