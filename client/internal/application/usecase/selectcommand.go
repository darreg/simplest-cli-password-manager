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
	commandNames := make([]string, len(commands))
	var i int
	for commandName := range commands {
		commandNames[i] = commandName
		i++
	}

	sort.Strings(commandNames)

	selectedCommandName, err := cliScript.SelectCommand(ctx, commandNames)
	if err != nil {
		return nil, err
	}

	return commands[selectedCommandName], nil
}
