#!/bin/bash

ID="<YOUR_PK>"
TOKEN="<YOUR_TOKEN>"
URL="http://127.0.0.1:8080/get_state?id=$ID"

# Send the GET request using curl
curl $URL -H "X-API-Secret: $TOKEN"
