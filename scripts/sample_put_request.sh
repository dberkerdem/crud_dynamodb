#!/bin/bash

URL="http://127.0.0.1:8080/update_state"
TOKEN="<YOUR-TOKEN>"  # Replace with your actual token

UPDATE_DATA='{
  "ID": <YOUR_PK>,
  "State": "IN-PROGRESS",
  "Date": "2023-10-11",
  "Details": {
    "key1": "newVal1",
    "key2": "newVal2",
    "key3": 2
  }
}'

# Send the PUT request using curl
curl -X PUT -H "Content-Type: application/json" \
-H "X-Dynamo-Secret: $TOKEN" \
-d "$UPDATE_DATA" $URL \
-v
