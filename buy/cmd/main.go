package main

import (
	"buy/internal/adapter"
	"buy/internal/app"
	"buy/internal/domain"
	"buy/internal/usecase"
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type LambdaHandler struct {
	myApp app.MyApp
}

func NewLambdaHandler(buyRepository domain.BuyRepository) *LambdaHandler {
	buyUsecase := usecase.NewBuyUsecase(buyRepository)
	myApp := app.NewMyApp(*buyUsecase)
	return &LambdaHandler{
		myApp: *myApp,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := []byte(request.Body)
	var body domain.Buy
	err := json.Unmarshal(b, &body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}

	date := time.Now()
	layout := "01-02-2006 15:04:05"
	dateString := date.Format(layout)
	dateTime, err := time.Parse(layout, dateString)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error parsing date-time",
		}, nil
	}

	body.CreateDate = dateTime.String()

	err = h.myApp.HandleRequest(&body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "buy: success create",
	}, nil
}

func main() {
	tableName := "Buy"
	sess := session.Must(session.NewSession())
	actualDynamoDBClient := dynamodb.New(sess)
	buyRepository := adapter.NewDynamoDBRepository(tableName, actualDynamoDBClient)
	handler := NewLambdaHandler(buyRepository)
	lambda.Start(handler.HandleRequest)
}
