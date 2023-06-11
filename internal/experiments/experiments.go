package experiments

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const envPrefix = "TASK_X_"

const (
	GentleForce = "GENTLE_FORCE"
)

var flags = map[string]bool{
	GentleForce: false,
}

func IsEnabled(xName string) bool {
	return flags[xName]
}

func Parse() error {
	if err := readDotEnv(); err != nil {
		return err
	}
	for xName := range flags {
		flags[xName] = parseEnv(xName)
	}
	return nil
}

func parseEnv(xName string) bool {
	envName := fmt.Sprintf("%s%s", envPrefix, xName)
	return os.Getenv(envName) == "1"
}

func readDotEnv() error {
	env, err := godotenv.Read()
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	// If the env var is an experiment, set it.
	for key, value := range env {
		if strings.HasPrefix(key, envPrefix) {
			os.Setenv(key, value)
		}
	}
	return nil
}
