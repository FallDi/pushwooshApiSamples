package main

import
(
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io/ioutil"
)

const (
	PW_APPLICATION = "EEA77-472F8"
	PW_AUTH = "bHmwnN4qBpgiHuvO9rWbxBDHfG6jjKFUTx5s69bIRVnpsTog6qTvPmafddGVsoW0TCxJUSYAyeaS3R29JBmJ"
	PW_ENDPOINT = "https://cp.pushwoosh.com/json/1.3/"
)

func pwCall(method string, data []byte) (bool) {
	url := PW_ENDPOINT + method
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error occur: " + err.Error())
		return false
	}
	defer response.Body.Close()

	fmt.Println("Response Status: ", response.Status)
	if (response.StatusCode == 200) {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Response Body: ", string(body))
		return true
	}
	return false
}

func main() {
	requestData := map[string]interface{}{
		"request": map[string]interface{} {
			"auth": PW_AUTH,
			"application": PW_APPLICATION,
			"notifications": []interface{}{
				map[string]interface{} {
					"send_date": "now",
					"content": "test",
					"link": "https://pushwoosh.com",
				},
			},
		},
	}
	jsonRequest, _ := json.Marshal(requestData)
	requestString := string(jsonRequest)
	fmt.Println("Request body: " + requestString)

	pwCall("createMessage", jsonRequest)
}
