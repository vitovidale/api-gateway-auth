# api-gateway-auth

> Função serverless para autenticação via Auth0 e API Gateway

---

## Introdução

Esta Lambda expõe o endpoint `/auth/login`, recebe JSON com `username` e `password`, autentica contra Auth0 e retorna tokens JWT.

---

## Deploy

O deploy é feito automaticamente via GitHub Actions (workflow_dispatch).

### Como executar

1. No GitHub, abra este repositório e clique em **Actions**  
2. Selecione o workflow **Deploy API Gateway Auth**  
3. Clique em **Run workflow**

---

## Manual Build & Deploy

```bash
docker build -t lambda-builder .
docker run --name temp-container lambda-builder
docker cp temp-container:/app/bootstrap ./bootstrap
docker rm temp-container
sam package --output-template-file packaged.yaml --s3-bucket local-lambdas
sam deploy --template-file packaged.yaml --stack-name auth --capabilities CAPABILITY_IAM
aws lambda invoke --function-name auth-AuthUser-<YOUR_ID> --cli-binary-format raw-in-base64-out --payload '{ "username": "email@pm.me", "password": "" }' outfile.json
