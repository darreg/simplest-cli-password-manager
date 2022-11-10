package usecase

import (
	"context"
	"fmt"

	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

func List(
	ctx context.Context,
	client port.GRPCClientListSupporter,
	cliScript port.CLIListOfEntriesSupporter,
	types []*model.Type,
) error {
	entries, err := client.GetAllEntries(ctx)
	if err != nil {
		return err
	}

	entryNames := make([]string, len(entries))
	for i, entry := range entries {
		entryNames[i] = entry.Name
	}

	var entryIndex int
	err = cliScript.ListOfEntries(ctx, entryNames, &entryIndex)
	if err != nil {
		return err
	}

	fullEntry, err := client.GetEntry(ctx, entries[entryIndex].ID)
	if err != nil {
		return err
	}

	fmt.Println(fullEntry)

	return nil
}
