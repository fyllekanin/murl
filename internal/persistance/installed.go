package persistance

import (
	"os"
)

func IsVersionInstalled(softwareName string, version string) bool {
	basePath := getBasePath()
	if _, err := os.Stat(basePath + "/" + softwareName + "/" + version); !os.IsNotExist(err) {
		return true
	}
	return false
}
