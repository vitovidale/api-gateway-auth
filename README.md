# Tech Challenge fase 3 - AWS Lambda

## Build

```bash
docker build -t lambda-builder .
```

## Package

```bash
sam package --output-template-file packaged.yaml --s3-bucket local-lambdas
```

## Deploy

```bash
sam deploy --template-file packaged.yaml --stack-name lambda-stack --capabilities CAPABILITY_IAM
```

## Invoke

```bash
aws lambda invoke --function-name lambda-stack-MyGoFunction-u3uY0rBXOJ3Z --cli-binary-format raw-in-base64-out --payload '{ "name": "Rodrigo Longhi" }' out.txt
```
