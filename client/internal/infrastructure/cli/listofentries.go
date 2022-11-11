package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/alrund/yp-2-project/client/internal/application/usecase"
)

// ListOfEntries displays a list of entries.
func (c *Collection) ListOfEntries(ctx context.Context, entries []string, data any) error {
	entryIndex, ok := data.(*int)
	if !ok {
		return usecase.ErrInvalidArgument
	}

	prompt := &survey.Select{
		Message: "Choose an entry:",
		Options: entries,
	}

	err := survey.AskOne(prompt, entryIndex)
	if err != nil {
		return err
	}

	return nil
}
