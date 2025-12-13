package env

import (
	"os"
	"strconv"
)

func GetString(key, fallbck string) string {

	val, ok := os.LookupEnv(key)
	if !ok {
		return fallbck
	}
	return val

}

func GetInt(key string, fallbck int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallbck
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallbck
	}
	return valAsInt
}
