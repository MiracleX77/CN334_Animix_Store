package usecases

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getAiDataFormAPI(data string) (string, error) {
	url := "https://api.aiforthai.in.th/ssense"

	payload := strings.NewReader("text=" + data)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Apikey", "KIiC1yQGOukHvbLoAUTl7IjJ78RH7fzH")
	req.Header.Add("Host", "ssense")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "PostmanRuntime/7.11.0")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("content-length", "284")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	reader := bytes.NewReader(body)
	gzipReader, err := gzip.NewReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	defer gzipReader.Close()

	decompressedData, err := ioutil.ReadAll(gzipReader)
	if err != nil {
		log.Fatal(err)
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

	type Associative struct {
		EntPos      []string `json:"ent-pos"`
		PolarityNeg bool     `json:"polarity-neg"`
		EndIndex    int      `json:"endIndex"`
		PolarityPos bool     `json:"polarity-pos"`
		BeginIndex  int      `json:"beginIndex"`
		Text        string   `json:"text"`
		EntNeg      []string `json:"ent-neg"`
		Asp         []string `json:"asp"`
	}

	type CustomAPIResponse struct {
		Sentiment   Sentiment     `json:"sentiment"`
		Preprocess  Preprocess    `json:"preprocess"`
		Alert       []string      `json:"alert"`
		Comparative []string      `json:"comparative"`
		Associative []Associative `json:"associative"`
		Intention   Intention     `json:"intention"`
	}

	var apiResponse CustomAPIResponse
	if err := json.Unmarshal(decompressedData, &apiResponse); err != nil {
		return "", err
	}
	return apiResponse.Sentiment.Polarity, nil
}

func getDataFormAPI(port string, typeData string, Id string, out interface{}, token string) error {
	url := "http://" + typeData + "-service:" + port + "/v1/" + typeData + "/" + Id

	// Create the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	type APIResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	var apiResponse APIResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return err
	}
	fmt.Println(apiResponse.Data)
	// Convert the data to the struct
	data, err := json.Marshal(apiResponse.Data)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, out); err != nil {
		return err
	}

	return nil
}
