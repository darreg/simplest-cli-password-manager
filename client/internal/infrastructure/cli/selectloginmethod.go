package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
)

// SelectLoginMethod displays a list of login methods.
func (c *Collection) SelectLoginMethod(ctx context.Context, options []string) (string, error) {
	prompt := &survey.Select{
		Message: "Choose a login method:",
		Options: options,
		Default: "Login",
	}

	var loginMethodIndex string
	err := survey.AskOne(prompt, &loginMethodIndex)
	if err != nil {
		return "", err
	}

	return loginMethodIndex, nil
}
