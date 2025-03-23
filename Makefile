# Alvo para construir a função Lambda
build-AuthUser:
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

# Alvo padrão que será invocado pelo SAM
build: build-AuthUser
