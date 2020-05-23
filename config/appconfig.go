package config

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sort"
)

var tokens []string

// InitConfigs initializes all the necessary configs once.
func InitConfigs() {

	path, _ := filepath.Abs("dictionaries/en-GB.txt")
	loadDictionary(path)

}

// GetPort returns the Web server port.
func GetPort() string {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8082"
	}
	return port
}

// SpellCheckerTopN returns the sounds like api url.
func SpellCheckerTopN() int {
	return 3
}

// APIVersion public API version.
func APIVersion() string {
	return "v1"
}

// GetTokensFromDictionary gets the word dictionary.
func GetTokensFromDictionary() []string {
	return tokens
}

func loadDictionary(dictionaryPath string) {
	tokens = []string{}

	file, err := os.Open(dictionaryPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		token := scanner.Text()
		tokens = append(tokens, token)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	sort.Strings(tokens)
}
