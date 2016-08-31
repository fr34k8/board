package daemon

import (
	"os"
	"github.com/kardianos/osext"
)

type Settings struct {
	ListenAddress string `json:"listenAddress"`
	DatabaseURL string `json:"databaseURL"`
	StaticDirectory string `json:"publicDirectory"`
}

func NewSettings() *Settings {

	// Default static file directory is /static in the same directory as
	// the executable
	staticDirectory, err := osext.ExecutableFolder()
    if err == nil {
    	staticDirectory = staticDirectory + "/static"
    } else {
        staticDirectory = ""
    }
	
	// Initialize with default values
	settings := &Settings{
		ListenAddress: ":8080",
		DatabaseURL: "",
		StaticDirectory: staticDirectory,
	}

	val := os.Getenv("LISTEN_ADDRESS")
	if val != "" {
		settings.ListenAddress = val
	}

	val = os.Getenv("DATABASE")
	if val != "" {
		settings.DatabaseURL = val
	}
	
	return settings
}
