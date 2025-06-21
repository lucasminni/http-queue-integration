package settings

import (
	"github.com/caarlos0/env/v7"
)

var envs Environment

type Environment struct {
	DynamoDBEnvironment
	HTTPServerEnvironment
	LocalStackEnvironment
	SetupEnvironment
}

type DynamoDBEnvironment struct {
}

type HTTPServerEnvironment struct {
	HTTPServerPort string `env:"HTTP_SERVER_PORT,required"  envDefault:"8080"`
}

type LocalStackEnvironment struct {
	LocalstackEndpointLocal    string `env:"LOCALSTACK_ENDPOINT_LOCAL,required" envDefault:"https://localhost.localstack.cloud:4566"`
	LocalstackEndpointExternal string `env:"LOCALSTACK_ENDPOINT_EXTERNAL,required" envDefault:"http://localstack:4566"`
	LocalStackRegion           string `env:"LOCALSTACK_REGION,required" envDefault:"sa-east-1"`
	LocalStackAWSId            string `env:"LOCALSTACK_AWS_ID,required" envDefault:"teste"`
	LocalStackAWSSecret        string `env:"LOCALSTACK_AWS_SECRET,required" envDefault:"teste"`
	LocalStackAWSToken         string `env:"LOCALSTACK_AWS_TOKEN,required" envDefault:"teste"`
}

type SetupEnvironment struct {
	Mode string `env:"MODE,required" envDefault:"debug"`
}

func GetEnvs() Environment {
	return envs
}

func LoadEnvs() error {
	if err := env.Parse(&envs); err != nil {
		return err
	}
	return nil
}
