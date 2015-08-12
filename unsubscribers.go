package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	UNSUBSCRIBER_RES = "unsubscribers"
)

type NewUnsubscriber struct {
	Id   uint32 `json:"id"`
	Href string `json:"href"`
}

type Unsubscriber struct {
	Id              uint32 `json:"id"`
	Phone           string `json:"phone"`
	UnsubscribeTime string `json:"unsubscribeTime"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

type UnsubscriberList struct {
	Page          uint32         `json:"page"`
	Limit         uint8          `json:"limit"`
	PageCount     uint32         `json:"pageCount"`
	Unsubscribers []Unsubscriber `json:"resources"`
}

/*
Get a single unsubscribed contact.

    Parameters:

        id: Unsubscribed contact id.
*/
func (client *TextmagicRestClient) GetUnsubscriber(id uint32) (*Unsubscriber, error) {
	unsubscriber := new(Unsubscriber)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), UNSUBSCRIBER_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return unsubscriber, err
	}

	err = json.Unmarshal(response, unsubscriber)

	return unsubscriber, err
}

/*
Unsubscribe contact from your communication by phone number.

    Parameters:

        phone: Contact phone number you want to unsubscribe. Required.
*/
func (client *TextmagicRestClient) UnsubscribePhone(phone string) (*NewUnsubscriber, error) {
	unsubscriber := new(NewUnsubscriber)

	method := "POST"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), UNSUBSCRIBER_RES)

	data := map[string]string{"phone": phone}
	params := preparePostParams(data)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return unsubscriber, err
	}

	err = json.Unmarshal(response, unsubscriber)

	return unsubscriber, err
}

/*
Get all contact have unsubscribed from your communication.

Parameters:

    Var `data` may contain next keys:

        page:  Fetch specified results page.
        limit: How many results on page.
*/
func (client *TextmagicRestClient) GetUnsubscriberList(data map[string]string) (*UnsubscriberList, error) {
	unsubscriberList := new(UnsubscriberList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), UNSUBSCRIBER_RES)

	params := transformGetParams(data)
	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return unsubscriberList, err
	}

	err = json.Unmarshal(response, unsubscriberList)

	return unsubscriberList, err
}
