package app

import (
	"context"

	"github.com/alrund/yp-2-project/client/internal/domain/port"
	"github.com/alrund/yp-2-project/client/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type App struct {
	Config *Config
	Logger port.Logger
}

func NewApp(
	config *Config,
	logger port.Logger,
) *App {
	return &App{
		Config: config,
		Logger: logger,
	}
}

//var qs = []*survey.Question{
//	{
//		Name:      "name",
//		Prompt:    &survey.Input{Message: "What is your name?"},
//		Validate:  survey.Required,
//		Transform: survey.Title,
//	},
//	{
//		Name: "color",
//		Prompt: &survey.Select{
//			Message: "Choose a color:",
//			Options: []string{"red", "blue", "green"},
//			Default: "red",
//		},
//	},
//	{
//		Name:   "age",
//		Prompt: &survey.Input{Message: "How old are you?"},
//	},
//}

func (a *App) Run(client port.Client) error {
	cred, err := credentials.NewClientTLSFromFile(a.Config.CertFile, "")
	if err != nil {
		return err
	}
	conn, err := grpc.Dial(a.Config.ServerAddress, grpc.WithTransportCredentials(cred))
	if err != nil {
		return err
	}
	defer conn.Close()

	err = client.SetGRPCClient(proto.NewAppClient(conn))
	if err != nil {
		return err
	}

	//client.SelectLoginMethod(context.Background())

	err = client.Login(context.Background())
	if err != nil {
		return err
	}

	//resp, err := a.Client.Registration(context.Background(), &proto.RegistrationRequest{
	//	Login:    "qqq",
	//	Password: "www",
	//})
	//
	//if err != nil {
	//	return err
	//}
	//// the answers will be written to this struct
	//answers := struct {
	//	Name          string // survey will match the question and field names
	//	FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
	//	Age           int    // if the types don't match, survey will convert it
	//}{}
	//
	//// perform the questions
	//err := survey.Ask(qs, &answers)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return nil
	//}
	//
	//fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)

	return nil
}
