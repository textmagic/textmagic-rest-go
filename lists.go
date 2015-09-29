package textmagic

import "strconv"

const listURI = "lists"

// NewList represents a new list.
type NewList struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// List represents a list item.
type List struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	MembersCount int    `json:"membersCount"`
	Shared       bool   `json:"shared"`
}

// Lists represents lists and pagination information.
type Lists struct {
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	PageCount int     `json:"pageCount"`
	Resources []*List `json:"resources"`
}

// GetList returns the list with the given ID.
func (c *Client) GetList(id int) (*List, error) {
	var l *List

	return l, c.get(listURI+"/"+strconv.Itoa(id), nil, nil, &l)
}

// CreateList creates a new list with
// the corresponding POST DATA.
//
// The data payload includes:
// - description: List description.
// - shared:      Should this list be shared with sub-accounts. Can be 1 or 0.
func (c *Client) CreateList(d Params) (*NewList, error) {
	var l *NewList

	return l, c.post(listURI, nil, d, &l)
}

// GetLists returns all user lists.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetLists(p Params) (*Lists, error) {
	var l *Lists

	return l, c.get(listURI, p, nil, &l)
}

// SearchLists returns all user lists for the given search.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
// - ids:	Find lists by ID(s).
// - query:	Find lists by specified search query.
func (c *Client) SearchLists(p Params) (*Lists, error) {
	var l *Lists

	return l, c.get(listURI+"/search", p, nil, &l)
}

// UpdateList updates the list for the given ID.
//
// The data payload includes:
// - name:        List name. Required.
// - description: List description.
// - shared:      Should this list be shared with sub-accounts. Can be 1 or 0. Default = 0.
func (c *Client) UpdateList(id int, d Params) (*NewList, error) {
	var l *NewList

	return l, c.put(listURI+"/"+strconv.Itoa(id), nil, d, &l)
}

// DeleteList deletes the list with the given ID.
func (c *Client) DeleteList(id int) error {
	return c.delete(listURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// GetContactsInList fetches the contacts for
// the given list ID.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetContactsInList(id int, p Params) (*ContactList, error) {
	var l *ContactList

	return l, c.get(listURI+"/"+strconv.Itoa(id)+"/contacts", p, nil, &l)
}

// PutContactsIntoList assigns comma separated contacts
// string to the list with the given ID.
func (c *Client) PutContactsIntoList(id int, contacts ...int) (*NewList, error) {
	var l *NewList

	return l, c.put(listURI+"/"+strconv.Itoa(id)+"/contacts", nil, NewParams("contacts", contacts), &l)
}

// DeleteContactsFromList deletes contacts from the given list.
func (c *Client) DeleteContactsFromList(id int, contacts ...int) error {
	return c.delete(listURI+"/"+strconv.Itoa(id)+"/contacts", nil, NewParams("contacts", contacts), nil)
}
