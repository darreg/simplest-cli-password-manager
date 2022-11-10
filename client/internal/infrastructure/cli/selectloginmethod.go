package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

func (c *Collection) SelectLoginMethod(ctx context.Context, options []string, data any) error {
	loginMethodIndex, ok := data.(*string)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	prompt := &survey.Select{
		Message: "Choose a login method:",
		Options: options,
		Default: "Login",
	}

	err := survey.AskOne(prompt, loginMethodIndex)
	if err != nil {
		return err
	}

	return nil
}
