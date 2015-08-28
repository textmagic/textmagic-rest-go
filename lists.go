package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	LIST_RES = "lists"
)

type NewList struct {
	Id   uint32 `json:"id"`
	Href string `json:"href"`
}

type List struct {
	Id           uint32 `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MembersCount uint32 `json:"membersCount"`
	Shared       bool   `json:"shared"`
}

type ListList struct {
	Page      uint32 `json:"page"`
	Limit     uint8  `json:"limit"`
	PageCount uint32 `json:"pageCount"`
	Resources []List `json:"resources"`
}

/*
Get a single list.

    Parameters:

        id: List id.
*/
func (client *TextmagicRestClient) GetList(id uint32) (*List, error) {
	list := new(List)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), LIST_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(response, list)

	return list, err
}

/*
Create a new list.

Parameters:

	Var `data` may contain next keys:

		name:        List name. Required.
		description: List description.
		shared:      Should this list be shared with sub-accounts. Can be 1 or 0.
*/
func (client *TextmagicRestClient) CreateList(data map[string]string) (*NewList, error) {
	list := new(NewList)

	method := "POST"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), LIST_RES)

	params := preparePostParams(data)
	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(response, list)

	return list, err
}

/*
Get all user lists.

	Parameters:

		Var `data` may contain next keys:

			page:   Fetch specified results page.
			limit:  How many results on page.
			ids:    Find lists by ID(s). Using with `search`=true.
			query:  Find lists by specified search query. Using with `search`=true.

		search: If true then search lists using `ids` and/or `query`.
*/
func (client *TextmagicRestClient) GetListList(data map[string]string, search bool) (*ListList, error) {
	listList := new(ListList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), LIST_RES)
	if search {
		uri += "/search"
	}

	params := transformGetParams(data)
	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return listList, err
	}

	err = json.Unmarshal(response, listList)

	return listList, err
}

/*
Updates the List for the given unique id.

Parameters:

	Var `data` may contain next keys:

		name:        List name. Required.
		description: List description.
		shared:      Should this list be shared with sub-accounts. Can be 1 or 0. Default=0.

	id: The unique id of the List to update. Required.
*/
func (client *TextmagicRestClient) UpdateList(id uint32, data map[string]string) (*NewList, error) {
	list := new(NewList)

	method := "PUT"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), LIST_RES, id)

	params := preparePostParams(data)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(response, list)

	return list, err
}

/*
Delete the specified List from TextMagic.

	Parameters:

		id: The unique id of the List to delete. Required.
*/
func (client *TextmagicRestClient) DeleteList(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), LIST_RES, id)

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
Fetch user contacts by given group id.
An useful synonym for "contacts/search" command with provided `groupId` parameter.

Parameters:

	id: The unique id of the List. Required.

	Var `data` may contain next keys:

		page:  Fetch specified results page.
		limit: How many results on page.
*/
func (client *TextmagicRestClient) GetContactsInList(id uint32, data map[string]string) (*ContactList, error) {
	contactList := new(ContactList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d/contacts", client.BaseUrl(), LIST_RES, id)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return contactList, err
	}

	err = json.Unmarshal(response, contactList)

	return contactList, err
}

/*
Assign contacts to the specified list.

	Parameters:

		id:       The unique id of the List. Required.
		contacts: Contact ID(s), separated by comma. Required.
*/
func (client *TextmagicRestClient) PutContactsIntoList(id uint32, contacts string) (*NewList, error) {
	list := new(NewList)

	method := "PUT"
	data := map[string]string{"contacts": contacts}
	params := preparePostParams(data)

	uri := fmt.Sprintf("%s/%s/%d/contacts", client.BaseUrl(), LIST_RES, id)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return list, err
	}

	err = json.Unmarshal(response, list)

	return list, err
}

/*
Unassign contacts from the specified list.
If contacts assign only to the specified list, then delete permanently.

	Parameters:
		id:      The unique id of the List. Required.
		contacts: Contact ID(s), separated by comma. Required.
*/
func (client *TextmagicRestClient) DeleteContactsFromList(id uint32, contacts string) (bool, error) {
	var success bool

	method := "DELETE"
	data := map[string]string{"contacts": contacts}
	params := preparePostParams(data)

	uri := fmt.Sprintf("%s/%s/%d/contacts", client.BaseUrl(), LIST_RES, id)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return false, err
	}
	if response[0] == 204 {
		success = true
	}

	return success, err
}
