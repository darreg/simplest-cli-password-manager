package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

// Login displays the authorization form.
func (c *Collection) Login(ctx context.Context) (any, error) {
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
	}

	credential := &usecase.LoginDTO{}
	err := survey.Ask(qs, credential)
	if err != nil {
		return nil, err
	}

	return credential, nil
}
