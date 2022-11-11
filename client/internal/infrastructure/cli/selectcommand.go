package cli

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

// SelectCommand displays a list of commands.
func (c *Collection) SelectCommand(ctx context.Context, options []string, data any) error {
	commandIndex, ok := data.(*string)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	fmt.Println("")
	prompt := &survey.Select{
		Message: "Select a command:",
		Options: options,
	}

	err := survey.AskOne(prompt, commandIndex)
	if err != nil {
		return err
	}

	return nil
}
