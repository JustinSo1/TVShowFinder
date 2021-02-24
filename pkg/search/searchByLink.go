package search

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/JustinSo1/TVShowFinder/internal"
	"github.com/JustinSo1/TVShowFinder/pkg/convert"
)

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

	var animeResp internal.JsonResponse
	json.Unmarshal(body, &animeResp)

	var res string
	res = "Title Romaji: " + animeResp.Docs[0].TitleRomanji + "\n" +
		"Title English: " + animeResp.Docs[0].TitleEnglish + "\n" +
		"Similarity: " + strconv.FormatFloat(animeResp.Docs[0].Similarity, 'f', 6, 64) + "\n" +
		"Episode Number: " + convert.IntArrToString(animeResp.Docs[0].Episode) + "\n" +
		"Year & Season: " + animeResp.Docs[0].Season + "\n" +
		"Is Adult: " + strconv.FormatBool(animeResp.Docs[0].IsAdult) + "\n"
	return res
}
