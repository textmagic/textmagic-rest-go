package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	SENDERID_RES = "senderids"
	SOURCE_RES   = "sources"
)

type NewSenderId struct {
	Id   uint32 `json:"id"`
	Href string `json:"href"`
}

type SenderId struct {
	Id       uint32 `json:"id"`
	SenderId string `json:"senderId"`
	User     User   `json:"user"`
	Status   string `json:"status"`
}

type SenderIdList struct {
	Page      uint32     `json:"page"`
	Limit     uint8      `json:"limit"`
	PageCount uint32     `json:"pageCount"`
	Resources []SenderId `json:"resources"`
}

type Sources struct {
	Dedicated []string `json:"dedicated"`
	User      []string `json:"user"`
	Shared    []string `json:"shared"`
	SenderIds []string `json:"senderIds"`
}

/*
Get a single Sender Id.

	Parameters:

		id: Sender id (numeric, not alphanumeric Sender ID itself).
*/
func (client *TextmagicRestClient) GetSenderId(id uint32) (*SenderId, error) {
	var senderid = new(SenderId)

	method := "GET"

	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SENDERID_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return senderid, err
	}

	err = json.Unmarshal(response, senderid)

	return senderid, err
}

/*
Get all user sender ids.

	Parameters:

		Var `data` may contain next keys:

			page:  Fetch specified results page.
			limit: How many results on page.
*/
func (client *TextmagicRestClient) GetSenderIdList(data map[string]string) (*SenderIdList, error) {
	var senderIdList = new(SenderIdList)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), SENDERID_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return senderIdList, err
	}

	err = json.Unmarshal(response, senderIdList)

	return senderIdList, err
}

/*
Apply for a new Sender ID.

	Parameters:

		Var `data` may contain next keys:

			senderId:    Alphanumeric Sender ID (maximum 11 characters). Required.
        	explanation: Explain why do you need this Sender ID. Required.
*/
func (client *TextmagicRestClient) CreateSenderId(data map[string]string) (*NewSenderId, error) {
	var senderId = new(NewSenderId)

	method := "POST"

	params := preparePostParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), SENDERID_RES)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return senderId, err
	}

	err = json.Unmarshal(response, senderId)

	return senderId, err
}

/*
Delete the specified Sender ID from TextMagic.

	Parameters:

		id: The unique id of the Sender ID to delete (numeric, not alphanumeric Sender ID itself). Required.
*/
func (client *TextmagicRestClient) DeleteSenderId(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SENDERID_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return false, err
	}
	if response[0] == 204 {
		success = true
	}

	return success, err
}

/*
Get all available sender settings which could be used in "from" parameter of POST messages method.

	Parameters:

		country:  Return sender settings available in specified country only. Optional.
*/
func (client *TextmagicRestClient) GetSources(data map[string]string) (*Sources, error) {
	var sources = new(Sources)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), SOURCE_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return sources, err
	}

	err = json.Unmarshal(response, sources)

	return sources, err
}
