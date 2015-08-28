package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	CUSTOM_FIELD_RES = "customfields"
)

type NewCustomField struct {
	Id   uint32 `json:"id"`
	Href string `json:"href"`
}

type CustomField struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type ContactCustomField struct {
	Id        uint32 `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Value     string `json:"value"`
}

type CustomFieldList struct {
	Page      uint32        `json:"page"`
	Limit     uint8         `json:"limit"`
	PageCount uint32        `json:"pageCount"`
	Resources []CustomField `json:"resources"`
}

/*
Get a single custom field.

    Parameters:

        id: Custom field id.
*/
func (client *TextmagicRestClient) GetCustomField(id uint32) (*CustomField, error) {
	customField := new(CustomField)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), CUSTOM_FIELD_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return customField, err
	}

	err = json.Unmarshal(response, customField)

	return customField, err
}

/*
Create a new custom field.

    Parameters:

        Var `data` may contain next keys:

            name: Name of custom field. Required.
*/
func (client *TextmagicRestClient) CreateCustomField(data map[string]string) (*NewCustomField, error) {
	customField := new(NewCustomField)

	method := "POST"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), CUSTOM_FIELD_RES)

	params := preparePostParams(data)
	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return customField, err
	}

	err = json.Unmarshal(response, customField)

	return customField, err
}

/*
Get all contact custom fields.

    Parameters:

        Var `data` may contain next keys:
            page:   Fetch specified results page.
            limit:  How many results on page.
*/
func (client *TextmagicRestClient) GetCustomFieldList(data map[string]string) (*CustomFieldList, error) {
	customFieldList := new(CustomFieldList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), CUSTOM_FIELD_RES)

	params := transformGetParams(data)
	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return customFieldList, err
	}

	err = json.Unmarshal(response, customFieldList)

	return customFieldList, err
}

/*
Updates the CustomField for the given unique id.

    Parameters:

        Var `data` may contain next keys:

            id:   The unique id of the CustomField to update. Required.
            name: Name of custom field. Required.
*/
func (client *TextmagicRestClient) UpdateCustomField(id uint32, data map[string]string) (*NewCustomField, error) {
	customField := new(NewCustomField)

	method := "PUT"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), CUSTOM_FIELD_RES, id)

	params := preparePostParams(data)
	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return customField, err
	}

	err = json.Unmarshal(response, customField)

	return customField, err
}

/*
Delete the specified CustomField from TextMagic.

    Parameters:

        id: The unique id of the CustomField to delete.
*/
func (client *TextmagicRestClient) DeleteCustomField(id uint32) (bool, error) {
	var success bool
	method := "DELETE"

	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), CUSTOM_FIELD_RES, id)

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
Updates contact's custom field value.

    Parameters:
        Var `data` may contain next keys:
            id:         The unique id of the CustomField to update a value. Required.
            contactId:  The unique id of the Contact to update value. Required.
            value:      Value of CustomField. Required.
*/
func (client *TextmagicRestClient) UpdateCustomFieldValue(id uint32, data map[string]string) (*NewContact, error) {
	contact := new(NewContact)

	method := "PUT"
	uri := fmt.Sprintf("%s/%s/%d/update", client.BaseUrl(), CUSTOM_FIELD_RES, id)

	params := preparePostParams(data)
	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return contact, err
	}

	err = json.Unmarshal(response, contact)

	return contact, err
}
