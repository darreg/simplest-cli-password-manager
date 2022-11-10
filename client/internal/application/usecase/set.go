package usecase

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/model"
	"github.com/alrund/yp-2-project/client/internal/domain/port"
)

type SetEntryDTO struct {
	TypeIndex int
	Name      string
	Metadata  string
	Data      string
}

func Set(
	ctx context.Context,
	client port.GRPCClientSetSupporter,
	cliScript port.CLISetEntrySupporter,
	types []*model.Type,
) error {
	typeNames := make([]string, len(types))
	for i, tp := range types {
		typeNames[i] = tp.Name
	}

	entryDTO := &SetEntryDTO{}
	err := cliScript.SetEntry(ctx, typeNames, entryDTO)
	if err != nil {
		return err
	}

	err = client.SetEntry(ctx, types[entryDTO.TypeIndex].ID, entryDTO.Name, entryDTO.Metadata, []byte(entryDTO.Data))
	if err != nil {
		return err
	}

	return nil
}
