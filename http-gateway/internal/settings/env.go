package settings

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

var envs Environment

type Environment struct {
	DynamoDBEnvironment
	HTTPServerEnvironment
	LocalStackEnvironment
	SetupEnvironment
}

type DynamoDBEnvironment struct {
	OrderDynamoTable string `env:"ORDER_DYNAMO_TABLE,required"`
}

type HTTPServerEnvironment struct {
	HTTPServerPort string `env:"HTTP_SERVER_PORT,required"`
}

type LocalStackEnvironment struct {
	LocalstackEndpointLocal    string `env:"LOCALSTACK_ENDPOINT_LOCAL,required"`
	LocalstackEndpointExternal string `env:"LOCALSTACK_ENDPOINT_EXTERNAL,required"`
	LocalStackRegion           string `env:"LOCALSTACK_REGION,required"`
	LocalStackAWSId            string `env:"LOCALSTACK_AWS_ID,required"`
	LocalStackAWSSecret        string `env:"LOCALSTACK_AWS_SECRET,required"`
	LocalStackAWSToken         string `env:"LOCALSTACK_AWS_TOKEN,required"`
}

type SetupEnvironment struct {
	SetupMode string `env:"SETUP_MODE,required"`
}

func GetEnvs() Environment {
	return envs
}

func LoadEnvs() error {
	_ = godotenv.Load(".env")

	if err := env.Parse(&envs); err != nil {
		return err
	}
	return nil
}
