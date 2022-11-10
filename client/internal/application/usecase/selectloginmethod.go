package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

func SelectLoginMethod(
	ctx context.Context,
	cliScript port.CLISelectLoginMethodSupporter,
	loginMethods []string,
) (int, error) {
	var loginMethodIndex int

	err := cliScript.SelectLoginMethod(ctx, loginMethods, &loginMethodIndex)
	if err != nil {
		return 0, err
	}

	return loginMethodIndex, nil
}
