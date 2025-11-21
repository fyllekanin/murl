package files

import "os"

func GetInstallPath() string {
	value, existing := os.LookupEnv("MURL_PATH")
	if !existing {
		panic("Missing installation path for murl")
	}
	return value
}
