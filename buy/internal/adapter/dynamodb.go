package adapter

import (
	"buy/internal/domain"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDBRepository struct {
	DynamoDBClient DynamoDBAPI
	TableName      string
}

type DynamoDBAPI interface {
	PutItem(*dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}

func NewDynamoDBRepository(tableName string, client DynamoDBAPI) *DynamoDBRepository {
	return &DynamoDBRepository{
		DynamoDBClient: client,
		TableName:      tableName,
	}
}

func (r *DynamoDBRepository) CreateBuy(buy *domain.Buy) error {
	totalAsString := fmt.Sprintf("%.2f", buy.Price)
	points := fmt.Sprintf("%v", buy.Price)

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(buy.ID),
			},
			"buyId": {
				S: aws.String(buy.BuyId),
			},
			"user": {
				S: aws.String(buy.User),
			},
			"product": {
				S: aws.String(buy.ProductName),
			},
			"price": {
				N: aws.String(totalAsString),
			},
			"points": {
				N: &points,
			},
			"createDate": {
				S: aws.String(buy.CreateDate),
			},
		},
	}

	_, err := r.DynamoDBClient.PutItem(input)
	return err
}
