# CRUD DynamoDB Service

This service is a Go application designed to provide CRUD (Create, Read, Update, Delete) operations on a DynamoDB table. It's containerized with Docker for easy deployment and scalability.

## Features

- CRUD operations on DynamoDB.
- Middleware for authentication and logging.
- Error handling utilities.
- Sample scripts for testing the API endpoints.

## Prerequisites

- Go 1.21.1 or later.
- Docker (for building and running the Docker image).
- AWS account with access to DynamoDB.
- Set the necessary environment variables: `AWS_REGION`, `DYNAMO_TABLE_NAME`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, and `X_API_Secret`.

## Getting Started

Clone the repository to your local machine:

```bash
git clone https://github.com/dberkerdem/crud_dynamodb.git
```
Navigate to the cloned directory:
```bash
cd crud_dynamodb
```
### Building Docker Image
Run the following command to build the Docker image:
```bash
make build-image
```

### Running Service Locally
Ensure you have the necessary environment variables set.\
To run the service locally using Docker, execute:
```bash
make run-service-locally
```

## API Endpoints

The service exposes the following endpoints:

- **GET /get_state:** Retrieves the state based on the provided PK.
- **POST /set_state:** Creates a new state entry.
- **PUT /update_state:** Updates an existing state.
- **DELETE /delete_state:** Deletes a state based on the provided PK.
Refer to the scripts/ directory for sample cURL requests to these endpoints.

# LICENSE
License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.