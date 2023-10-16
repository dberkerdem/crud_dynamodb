#!/bin/bash

ID="<YOUR_PK>"
TOKEN="<YOUR_TOKEN>"
URL="http://127.0.0.1:8080/delete_state?id=$ID"

# Send the DELETE request using curl
curl -X DELETE $URL \
-H "X-Dynamo-Secret: $TOKEN"
