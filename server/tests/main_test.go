//go:build integration

package tests

import (
	"testing"

	"github.com/alrund/yp-1-project/internal/infrastructure/adapter"
	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	testSuite := NewIntegrationTestSuite(
		adapter.NewLogger(),
		"../docker-compose.test.yml",
		map[string]string{
			"POSTGRES_PORT": "5433",
		},
	)
	suite.Run(t, testSuite)
}
