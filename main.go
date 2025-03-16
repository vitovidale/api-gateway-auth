package main

import (
	"context"
	"encoding/json"
	"log"
	"tech-challenge-lambda/config"

	"github.com/auth0/go-auth0/authentication"
	"github.com/auth0/go-auth0/authentication/oauth"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	env *config.Container
	// jwks *keyfunc.JWKS
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	var err error

	env, err = config.New()

	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// TODO: uncomment this when JWT validation is required
	// jwks, err = keyfunc.Get(env.Jwt.JwksUrl, keyfunc.Options{
	// 	RefreshInterval: time.Hour,
	// 	Client:          &http.Client{Timeout: 10 * time.Second},
	// })

	// if err != nil {
	// 	log.Fatal("Error loading JWKS:", err)
	// }
}

func main() {
	lambda.Start(HandleRequest)
	// localTest()
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User

	err := json.Unmarshal([]byte(request.Body), &user)

	if err != nil || user.Username == "" || user.Password == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Invalid request body",
		}, nil
	}

	auth, err := auth(user.Username, user.Password)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
			Body:       "Invalid credentials",
		}, nil
	}

	body, err := json.Marshal(auth)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal server error",
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

// MARK: test the function locally
// func localTest() {
// 	res, err := HandleRequest(context.TODO(), events.APIGatewayProxyRequest{
// 		Body: `{"username": "", "password": ""}`,
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	spew.Dump(res)
// }

func auth(username string, password string) (*oauth.TokenSet, error) {
	a, err := authentication.New(
		context.TODO(),
		env.Auth.Domain,
		authentication.WithClientID(env.Auth.Client),
		authentication.WithClientSecret(env.Auth.Secret),
	)

	if err != nil {
		return nil, err
	}

	tokens, err := a.OAuth.LoginWithPassword(context.TODO(), oauth.LoginWithPasswordRequest{
		Username: username,
		Password: password,
	}, oauth.IDTokenValidationOptions{})

	if err != nil {
		if err2, ok := err.(*authentication.Error); ok {
			if err2.Err == "mfa_required" {
				return nil, err2
			}
		}
		return nil, err
	}

	return tokens, nil
}
