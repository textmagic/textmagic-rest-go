package textmagic

import "strconv"

const contactURI = "contacts"

// NewContact represents a new contact object.
type NewContact struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// Contact represents a contact object.
type Contact struct {
	ID           int                   `json:"id"`
	Phone        string                `json:"phone"`
	FirstName    string                `json:"firstName"`
	LastName     string                `json:"lastName"`
	Company      string                `json:"companyName"`
	Country      map[string]string     `json:"country"`
	CustomFields []*ContactCustomField `json:"customFields"`
}

// ContactList represents a contact list object,
// including pagination and statistical information.
type ContactList struct {
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
	PageCount int        `json:"pageCount"`
	Resources []*Contact `json:"resources"`
}

// GetContact returns a single contact by ID.
func (c *Client) GetContact(id int) (*Contact, error) {
	var contact *Contact

	return contact, c.get(contactURI+"/"+strconv.Itoa(id), nil, nil, &contact)
}

// CreateContact creates a new contact with
// the corresponding POST DATA.
//
// The data payload includes:
// - firstName:
// - lastName:
// - phone:         Contact's phone number. Required.
// - email:
// - companyName:
// - country:       2-letter ISO country code.
// - lists:         String of Lists separated by commas to assign contact. Required.
func (c *Client) CreateContact(d Params) (*NewContact, error) {
	var contact *NewContact

	return contact, c.post(contactURI, nil, d, &contact)
}

// GetContactList returns the contact list.
//
// The parameter payload includes:
// - page:		Fetch specified results page.
// - limit:     How many results on page.
// - shared:    Should shared contacts to be included.
func (c *Client) GetContactList(p Params) (*ContactList, error) {
	var l *ContactList

	return l, c.get(contactURI, p, nil, &l)
}

// SearchContactList returns a contact list in relation
// to search filters.
//
// The parameter payload includes:
// - page:      Fetch specified results page.
// - limit:     How many results on page.
// - shared:	Should shared contacts to be included.
// - ids:       Find contact by ID(s).
// - listId:    Find contact by List ID.
// - query:     Find contact by specified search query.
func (c *Client) SearchContactList(p Params) (*ContactList, error) {
	var l *ContactList

	return l, c.get(contactURI+"/search", p, nil, &l)
}

// UpdateContact updates an existing contact with
// the corresponding POST DATA.
//
// The data payload includes:
// - firstName:
// - lastName:
// - phone:         Contact's phone number. Required.
// - email:
// - companyName:
// - country:		2-letter ISO country code.
// - lists:         String of Lists separated by commas to assign contact. Required.
func (c *Client) UpdateContact(id int, d Params) (*NewContact, error) {
	var contact *NewContact

	return contact, c.put(contactURI+"/"+strconv.Itoa(id), nil, d, &contact)
}

// DeleteContact deletes the contact with
// the given ID.
func (c *Client) DeleteContact(id int) error {
	return c.delete(contactURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// GetContactLists returns the lists the given
// contact belongs to.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetContactLists(id int, p Params) (*Lists, error) {
	var l *Lists

	return l, c.get(contactURI+"/"+strconv.Itoa(id)+"/lists", p, nil, &l)
}
