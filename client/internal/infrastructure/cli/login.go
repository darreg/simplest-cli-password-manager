package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

func Login(ctx context.Context, data any) error {
	credential, ok := data.(*usecase.Credential)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	var qs = []*survey.Question{
		{
			Name:     "login",
			Prompt:   &survey.Input{Message: "Login"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Password"},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(qs, credential)
	if err != nil {
		return err
	}

	return nil
}
