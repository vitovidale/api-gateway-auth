package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		Auth *Auth
		Jwt  *Jwt
	}

	Jwt struct {
		JwksUrl string
	}

	Auth struct {
		Domain   string
		Audience string
		Client   string
		Secret   string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	auth := &Auth{
		Domain:   os.Getenv("AUTH0_DOMAIN"),
		Audience: os.Getenv("AUTH0_AUDIENCE"),
		Client:   os.Getenv("AUTH0_CLIENT_ID"),
		Secret:   os.Getenv("AUTH0_CLIENT_SECRET"),
	}

	jwt := &Jwt{
		JwksUrl: fmt.Sprintf("https://%s/.well-known/jwks.json", auth.Domain),
	}

	return &Container{
		auth,
		jwt,
	}, nil
}
