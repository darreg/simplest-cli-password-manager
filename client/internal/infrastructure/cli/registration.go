package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

func (c *Collection) Registration(ctx context.Context, data any) error {
	registrationData, ok := data.(*usecase.RegistrationDTO)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	qs := []*survey.Question{
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
		{
			Name:     "repeatpassword",
			Prompt:   &survey.Password{Message: "Repeat password"},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(qs, registrationData)
	if err != nil {
		return err
	}

	return nil
}
