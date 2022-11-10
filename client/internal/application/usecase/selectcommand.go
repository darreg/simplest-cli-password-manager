package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

func SelectCommand(
	ctx context.Context,
	cliScript port.CLISelectCommandSupporter,
	commands map[string]func() (string, error),
) (func() (string, error), error) {
	var selectedCommandName string

	var commandNames []string
	for commandName := range commands {
		commandNames = append(commandNames, commandName)
	}

	err := cliScript.SelectCommand(ctx, commandNames, &selectedCommandName)
	if err != nil {
		return nil, err
	}

	return commands[selectedCommandName], nil
}
