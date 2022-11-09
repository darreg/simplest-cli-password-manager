package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

func (c *Collection) SetEntry(ctx context.Context, types []string, data any) error {
	setEntryDTO, ok := data.(*usecase.SetEntryDTO)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	qs := []*survey.Question{
		{
			Name: "typeindex",
			Prompt: &survey.Select{
				Message: "Choose a type:",
				Options: types,
			},
			Validate: survey.Required,
		},
		{
			Name: "name",
			Prompt: &survey.Input{
				Message: "Name",
			},
			Validate: survey.Required,
		},
		{
			Name: "metadata",
			Prompt: &survey.Multiline{
				Message: "Metadata",
			},
			Validate: survey.Required,
		},
		{
			Name: "data",
			Prompt: &survey.Multiline{
				Message: "Data",
			},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(qs, setEntryDTO)
	if err != nil {
		return err
	}

	return nil
}
