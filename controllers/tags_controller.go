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
	"github.com/saiprasadkrishnamurthy/reference-data-api/utils"
)

// TagsController echos.
type TagsController struct {
	BaseController
}

// Tags from database.
// Tags echos.
// @Summary Get tags for domains.
// @Description Get tags for domains.
// @Produce  json
// @Param text query string true "Input Text"
// @Success 200 {object} models.Tags
// @Header 200 {string} Token "qwerty"
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /tags [get]
func (c *TagsController) Tags(rw http.ResponseWriter, r *http.Request, p httprouter.Params) error {
	c1 := make(chan map[string]interface{})
	c2 := make(chan map[string]interface{})
	inputQuery := r.URL.Query().Get("text")
	//domainType := r.URL.Query().Get("domain")

	tagValues := []string{}

	go dictionaryAPI(config.GetSpellCheckerAPI(), inputQuery, c1)
	spellCheckerResponse := fmt.Sprintf("%v", utils.ExtractResult(c1, 2)["word"])

	go dictionaryAPI(config.GetSoundsLikeAPI(), inputQuery, c2)
	soundsLikeResponse := fmt.Sprintf("%v", utils.ExtractResult(c2, 2)["word"])

	tagValues = append(tagValues, spellCheckerResponse, soundsLikeResponse)
	tags := models.Tags{InputText: inputQuery, Tags: utils.Unique(tagValues)}
	response, _ := json.Marshal(tags)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(response)
	return nil // no error
}

func dictionaryAPI(api string, word string, c chan map[string]interface{}) {
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
