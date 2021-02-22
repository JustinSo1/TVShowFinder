package internal

// Response json struct
type JsonResponse struct {
	Docs []struct {
		TitleRomanji string  `json:"title_romaji"`
		TitleEnglish string  `json:"title_english"`
		Similarity   float64 `json:"similarity"`
		Episode      []int   `json:"episode"`
		Season       string  `json:"season"`
		IsAdult      bool    `json:"is_adult"`
	} `json:"docs"`
}
