package botutils

import (
	"log"
	"os"
)

func (*env) Check(envVars []string) {
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			log.Fatal("[ERROR] Environment variable " + envVar + " was not supplied.\n")
		}
	}
}
