package textmagic

import "strconv"

const customFieldURI = "customfields"

// NewCustomField represents a new custom field.
type NewCustomField struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// CustomField represents a custom field.
type CustomField struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

// ContactCustomField represents a contact custom field.
type ContactCustomField struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	Value     string `json:"value"`
}

// CustomFieldList represents a custom field list.
type CustomFieldList struct {
	Page      int            `json:"page"`
	Limit     int            `json:"limit"`
	PageCount int            `json:"pageCount"`
	Resources []*CustomField `json:"resources"`
}

// GetCustomField returns a single custom field by ID.
func (c *Client) GetCustomField(id int) (*CustomField, error) {
	var f *CustomField

	return f, c.get(customFieldURI+"/"+strconv.Itoa(id), nil, nil, &f)
}

// CreateCustomField creates a new custom field
// with the given name.
func (c *Client) CreateCustomField(name string) (*NewCustomField, error) {
	var f *NewCustomField

	return f, c.post(customFieldURI, nil, NewParams("name", name), &f)
}

// GetCustomFieldList returns the custom field list.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetCustomFieldList(p Params) (*CustomFieldList, error) {
	var l *CustomFieldList

	return l, c.get(customFieldURI, p, nil, &l)
}

// UpdateCustomField updates the given custom field
// to the provided name value.
func (c *Client) UpdateCustomField(id int, name string) (*NewCustomField, error) {
	var f *NewCustomField

	return f, c.put(customFieldURI+"/"+strconv.Itoa(id), nil, NewParams("name", name), &f)
}

// DeleteCustomField deletes the custom field
// with the given ID.
func (c *Client) DeleteCustomField(id int) error {
	return c.delete(customFieldURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// UpdateCustomFieldValue updates the contact's
// custom field value.
//
// The parameter payload includes:
// - contactId:	The unique id of the Contact to update value. Required.
// - value:     Value of CustomField. Required.
func (c *Client) UpdateCustomFieldValue(id int, d Params) (*NewContact, error) {
	var contact *NewContact

	return contact, c.put(customFieldURI+"/"+strconv.Itoa(id)+"/update", nil, d, &contact)
}
