# Phony Targets
.PHONY: build-image run-service-locally

build-image:
	docker build -t crud_dynamo:latest .

run-service-locally:
	docker run --rm -it -p 8080:8080 -e AWS_REGION=${AWS_REGION} -e DYNAMO_TABLE_NAME=${DYNAMO_TABLE_NAME} -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} -e X_API_Secret=${X_API_Secret} crud_dynamo:latest
