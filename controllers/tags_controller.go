package controllers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/julienschmidt/httprouter"
	"github.com/saiprasadkrishnamurthy/reference-data-api/config"
	"github.com/saiprasadkrishnamurthy/reference-data-api/models"
	"github.com/saiprasadkrishnamurthy/reference-data-api/utils"
)

// TagsController echos.
type TagsController struct {
	BaseController
}

const totalCandidates = 20

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
	input := r.URL.Query().Get("text")
	//domainType := r.URL.Query().Get("domain")

	inputTokens := strings.Fields(input)

	tokens := config.GetTokensFromDictionary()

	var tagsOfAllWords = [][]string{}

	for _, t := range inputTokens {
		s := findWords(t, tokens)
		tagsOfAllWords = append(tagsOfAllWords, s)
	}
	o := []string{}
	for i := 0; i <= config.SpellCheckerTopN(); i++ {
		s := ""
		for _, v := range tagsOfAllWords {
			if i < len(v) {
				s += v[i] + " "
			}
		}
		if len(strings.TrimSpace(s)) > 0 {
			o = append(o, strings.TrimSpace(s))
		}
	}

	tags := models.Tags{InputText: input, Tags: utils.Unique(o)}
	response, _ := json.Marshal(tags)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(response)
	return nil // no error
}

func findWords(input string, tokens []string) []string {
	closestIndex := sort.Search(len(tokens), func(i int) bool { return input <= tokens[i] })
	radius := totalCandidates / 2

	var start, end int
	if start = closestIndex - radius; start <= 0 {
		start = 0
	}
	if end = closestIndex + radius; start >= len(tokens) {
		end = 0
	}
	probables := make([]string, totalCandidates)
	copy(probables, tokens[start:end])
	sort.Slice(probables, func(i, j int) bool {
		iDistance := levenshtein.ComputeDistance(input, probables[i])
		jDistance := levenshtein.ComputeDistance(input, probables[j])
		return iDistance <= jDistance
	})
	if len(probables) <= config.SpellCheckerTopN() {
		return probables
	}
	return probables[0:config.SpellCheckerTopN()]

}
