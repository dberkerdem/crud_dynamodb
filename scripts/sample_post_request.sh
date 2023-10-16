#!/bin/bash

URL="http://127.0.0.1:8080/set_state"
TOKEN="<YOUR_TOKEN>"

POST_DATA='{
  "ID": <YOUR_PK>,
  "State": "Test1",
  "Date": "2023-10-10",
  "Details": {
    "key1": "val1",
    "key2": "val2",
    "key3": 1
  }
}'

# Send the POST request using curl
curl -X POST -H "Content-Type: application/json" \
-H "X-API-Secret: $TOKEN" \
-d "$POST_DATA" $URL \
-v
