package cli

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
)

// ListOfEntries displays a list of entries.
func (c *Collection) ListOfEntries(ctx context.Context, entries []string) (int, error) {
	prompt := &survey.Select{
		Message: "Choose an entry:",
		Options: entries,
	}

	var entryIndex int
	err := survey.AskOne(prompt, &entryIndex)
	if err != nil {
		return 0, err
	}

	return entryIndex, nil
}
