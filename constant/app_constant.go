package constant

import (
	"log"
	"os"
	"path/filepath"
)

var Key string
var HashSum string

func GetFolderPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(exePath)
}
