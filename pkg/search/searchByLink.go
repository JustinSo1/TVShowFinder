package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/JustinSo1/TVShowFinder/internal"
)

// JSONResponse json struct
type JSONResponse struct {
	Docs []struct {
		TitleRomanji string      `json:"title_romaji"`
		TitleEnglish string      `json:"title_english"`
		Similarity   float64     `json:"similarity"`
		Episode      interface{} `json:"episode"` // Can be int or array of ints
		Season       string      `json:"season"`
		IsAdult      bool        `json:"is_adult"`
	} `json:"docs"`
}

const searchByLinkURL = "https://trace.moe/api/search?url="

// ByLink searches image by image url
func ByLink(link string) string {
	reqBody, err := json.Marshal(map[string]string{})
	internal.HandleError(err)

	resp, err := http.Post(searchByLinkURL+link, "application/json", bytes.NewBuffer(reqBody))
	internal.HandleError(err)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	internal.HandleError(err)

	var animeResp JSONResponse
	json.Unmarshal(body, &animeResp)

	res := "Link: " + link + "\n" +
		"Title Romaji: " + animeResp.Docs[0].TitleRomanji + "\n" +
		"Title English: " + animeResp.Docs[0].TitleEnglish + "\n" +
		"Similarity: " + strconv.FormatFloat(animeResp.Docs[0].Similarity, 'f', 6, 64) + "\n" +
		"Episode Number: " + fmt.Sprintf("%v", animeResp.Docs[0].Episode) + "\n" +
		"Year & Season: " + animeResp.Docs[0].Season + "\n" +
		"Is Adult: " + strconv.FormatBool(animeResp.Docs[0].IsAdult) + "\n\n"
	return res
}
