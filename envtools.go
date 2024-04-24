package envtools

import (
	"log"
	"github.com/joho/godotenv"
)

func GetSingleEnv(envvar string, envpath string) string {

	if envpath == "" {
		envpath = ".env"
	}

	envvalue, err := godotenv.Read(envpath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	param := envvalue[envvar]

	if param == "" {
		log.Fatalf("%v is not set.", envvar)
	}


	return param

}

func GetAllEnv(envpath string) map[string]string {

	if envpath == "" {
		envpath = ".env"
	}

	envMap, err := godotenv.Read(envpath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return envMap

}
