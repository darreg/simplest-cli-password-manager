package usecase

import (
	"context"
	"sort"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

func SelectCommand(
	ctx context.Context,
	cliScript port.CLISelectCommandSupporter,
	commands map[string]func() (string, error),
) (func() (string, error), error) {
	var selectedCommandName string

	commandNames := make([]string, len(commands))
	var i int
	for commandName := range commands {
		commandNames[i] = commandName
		i++
	}

	sort.Strings(commandNames)

	err := cliScript.SelectCommand(ctx, commandNames, &selectedCommandName)
	if err != nil {
		return nil, err
	}

	return commands[selectedCommandName], nil
}
