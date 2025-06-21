package nosql

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	model "http-gateway/internal/domain/models/order"
	"http-gateway/internal/settings"
)

func getLocalStackEndpoint() string {
	if settings.GetEnvs().SetupMode == "debug" {
		return settings.GetEnvs().LocalstackEndpointLocal
	}
	return settings.GetEnvs().LocalstackEndpointExternal
}

func CreateSession() (*dynamodb.DynamoDB, error) {

	newSession, err := session.NewSession(&aws.Config{
		Region:      aws.String(settings.GetEnvs().LocalStackRegion),
		Endpoint:    aws.String(getLocalStackEndpoint()),
		Credentials: credentials.NewStaticCredentials(settings.GetEnvs().LocalStackAWSId, settings.GetEnvs().LocalStackAWSSecret, settings.GetEnvs().LocalStackAWSToken),
	})

	if err != nil {
		return nil, err
	}

	return dynamodb.New(newSession), nil
}

func Insert(table string, item any) error {
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

	_, err = session.PutItem(input)
	if err != nil {
		log.Print("Error trying to insert a new item")
		return err
	}

	return nil
}

func Scan(table string) ([]model.Order, error) {
	var orders []model.Order
	var order model.Order

	session, err := CreateSession()

	if err != nil {
		log.Print("Unable to create a dynamo session")
		return nil, err
	}

	items, err := session.Scan(&dynamodb.ScanInput{
		TableName: aws.String(table),
	})

	if err != nil {
		return nil, err
	}

	for _, item := range items.Items {
		err = dynamodbattribute.UnmarshalMap(item, &order)
		if err != nil {
			log.Println("Deserializing error: " + err.Error())
			continue
		}
		orders = append(orders, order)
	}

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func ScanById(table string, id string) ([]model.Order, error) {
	var orders []model.Order
	var order model.Order

	session, err := CreateSession()

	if err != nil {
		log.Print("Unable to create a dynamo session")
		return nil, err
	}

	items, err := session.Scan(&dynamodb.ScanInput{
		TableName:        aws.String(table),
		FilterExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(id),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	for _, item := range items.Items {
		err = dynamodbattribute.UnmarshalMap(item, &order)
		if err != nil {
			log.Println("Deserializing error: " + err.Error())
			continue
		}
		orders = append(orders, order)
	}

	if err != nil {
		return nil, err
	}

	return orders, nil
}
