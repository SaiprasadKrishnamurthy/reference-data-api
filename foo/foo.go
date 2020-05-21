package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	c1 := make(chan map[string]interface{})
	c2 := make(chan map[string]interface{})
	word := "Mangement Studies"
	go spellChecker(word, c1)
	spellCheckerResponse := fmt.Sprintf("%v", extractResult(c1, 2)["word"])
	go domainSpecificWordsChecker(spellCheckerResponse, c2)
	domainSpecificWordsCheckerResponse := fmt.Sprintf("%v", extractResult(c2, 2)["word"])
	fmt.Println(domainSpecificWordsCheckerResponse)
}

func spellChecker(word string, c chan map[string]interface{}) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := "https://api.datamuse.com/words?sp=" + strings.Replace(word, " ", "%20", -1)
	resp, error := client.Get(url)
	if error == nil {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		var result []map[string]interface{}
		e := json.Unmarshal(bodyBytes, &result)
		if e == nil && len(result) > 0 {
			c <- result[0]
			return
		}
		c <- map[string]interface{}{"word": word}
		return
	}
	c <- map[string]interface{}{"word": word}
}

func domainSpecificWordsChecker(word string, c chan map[string]interface{}) {
	c <- map[string]interface{}{"word": " ==> " + word}
}

func extractResult(c chan map[string]interface{}, timeoutSeconds int) map[string]interface{} {
	select {
	case res := <-c:
		return res
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		return map[string]interface{}{}
	}
}
