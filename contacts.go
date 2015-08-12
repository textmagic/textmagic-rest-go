package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	CONTACT_RESOURCE = "contacts"
)

type NewContact struct {
	Id   uint32 `json:"id"`
	Href string `json:"href"`
}

type Contact struct {
	Id           uint32               `json:"id"`
	Phone        string               `json:"phone"`
	FirstName    string               `json:"firstName"`
	LastName     string               `json:"lastName"`
	Company      string               `json:"companyName"`
	Country      map[string]string    `json:"country"`
	CustomFields []ContactCustomField `json:"customFields"`
}

type ContactList struct {
	Page      uint32    `json:"page"`
	Limit     uint8     `json:"limit"`
	PageCount uint32    `json:"pageCount"`
	Contacts  []Contact `json:"resources"`
}

/*
Get a single contact.

    Parameters:

        id: Contact id.
*/
func (client *TextmagicRestClient) GetContact(id uint32) (*Contact, error) {
	contact := new(Contact)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), CONTACT_RESOURCE, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return contact, err
	}

	err = json.Unmarshal(response, contact)

	return contact, err
}

/*
Create a new contact.

    Parameters:

        Var `data` may contain next keys:

            firstName:
            lastName:
            phone:       Contact's phone number. Required.
            email:
            companyName:
            country:     2-letter ISO country code.
            lists:       String of Lists separated by commas to assign contact. Required.
*/
func (client *TextmagicRestClient) CreateContact(data map[string]string) (*NewContact, error) {
	var contact = new(NewContact)

	method := "POST"

	params := preparePostParams(data)
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), CONTACT_RESOURCE)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return contact, err
	}

	err = json.Unmarshal(response, contact)

	return contact, err
}

/*
Get all user contacts.

    Parameters:

        Var `data` may contain next keys:

            page:   Fetch specified results page.
            limit:  How many results on page.
            shared: Should shared contacts to be included.
            ids:    Find contact by ID(s). Using with `search`=true.
            listId: Find contact by List ID. Using with `search`=true.
            query:  Find contact by specified search query. Using with `search`=true.

        search:   If true then search contacts using `query`, `ids` and/or `group_id`.
*/
func (client *TextmagicRestClient) GetContactList(data map[string]string, search bool) (*ContactList, error) {
	contactList := new(ContactList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), CONTACT_RESOURCE)

	if search {
		uri += "/search"
	}

	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return contactList, err
	}

	err = json.Unmarshal(response, contactList)

	return contactList, err
}

/*
Updates the existing Contact for the given unique id.

    Parameters:

        id: The unique id of the Contact to update. Required.

        Var `data` may contain next keys:

            firstName:
            lastName:
            phone:       Contact's phone number. Required.
            email:
            companyName:
            lists:       String of Lists separated by commas to assign contact. Required.
*/
func (client *TextmagicRestClient) UpdateContact(id uint32, data map[string]string) (*NewContact, error) {
	contact := new(NewContact)

	method := "PUT"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), CONTACT_RESOURCE, id)

	params := preparePostParams(data)
	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return contact, err
	}

	err = json.Unmarshal(response, contact)

	return contact, err
}

/*
Delete the specified Contact from TextMagic.

    Parameters:
        id: The unique id of the Contact to delete.
*/
func (client *TextmagicRestClient) DeleteContact(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), CONTACT_RESOURCE, id)
	response, err := client.Request(method, uri, nil, nil)

	if err != nil {
		return success, err
	}
	if response[0] == 204 {
		success = true
	}

	return success, err
}

/*
Return lists which contact belongs to.

    Parameters:

        id:   The unique id of the Contact to update. Required.

        Var `data` may contain next keys:

            page:  Fetch specified results page. Default=1
            limit: How many results on page. Default=10
*/
func (client *TextmagicRestClient) GetContactLists(id uint32, data map[string]string) (*ListList, error) {
	listList := new(ListList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d/lists", client.BaseUrl(), CONTACT_RESOURCE, id)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return listList, err
	}

	err = json.Unmarshal(response, listList)

	return listList, err
}
