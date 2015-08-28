package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	NUMBER_RES = "numbers"
)

type Country struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type NewNumber struct {
	Id   uint32 `json:"id"`
	Href string `json:"href"`
}

type Number struct {
	Id          uint32  `json:"id"`
	User        User    `json:"user"`
	PurchasedAt string  `json:"purchasedAt"`
	ExpireAt    string  `json:"expireAt"`
	Phone       string  `json:"phone"`
	Country     Country `json:"country"`
	Status      string  `json:"status"`
}

type NumberList struct {
	Page      uint32   `json:"page"`
	Limit     uint8    `json:"limit"`
	PageCount uint32   `json:"pageCount"`
	Resources []Number `json:"resources"`
}

type AvailableNumbers struct {
	Numbers []string `json:"numbers"`
	Price   float32  `json:"price"`
}

/*
Get a single dedicated number.

	Parameters:

		id: Dedicated number id.
*/
func (client *TextmagicRestClient) GetNumber(id uint32) (*Number, error) {
	var number = new(Number)

	method := "GET"

	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), NUMBER_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return number, err
	}

	err = json.Unmarshal(response, number)

	return number, err
}

/*
Get all user dedicated numbers.

	Parameters:

		Var `data` may contain next keys:

			page:  Fetch specified results page.
			limit: How many results on page.
*/
func (client *TextmagicRestClient) GetNumberList(data map[string]string) (*NumberList, error) {
	var numberList = new(NumberList)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), NUMBER_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return numberList, err
	}

	err = json.Unmarshal(response, numberList)

	return numberList, err
}

/*
Buy a dedicated number and assign it to the specified account.

	Parameters:

		Var `data` may contain next keys:

			phone:   Desired dedicated phone number in international E.164 format. Required.
	        country: Dedicated number country. Required.
	        userId:  User ID this number will be assigned to. Required.
*/
func (client *TextmagicRestClient) BuyNumber(data map[string]string) (*NewNumber, error) {
	var number = new(NewNumber)

	method := "POST"

	params := preparePostParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), NUMBER_RES)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return number, err
	}

	err = json.Unmarshal(response, number)

	return number, err
}

/*
Find available dedicated numbers to buy.

	Parameters:

		Var `data` may contain next keys:

			country: Dedicated number country. Required.
        	prefix:  Desired number prefix. Should include country code (i.e. 447 for GB)
*/
func (client *TextmagicRestClient) GetAvailableNumbers(data map[string]string) (*AvailableNumbers, error) {
	var numbers = new(AvailableNumbers)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s/available", client.BaseUrl(), NUMBER_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return numbers, err
	}

	err = json.Unmarshal(response, numbers)

	return numbers, err
}

/*
Cancel dedicated number subscription.

	Parameters:

		id: The unique id of the Number. Required.
*/
func (client *TextmagicRestClient) CancelNumber(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), NUMBER_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return false, err
	}
	if response[0] == 204 {
		success = true
	}

	return success, err
}
