package persistance

import (
	"log"
	"os"
)

func getBasePath() string {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return dirname + "/.murl"
}
