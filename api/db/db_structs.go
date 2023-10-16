// db/db_structs.go
package db

// Modify this struct based on the dynamodb table
type State struct {
	ID      string                 `json:"ID"`
	State   string                 `json:"State"`
	Date    string                 `json:"Date"`
	Details map[string]interface{} `json:"Details"`
}
