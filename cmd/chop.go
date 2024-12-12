package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func createNewURL(ogUrl string) (result string) {
	data := make(map[string]string)
	data["url"] = ogUrl
	jsonString, err := json.Marshal(data)

	if err != nil {
		return fmt.Sprintln("Some error occurred:", err.Error())
	}

	res, err := http.Post(
		"https://hr8.fun/url/create",
		"application/json",
		bytes.NewBuffer(jsonString),
	)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Sprintf("Error reading response body: %v", err)
	}

	type RespBody struct {
		Data string `json:"data"`
	}
	var dataObj RespBody
	err = json.Unmarshal(body, &dataObj)

	if err != nil {
		return fmt.Sprintln("Some error occurred:", err.Error())
	}

	return fmt.Sprintf(dataObj.Data)
}
