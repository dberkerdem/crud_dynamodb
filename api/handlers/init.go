package handlers

import (
	"crud_dynamo/config"
)

var tableName string

func InitHandlers(cfg config.Config) {
	tableName = cfg.AWS.DynamoTable
}
