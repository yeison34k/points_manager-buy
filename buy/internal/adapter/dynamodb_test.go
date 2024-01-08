package adapter_test

import (
	"buy/internal/adapter"
	"buy/internal/domain"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDynamoDB struct {
	dynamodbiface.DynamoDBAPI
	mock.Mock
}

func (m *MockDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.GetItemOutput), args.Error(1)
}

func (m *MockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	args := m.Called(input)
	return args.Get(0).(*dynamodb.PutItemOutput), args.Error(1)
}

func TestDynamoDBRepository_CreateBuy(t *testing.T) {
	mockDynamoDB := new(MockDynamoDB)

	repo := &adapter.DynamoDBRepository{
		DynamoDBClient: mockDynamoDB,
		TableName:      "TestTableName",
	}

	mockDynamoDB.On("PutItem", mock.AnythingOfType("*dynamodb.PutItemInput")).Return(
		&dynamodb.PutItemOutput{},
		nil,
	)

	err := repo.CreateBuy(&domain.Buy{
		ID:          "123",
		User:        "TestUser",
		ProductName: "TestPoint",
		Price:       123.45,
		Points:      10,
	})

	mockDynamoDB.AssertExpectations(t)
	assert.NoError(t, err)
}
