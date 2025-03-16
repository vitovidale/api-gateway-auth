# Tech Challenge fase 3 - AWS Lambda

## Introduction

This is a simple lambda function that authenticates a user with username and password using the Auth0 provider.

## Build the Docker image to build the lambda (agnostic solution)

```bash
docker build -t lambda-builder .
```

# Run the Docker image to copy the binary to the local machine

```bash
docker run --name temp-container lambda-builder
docker cp temp-container:/app/bootstrap ./bootstrap
docker rm temp-container
```

## Package the Lambda code and create a SAM template

```bash
sam package --output-template-file packaged.yaml --s3-bucket local-lambdas
```

## Deploy the Lambda using SAM

```bash
sam deploy --template-file packaged.yaml --stack-name auth --capabilities CAPABILITY_IAM
```

## Invoke

```bash
aws --cli-auto-prompt
aws lambda invoke --function-name XPTO --cli-binary-format raw-in-base64-out --payload '{ "username": "email@pm.me", "password": "" }' outfile.json
```
