package nosql

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	region             = "sa-east-1"
	localstackEndpoint = "https://localhost.localstack.cloud:4566"
	id                 = "teste"
	secret             = "teste"
	token              = "teste"
)

func CreateSession() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Endpoint:    aws.String(localstackEndpoint),
		Credentials: credentials.NewStaticCredentials(id, secret, token),
	})

	if err != nil {
		return nil, err
	}

	return dynamodb.New(sess), nil
}

func Insert(table string, item interface{}) error {
	session, err := CreateSession()

	if err != nil {
		log.Print("Unable to create a dynamo session")
		return err
	}

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		log.Print("Error trying to marshal item")
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      av,
	}

	log.Println(input)

	_, err = session.PutItem(input)
	if err != nil {
		log.Print("Error trying to insert a new item")
		return err
	}

	return nil
}
