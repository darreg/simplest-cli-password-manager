//go:build unit

package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	var hasher = NewHasher()

	tests := []struct {
		name string
		data string
		want string
	}{
		{
			name: "success",
			data: "раз два три",
			want: "9329adc004377201acd27247a0915d04be7214604ef02947f6752855c4881917",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashed := hasher.Hash(tt.data)
			assert.NotEqual(t, hashed, tt.data)
			assert.Equal(t, tt.want, hashed)
		})
	}
}
