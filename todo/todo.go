package todo

import (
	"log"

	homedir "github.com/mitchellh/go-homedir"
)

// GetHomeDir returns homedir
func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory.")
	}
	return home
}

// GetVersion returns version name, and code
func GetVersion() string {
	var version = "1.0.0"
	return version
}
