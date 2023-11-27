package configs

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type ProgrammingConfig struct {
	ServerPort              int
	DBPort                  int
	DBHost                  string
	DBUser                  string
	DBPass                  string
	DBName                  string
	Secret                  string
	RefSecret               string
	MidtransServerKey       string
	MidtransClientKey       string
	MidtransEnvironment     string
  CloudinaryURL           string
	OauthGoogleClientID     string
	OauthGoogleClientSecret string
	OauthGoogleRedirectURL  string

func InitConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)

	// err := godotenv.Load(".env")

	// if err != nil {
	// 	logrus.Error("Config : Cannot load config file, ", err.Error())
	// 	return nil
	// }

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			return nil
		}

		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			return nil
		}

		res.DBPort = port
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefSecret = val
	}

	if val, found := os.LookupEnv("MT_SERVER_KEY"); found {
		res.MidtransServerKey = val
	}

	if val, found := os.LookupEnv("MT_CLIENT_KEY"); found {
		res.MidtransClientKey = val
	}

	if val, found := os.LookupEnv("MT_ENV"); found {
		res.MidtransEnvironment = val
	}

	if val, found := os.LookupEnv("OAUTH_GOOGLE_CLIENT_ID"); found {
		res.OauthGoogleClientID = val
	}

	if val, found := os.LookupEnv("OAUTH_GOOGLE_CLIENT_SECRET"); found {
		res.OauthGoogleClientSecret = val
	}

	if val, found := os.LookupEnv("OAUTH_GOOGLE_REDIRECT_URL"); found {
		res.OauthGoogleRedirectURL = val

	if val, found := os.LookupEnv("CloudURL"); found {
		res.CloudinaryURL = val
	}

	return res
}
