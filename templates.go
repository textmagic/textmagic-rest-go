package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	TEMPLATE_RES = "templates"
)

type NewTemplate struct {
	Id   uint32 `json:"id"`
	Href string `json:"href"`
}

type Template struct {
	Id           uint32 `json:"id"`
	Name         string `json:"name"`
	Content      string `json:"content"`
	LastModified string `json:"lastModified"`
}

type TemplateList struct {
	Page      uint32     `json:"page"`
	Limit     uint8      `json:"limit"`
	PageCount uint32     `json:"pageCount"`
	Templates []Template `json:"resources"`
}

/*
Get a single message template.

	Parameters:

		id: Message template id.
*/
func (client *TextmagicRestClient) GetTemplate(id uint32) (*Template, error) {
	template := new(Template)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), TEMPLATE_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return template, err
	}

	err = json.Unmarshal(response, template)

	return template, err
}

/*
Create a new template.

	Parameters:

		Var `data` may contain next keys:

			name:    Template name. Required.
			content: Template text. May contain tags inside braces. Required.
*/
func (client *TextmagicRestClient) CreateTemplate(data map[string]string) (*NewTemplate, error) {
	template := new(NewTemplate)

	method := "POST"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), TEMPLATE_RES)
	params := preparePostParams(data)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return template, err
	}

	err = json.Unmarshal(response, template)

	return template, err
}

/*
Get all user message templates.

	Parameters:

		Var `data` may contain next keys:

			page:     Fetch specified results page.
	        limit:    How many results on page.
	        name:     Find template by name. Using with `search`=true.
	        content:  Find template by content. Using with `search`=true.

		search: If true then search templates using `name` and/or `content`.

*/
func (client *TextmagicRestClient) GetTemplateList(data map[string]string, search bool) (*TemplateList, error) {
	templateList := new(TemplateList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), TEMPLATE_RES)
	if search {
		uri += "/search"
	}
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return templateList, err
	}

	err = json.Unmarshal(response, templateList)

	return templateList, err
}

/*
Updates the Template for the given unique id.

	Parameters:

		id: Unique id of the template to update. Required.

		Var `data` may contain next keys:

			name:    Template name. Required.
			content: Template text. May contain tags inside braces. Required.
*/
func (client *TextmagicRestClient) UpdateTemplate(id uint32, data map[string]string) (*NewTemplate, error) {
	template := new(NewTemplate)

	method := "PUT"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), TEMPLATE_RES, id)
	params := preparePostParams(data)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return template, err
	}

	err = json.Unmarshal(response, template)

	return template, err
}

/*
Delete the specified Template from TextMagic.

	Parameters:

		id: Template id.
*/
func (client *TextmagicRestClient) DeleteTemplate(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), TEMPLATE_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return false, err
	}
	if response[0] == 204 {
		success = true
	}

	return success, err
}
