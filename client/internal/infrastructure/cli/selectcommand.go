package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

func SelectCommand(ctx context.Context, options []string, data any) error {
	CommandIndex, ok := data.(*int)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	prompt := &survey.Select{
		Message: "Select a command:",
		Options: options,
	}

	err := survey.AskOne(prompt, CommandIndex)
	if err != nil {
		return err
	}

	return nil
}
