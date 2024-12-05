package env

import (
	"os"
	"strconv"
)

func GetEnvString(key, fallback string) string {
	val := os.Getenv(key)

	if (len(val) == 0) {
		return fallback
	}

	return val
}

func GetEnvInt(key string, fallback int) int  {
	val := os.Getenv(key)

	if (len(val) == 0) {
		return fallback
	}

	res, err := strconv.Atoi(val);

	if(err != nil){
		return fallback
	}

	return res
}
