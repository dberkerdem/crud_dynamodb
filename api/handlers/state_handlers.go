package handlers

import (
	"crud_dynamo/db"
	"crud_dynamo/utils"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/valyala/fasthttp"
)

func GetStateHandler(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")

	if id == nil || string(id) == "" {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusBadRequest, "Invalid input: ID is required"))
		return
	}

	result, err := db.DB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(string(id)),
			},
		},
	})
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusInternalServerError, fmt.Sprintf("Error fetching state: %v", err)))
		return
	}

	if result.Item == nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusNotFound, "State not found"))
		return
	}

	var stateEntry db.State
	err = dynamodbattribute.UnmarshalMap(result.Item, &stateEntry)
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusInternalServerError, fmt.Sprintf("Error unmarshalling state: %v", err)))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.SetContentType("application/json")
	json.NewEncoder(ctx).Encode(stateEntry)
}

func PostStateHandler(ctx *fasthttp.RequestCtx) {
	var stateEntry db.State
	err := json.Unmarshal(ctx.PostBody(), &stateEntry)
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusBadRequest, "Invalid JSON format"))
		return
	}

	if stateEntry.ID == "" || stateEntry.State == "" || stateEntry.Date == "" {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusBadRequest, "Invalid input: ID, State, and Date are required"))
		return
	}

	av, err := dynamodbattribute.MarshalMap(stateEntry)
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusInternalServerError, fmt.Sprintf("Error marshalling state: %v", err)))
		return
	}

	_, err = db.DB.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusInternalServerError, fmt.Sprintf("Error creating state: %v", err)))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func UpdateStateHandler(ctx *fasthttp.RequestCtx) {
	var stateEntry db.State
	err := json.Unmarshal(ctx.PostBody(), &stateEntry)
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusBadRequest, "Invalid JSON format"))
		return
	}

	if stateEntry.ID == "" {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusBadRequest, "Invalid input: ID is required"))
		return
	}

	// Convert the Details field to the format expected by DynamoDB
	detailsAttributeValue, err := dynamodbattribute.MarshalMap(stateEntry.Details)
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusInternalServerError, fmt.Sprintf("Error marshalling Details: %v", err)))
		return
	}

	updateExpression := "set #state = :state, #date = :date, #details = :details"
	expressionAttributeNames := map[string]*string{
		"#state":   aws.String("State"),
		"#date":    aws.String("Date"),
		"#details": aws.String("Details"),
	}
	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":state": {
			S: aws.String(stateEntry.State),
		},
		":date": {
			S: aws.String(stateEntry.Date),
		},
		":details": {
			M: detailsAttributeValue,
		},
	}

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames:  expressionAttributeNames,
		ExpressionAttributeValues: expressionAttributeValues,
		TableName:                 aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(stateEntry.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String(updateExpression),
	}

	_, err = db.DB.UpdateItem(input)
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusInternalServerError, fmt.Sprintf("Failed to update state: %v", err)))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.SetBodyString("State updated successfully")
}

func DeleteStateHandler(ctx *fasthttp.RequestCtx) {
	id := ctx.QueryArgs().Peek("id")

	if id == nil || string(id) == "" {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusBadRequest, "Invalid input: ID is required"))
		return
	}

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(string(id)),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := db.DB.DeleteItem(input)
	if err != nil {
		utils.ErrorHandler(ctx, utils.NewHTTPError(fasthttp.StatusInternalServerError, fmt.Sprintf("Failed to delete state: %v", err)))
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}
