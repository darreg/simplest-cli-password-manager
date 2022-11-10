package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

func SelectLoginMethod(
	ctx context.Context,
	cliScript port.CLISelectLoginMethodSupporter,
	loginMethods map[string]func() (string, error),
) (func() (string, error), error) {
	var loginMethodName string

	var names []string
	for methodName := range loginMethods {
		names = append(names, methodName)
	}

	err := cliScript.SelectLoginMethod(ctx, names, &loginMethodName)
	if err != nil {
		return nil, err
	}

	return loginMethods[loginMethodName], nil
}
