//go:build integration

package tests

import (
	"context"
	"fmt"
	"github.com/alrund/yp-2-project/server/internal/application/app"
	"github.com/alrund/yp-2-project/server/internal/domain/port"
	"github.com/alrund/yp-2-project/server/internal/infrastructure/builder"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/wait"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"time"
)

type request struct {
	method             string
	target             string
	sessionCookieName  string
	sessionCookieValue string
	body               string
	contentType        string
}
type want struct {
	code        int
	response    string
	contentType string
}

type IntegrationTestSuite struct {
	suite.Suite
	logger            port.Logger
	dockerComposeFile string
	identifier        string
	env               map[string]string
	app               *app.App
}

func NewIntegrationTestSuite(
	logger port.Logger,
	dockerComposeFile string,
	env map[string]string,
) suite.TestingSuite {
	ts := new(IntegrationTestSuite)
	ts.logger = logger
	ts.dockerComposeFile = dockerComposeFile
	ts.env = env
	return ts
}

func (s *IntegrationTestSuite) BuildApp() {
	a, err := builder.Builder(&app.Config{
		Debug:        false,
		RunAddress:   "localhost:3000",
		DatabaseURI:  "postgres://dev:dev@localhost:" + s.env["POSTGRES_PORT"] + "/dev?sslmode=disable",
		CipherPass:   "J53RPX6",
		MigrationDir: "../migrations",
	}, s.logger)
	if err != nil {
		s.TearDownSuite()
		s.logger.Fatal(err)
	}
	s.app = a
}

func (s *IntegrationTestSuite) SetupSuite() {
	postgresPort, err := strconv.Atoi(s.env["POSTGRES_PORT"])
	if err != nil {
		s.logger.Fatal(fmt.Errorf("invalid env value: %s - %v", "POSTGRES_PORT", err))
	}

	composeFilePaths := []string{s.dockerComposeFile}
	s.identifier = strings.ToLower(uuid.New().String())
	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, s.identifier)
	err = compose.
		WithCommand([]string{"up", "-d"}).
		WithEnv(s.env).
		WithExposedService(
			"postgres_1",
			postgresPort,
			wait.NewLogStrategy("database system is ready to accept connections").
				WithStartupTimeout(10*time.Second).
				WithOccurrence(2),
		).
		Invoke().
		Error
	if err != nil {
		s.logger.Fatal(fmt.Errorf("could not run compose file: %v - %v", composeFilePaths, err))
	}

	s.BuildApp()
}

func (s *IntegrationTestSuite) TearDownSuite() {
	composeFilePaths := []string{s.dockerComposeFile}
	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, s.identifier)
	err := compose.WithEnv(s.env).Down().Error
	if err != nil {
		s.logger.Fatal(fmt.Errorf("could not run compose file: %v - %v", composeFilePaths, err))
	}
}

func (s *IntegrationTestSuite) SetupTest()                            {}
func (s *IntegrationTestSuite) BeforeTest(suiteName, testName string) {}
func (s *IntegrationTestSuite) AfterTest(suiteName, testName string)  {}
func (s *IntegrationTestSuite) TearDownTest()                         {}

func (s *IntegrationTestSuite) MakeTestRequest(r request, h http.Handler) *httptest.ResponseRecorder {
	request := httptest.NewRequest(r.method, r.target, strings.NewReader(r.body))
	request = request.WithContext(context.WithValue(
		request.Context(),
		app.SessionContextKey(r.sessionCookieName),
		r.sessionCookieValue,
	))
	request.Header.Set("Content-type", r.contentType)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, request)

	return w
}
