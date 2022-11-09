package usecase

import (
	"context"
)

type Credential struct {
	Login    string
	Password string
}

func Login(
	ctx context.Context,
	cli func(ctx context.Context, data any) error,
) (*Credential, error) {
	credential := &Credential{}

	err := cli(ctx, credential)
	if err != nil {
		return nil, err
	}

	return credential, nil
}
