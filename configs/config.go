package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

type ProgrammingConfig struct {
	ServerPort              int
	DBPort                  uint16
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
  OpenAI                  string
	DbMongoUrl              string
	DbMongoName             string
	EmailSender             string
	EmailPasswordSender     string
	BaseURLFE               string
}

func InitConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func readData() *ProgrammingConfig {
	var data = new(ProgrammingConfig)
	data = loadConfig()

	if data == nil {
		err := godotenv.Load(".env")
		data = loadConfig()
		if err != nil || data == nil {
			return nil
		}
	}
	return data
}

func loadConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	var permit = true

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			permit = false
		}
		res.ServerPort = port
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			permit = false
		}

		res.DBPort = uint16(port)
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBHOST"); found {
		res.DBHost = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBUSER"); found {
		res.DBUser = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBPASS"); found {
		res.DBPass = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DBNAME"); found {
		res.DBName = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.RefSecret = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("MT_SERVER_KEY"); found {
		res.MidtransServerKey = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("MT_CLIENT_KEY"); found {
		res.MidtransClientKey = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("MT_ENV"); found {
		res.MidtransEnvironment = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("OAUTH_GOOGLE_CLIENT_ID"); found {
		res.OauthGoogleClientID = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("OAUTH_GOOGLE_CLIENT_SECRET"); found {
		res.OauthGoogleClientSecret = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("OAUTH_GOOGLE_REDIRECT_URL"); found {
		res.OauthGoogleRedirectURL = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("CloudURL"); found {
		res.CloudinaryURL = val
	} else {
		permit = false
	}
  
	if val, found := os.LookupEnv("KEY_OPEN_AI"); found {
		res.OpenAI = val
  } else {
		permit = false
	}
  
	if val, found := os.LookupEnv("EMAIL_SENDER"); found {
		res.EmailSender = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DB_MONGO_URL"); found {
		res.DbMongoUrl = val
  } else {
		permit = false
	}
  
	if val, found := os.LookupEnv("EMAIL_PASSWORD_SENDER"); found {
		res.EmailPasswordSender = val
	} else {
		permit = false
	}

	if val, found := os.LookupEnv("DB_MONGO_NAME"); found {
		res.DbMongoName = val
  } else {
		permit = false
	}
  
	if val, found := os.LookupEnv("BASE_URL_FE"); found {
		res.BaseURLFE = val
	} else {
		permit = false
	}

	if !permit {
		return nil
	}

	return res
}
