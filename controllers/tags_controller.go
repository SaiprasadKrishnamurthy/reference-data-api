package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"io/ioutil"

	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/reference-data-api/config"
	"github.com/saiprasadkrishnamurthy/reference-data-api/models"
)

// TagsController echos.
type TagsController struct {
	BaseController
}

// Tags get tags.
func (c *TagsController) Tags(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	c1 := make(chan map[string]interface{})
	c2 := make(chan map[string]interface{})
	inputQuery := r.URL.Query().Get("text")
	//domainType := r.URL.Query().Get("domain")

	go dictionaryApi(config.GetSpellCheckerAPI(), inputQuery, c1)
	spellCheckerResponse := fmt.Sprintf("%v", extractResult(c1, 2)["word"])

	go dictionaryApi(config.GetSoundsLikeAPI(), inputQuery, c2)
	soundsLikeResponse := fmt.Sprintf("%v", extractResult(c2, 2)["word"])

	tagValues := []string{spellCheckerResponse, soundsLikeResponse}

	temp := tagValues[:0]
	for _, x := range tagValues {
		if x != inputQuery {
			temp = append(temp, x)
		}
	}
	tagValues = temp
	tags := models.Tags{InputText: inputQuery, Tags: tagValues}
	response, _ := json.Marshal(tags)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(response)
	return nil // no error
}

func dictionaryApi(api string, word string, c chan map[string]interface{}) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	url := fmt.Sprintf(api, strings.Replace(word, " ", "%20", -1))
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

func extractResult(c chan map[string]interface{}, timeoutSeconds int) map[string]interface{} {
	select {
	case res := <-c:
		return res
	case <-time.After(time.Duration(timeoutSeconds) * time.Second):
		return map[string]interface{}{}
	}
}
