package driver

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func ConfigAws() (*dynamodb.Client, error) {
	configAws, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedCredentialsFiles([]string{"driver/data/credentials.aws"}),
		config.WithSharedConfigFiles([]string{"driver/data/config.aws"}),
	)

	return dynamodb.NewFromConfig(configAws), err
}
