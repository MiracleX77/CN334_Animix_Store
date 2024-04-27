package usecases

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getDataFormAPI(data string) (string, error) {
	url := "https://api.aiforthai.in.th/ssense?text=" + data

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Apikey", "KIiC1yQGOukHvbLoAUTl7IjJ78RH7fzH")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	type Sentiment struct {
		Score       string `json:"score"`
		PolarityNeg bool   `json:"polarity-neg"`
		PolarityPos bool   `json:"polarity-pos"`
		Polarity    string `json:"polarity"`
	}

	type Preprocess struct {
		Input     string   `json:"input"`
		Neg       []string `json:"neg"`
		Pos       []string `json:"pos"`
		Segmented []string `json:"segmented"`
		Keyword   []string `json:"keyword"`
	}

	type Intention struct {
		Request      string `json:"request"`
		Sentiment    string `json:"sentiment"`
		Question     string `json:"question"`
		Announcement string `json:"announcement"`
	}

	type CustomAPIResponse struct {
		Sentiment   Sentiment  `json:"sentiment"`
		Preprocess  Preprocess `json:"preprocess"`
		Alert       []string   `json:"alert"`
		Comparative []string   `json:"comparative"`
		Associative []string   `json:"associative"`
		Intention   Intention  `json:"intention"`
	}
	var apiResponse CustomAPIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return "", err
	}

	return apiResponse.Sentiment.Polarity, nil
}
