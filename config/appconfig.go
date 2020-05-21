package config

import (
	"os"
)

// InitConfigs initializes all the necessary configs once.
func InitConfigs() {
	initDB()
}

func initDB() {

}

// GetPort returns the Web server port.
func GetPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8082"
	}
	return port
}

// GetSpellCheckerAPI returns the spell checker api url.
func GetSpellCheckerAPI() string {
	return "https://api.datamuse.com/words?sp=%s"
}

// GetSoundsLikeAPI returns the sounds like api url.
func GetSoundsLikeAPI() string {
	return "https://api.datamuse.com/words?sl=%s"
}

// APIVersion public API version.
func APIVersion() string {
	return "v1"
}
