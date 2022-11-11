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

// Set adds an entry.
func Set(
	ctx context.Context,
	client port.GRPCClientSetSupporter,
	cliScript port.CLISetEntrySupporter,
	types []*model.Type,
) (string, error) {
	typeNames := make([]string, len(types))
	for i, tp := range types {
		typeNames[i] = tp.Name
	}

	dto, err := cliScript.SetEntry(ctx, typeNames)
	if err != nil {
		return "", err
	}

	entryDTO, ok := dto.(*SetEntryDTO)
	if !ok {
		return "", ErrInternalError
	}

	err = client.SetEntry(ctx, types[entryDTO.TypeIndex].ID, entryDTO.Name, entryDTO.Metadata, []byte(entryDTO.Data))
	if err != nil {
		return "", err
	}

	return "", nil
}
