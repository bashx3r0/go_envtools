package envtools

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetSingleEnv(envvar string, envpath string) string {

	if envpath == "" {
		envpath = ".env"
	}

	err := godotenv.Load(envpath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	param := os.Getenv(envvar)

	if param == "" {
		log.Fatalf("%v is not set.", envvar)
	}


	return param

}

func GetAllEnv(envpath string) map[string]string {

	if envpath == "" {
		envpath = ".env"
	}

	err := godotenv.Load(envpath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	envMap := make(map[string]string)

	for _, env := range os.Environ() {
        
        pair := strings.SplitN(env, "=", 2)
        if len(pair) != 2 {
            continue // will variable with invalid format
        }
        
        envMap[pair[0]] = pair[1]
    }

	return envMap

}