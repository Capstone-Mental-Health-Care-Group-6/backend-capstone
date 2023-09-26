package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ProgrammingConfig struct {
	ServerPort int
	DBPort     string
	DBHost     string
	DBUser     string
	DBPass     string
	DBName     string
	Secret     string
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

func loadConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file,", err.Error())
	}

	if val, found := os.LookupEnv("SERVER"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config : Invalid port value,", err.Error())
			return nil
		}

		res.ServerPort = port
	}

	if val, found := os.LookupEnv("DBPORT"); found {
		res.DBPort = val
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

	return res
}
