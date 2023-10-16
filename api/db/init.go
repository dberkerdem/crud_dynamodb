// db/db_conn.go
package db

import (
	"crud_dynamo/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var DB *dynamodb.DynamoDB

func InitDB(cfg config.Config) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(cfg.AWS.Region),
	}))
	DB = dynamodb.New(sess)
}
