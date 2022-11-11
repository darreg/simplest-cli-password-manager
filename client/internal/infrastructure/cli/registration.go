package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

const (
	MinLoginLength    = 3
	MinPasswordLength = 6
	MaxLength         = 255
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
			Name:   "login",
			Prompt: &survey.Input{Message: "Login"},
			Validate: survey.ComposeValidators(
				survey.Required,
				survey.MinLength(MinLoginLength),
				survey.MaxLength(MaxLength),
			),
		},
		{
			Name:   "password",
			Prompt: &survey.Password{Message: "Password"},
			Validate: survey.ComposeValidators(
				survey.Required,
				survey.MinLength(MinPasswordLength),
				survey.MaxLength(MaxLength),
			),
		},
		{
			Name:   "repeatpassword",
			Prompt: &survey.Password{Message: "Repeat password"},
			Validate: survey.ComposeValidators(
				survey.Required,
				survey.MinLength(MinPasswordLength),
				survey.MaxLength(MaxLength),
			),
		},
	}

	registrationData := &usecase.RegistrationDTO{}
	err := survey.Ask(qs, registrationData)
	if err != nil {
		return nil, err
	}

	return registrationData, nil
}
