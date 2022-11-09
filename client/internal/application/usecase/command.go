package usecase

import (
	"context"
	"fmt"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

const (
	ListIndex = iota
	SetIndex
)

type SetEntryDTO struct {
	TypeIndex int
	Name      string
	Metadata  string
	Data      string
}

func Command(
	ctx context.Context,
	client port.GRPCClientSupporter,
	cliScript port.CLICommandSupporter,
) error {
	var (
		commands     = []string{ListIndex: "List", SetIndex: "Set"}
		commandIndex int
	)

	err := cliScript.SelectCommand(ctx, commands, &commandIndex)
	if err != nil {
		return err
	}

	switch commandIndex {
	case ListIndex:
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

	case SetIndex:
		types, err := client.GetAllTypes(ctx)
		if err != nil {
			return err
		}

		typeNames := make([]string, len(types))
		for i, tp := range types {
			typeNames[i] = tp.Name
		}

		entryDTO := &SetEntryDTO{}
		err = cliScript.SetEntry(ctx, typeNames, entryDTO)
		if err != nil {
			return err
		}

		err = client.SetEntry(ctx, types[entryDTO.TypeIndex].ID, entryDTO.Name, entryDTO.Metadata, []byte(entryDTO.Data))
		if err != nil {
			return err
		}
	}

	return nil
}
