package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

// Registration displays the registration form.
func (c *Collection) Registration(ctx context.Context) (any, error) {
	qs := []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Name"},
			Validate: survey.Required,
		},
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

	registrationData := &usecase.RegistrationDTO{}
	err := survey.Ask(qs, registrationData)
	if err != nil {
		return nil, err
	}

	return registrationData, nil
}
