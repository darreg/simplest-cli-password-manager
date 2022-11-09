package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

func SelectLoginMethod(ctx context.Context, options []string, data any) error {
	loginMethodIndex, ok := data.(*int)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	prompt := &survey.Select{
		Message: "Your login method:",
		Options: options,
		Default: "Login",
	}

	err := survey.AskOne(prompt, loginMethodIndex)
	if err != nil {
		return err
	}

	return nil
}
