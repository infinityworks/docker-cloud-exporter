package config

import (
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
)

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func CheckConfig() {
	if os.Getenv("DOCKERCLOUD_USER") == "" {
		log.Error("User not set for API Access, required.")
	} else if os.Getenv("DOCKERCLOUD_APIKEY") == "" {
		log.Error("API Key not set for API Access, required.")
	}
}

func ConvertTime(input string) float64 {
	t, err := time.Parse(time.RFC1123Z, input)
	if err != nil {
		log.Println(err)
	}

	ts := t.Unix()

	return float64(ts)
}
