package usecases

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getDataFormAPI(port string, typeService string, typeData string, Id string, out interface{}, token string) error {
	url := "http://" + typeService + "-service:" + port + "/v1/" + typeData + "/" + Id

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
