package usecase

import (
	"context"
	"sort"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

// SelectLoginMethod processes a list of login methods.
func SelectLoginMethod(
	ctx context.Context,
	cliScript port.CLISelectLoginMethodSupporter,
	loginMethods map[string]func() (string, error),
) (func() (string, error), error) {
	names := make([]string, len(loginMethods))
	var i int
	for methodName := range loginMethods {
		names[i] = methodName
		i++
	}

	sort.Strings(names)

	loginMethodName, err := cliScript.SelectLoginMethod(ctx, names)
	if err != nil {
		return nil, err
	}

	return loginMethods[loginMethodName], nil
}
