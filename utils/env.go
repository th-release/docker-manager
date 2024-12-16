package utils

import (
	"os"
)

func GetEnv(value string) (string, error) {
	return os.Getenv(value), nil
}
