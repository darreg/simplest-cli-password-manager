package usecase

import (
	"context"
	"sort"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

// SelectCommand processes a list of commands.
func SelectCommand(
	ctx context.Context,
	cliScript port.CLISelectCommandSupporter,
	commands map[string]func() (string, error),
) (func() (string, error), error) {
	commandNames := make([]string, 0, len(commands))
	for commandName := range commands {
		commandNames = append(commandNames, commandName)
	}

	sort.Strings(commandNames)

	selectedCommandName, err := cliScript.SelectCommand(ctx, commandNames)
	if err != nil {
		return nil, err
	}

	return commands[selectedCommandName], nil
}
