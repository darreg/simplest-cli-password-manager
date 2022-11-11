package cli

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// SelectCommand displays a list of commands.
func (c *Collection) SelectCommand(ctx context.Context, options []string) (string, error) {
	fmt.Println("")
	prompt := &survey.Select{
		Message: "Select a command:",
		Options: options,
	}

	var commandIndex string
	err := survey.AskOne(prompt, &commandIndex)
	if err != nil {
		return "", err
	}

	return commandIndex, nil
}
