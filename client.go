package textmagic

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const BASE_URL = "https://rest.textmagic.com/api/v2"

type TextmagicRestClient struct {
	username string
	token    string
	baseUrl  string
}

func NewClient(username, token string) *TextmagicRestClient {
	baseUrl := BASE_URL
	return &TextmagicRestClient{username, token, baseUrl}
}

func (client *TextmagicRestClient) Username() string {
	return client.username
}

func (client *TextmagicRestClient) Token() string {
	return client.token
}

func (client *TextmagicRestClient) BaseUrl() string {
	return client.baseUrl
}

func preparePostParams(params map[string]string) url.Values {
	prepared := url.Values{}
	for key, value := range params {
		prepared.Set(key, value)
	}

	return prepared
}

func transformGetParams(params map[string]string) url.Values {
	transformedParams := url.Values{}
	for key, value := range params {
		transformedParams.Set(key, value)
	}

	return transformedParams
}

func (client *TextmagicRestClient) Request(method string, uri string, params url.Values, data url.Values) ([]byte, error) {
	var requestData = strings.NewReader(url.Values{}.Encode())

	if data != nil {
		requestData = strings.NewReader(data.Encode())
	}
	if params != nil {
		uri += fmt.Sprintf("?%s", params.Encode())
	}

	request, err := http.NewRequest(method, uri, requestData)

	if err != nil {
		return nil, err
	}

	if method == "POST" || method == "PUT" || method == "DELETE" {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	request.Header.Add("Accept-Charset", "utf-8")
	request.Header.Add("Accept-Language", "en-us")

	// To avoid Header.Add key capitalization.
	request.Header["X-TM-Username"] = []string{client.Username()}
	request.Header["X-TM-Key"] = []string{client.Token()}

	httpClient := &http.Client{}

	response, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return body, err
	}

	if response.StatusCode != 200 && response.StatusCode != 201 && response.StatusCode != 204 {
		textmagicError := new(TextmagicError)

		json.Unmarshal(body, textmagicError)

		return body, textmagicError
	}

	// Convert response StatusCode to []byte to return
	if method == "DELETE" || response.StatusCode == 204 {
		var b = make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(response.StatusCode))

		return b, err
	}

	return body, err
}

type PingResp struct {
	Ping string `json:"ping"`
}

func (client *TextmagicRestClient) Ping() (string, error) {
	var ping string
	pingResp := new(PingResp)

	method := "GET"
	uri := fmt.Sprintf("%s/ping", client.BaseUrl())
	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return ping, err
	}

	err = json.Unmarshal(response, pingResp)

	return pingResp.Ping, err
}
