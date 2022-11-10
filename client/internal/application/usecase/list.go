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
) (string, error) {
	entries, err := client.GetAllEntries(ctx)
	if err != nil {
		return "", err
	}

	entryNames := make([]string, len(entries))
	for i, entry := range entries {
		entryNames[i] = entry.Name
	}

	var entryIndex int
	err = cliScript.ListOfEntries(ctx, entryNames, &entryIndex)
	if err != nil {
		return "", err
	}

	fullEntry, err := client.GetEntry(ctx, entries[entryIndex].ID)
	if err != nil {
		return "", err
	}

	var tpName string
	for _, tp := range types {
		if tp.ID == fullEntry.TypeID {
			tpName = tp.Name
		}
	}

	return fmt.Sprintf(
		"ID:\t%s\nName:\t%s\nType:\t%s\nMetadata:\n%s\nData:\n%s",
		fullEntry.ID,
		fullEntry.Name,
		tpName,
		fullEntry.Metadata,
		fullEntry.Data,
	), nil
}
