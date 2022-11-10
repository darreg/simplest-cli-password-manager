package usecase

import (
	"context"
	"fmt"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

func Greetings(
	ctx context.Context,
	client port.GRPCClientUserGetter,
) (string, error) {
	user, err := client.GetUser(ctx)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Greetings, %s!", user.Name), nil
}
