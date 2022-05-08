package configs

import (
    "os"
    "time"

    _ "github.com/joho/godotenv/autoload" // This will load .env file if exists
)

var (
	Mode     string
	Port     string
	LogLevel string

	GRPC struct {
		ConnectTimeout time.Duration
		HTTP           string
		Port           string
		Server         struct {
			Authz     string
			Deploy    string
			Account   string
			Component string
		}
	}

	Sentry struct {
		Dsn              string
		TracesSampleRate float64
	}

	Database struct {
		Host     string
		Port     string
		Name     string
		Username string
		Password string
	}

	Github struct {
		User  string
		Token string
	}

	RSA struct {
		PublicKey  string
		PrivateKey string
	}
	Dev struct {
		FrontendDevPort string
	}
)

func init() {
    Mode = os.Getenv("MODE")
    Port = os.Getenv("PORT")

	GRPC.Port = os.Getenv("GRPC_PORT")
	GRPC.Server.Authz = os.Getenv("AUTHZ_SERVER")
	GRPC.Server.Deploy = os.Getenv("GRPC_DEPLOY_SERVER")
	GRPC.Server.Account = os.Getenv("GRPC_ACCOUNT_SERVICE")
	GRPC.Server.Component = os.Getenv("GRPC_COMPONENT_SERVER")

    Database.Host = os.Getenv("DB_HOST")
    Database.Port = os.Getenv("DB_PORT")
    Database.Name = os.Getenv("DB_NAME")
    Database.Username = os.Getenv("DB_USERNAME")
    Database.Password = os.Getenv("DB_PASSWORD")

    Github.User = os.Getenv("GITHUB_USER")
    Github.Token = os.Getenv("GITHUB_TOKEN")

    RSA.PublicKey = os.Getenv("RSA_PUBLIC_KEY")
    RSA.PublicKey = os.Getenv("RSA_PRIVATE_KEY")

    Dev.FrontendDevPort = os.Getenv("FRONTEND_PORT")
}

func IsProduction() bool {
    return os.Getenv("MODE") == "production"
}

func IsStaging() bool {
    return os.Getenv("MODE") == "staging"
}

func IsDev() bool {
    return !IsProduction() && !IsStaging()
}
