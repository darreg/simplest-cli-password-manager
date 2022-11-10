package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

func SelectCommand(
	ctx context.Context,
	cliScript port.CLISelectCommandSupporter,
	commands []string,
) (int, error) {
	var commandIndex int

	err := cliScript.SelectCommand(ctx, commands, &commandIndex)
	if err != nil {
		return 0, err
	}

	return commandIndex, nil
}
