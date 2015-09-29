package textmagic

import "strconv"

const unsubscriberURI = "unsubscribers"

// NewUnsubscriber represents a new unsubscriber.
type NewUnsubscriber struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// Unsubscriber represents a unsubscriber.
type Unsubscriber struct {
	ID              int    `json:"id"`
	Phone           string `json:"phone"`
	UnsubscribeTime string `json:"unsubscribeTime"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
}

// UnsubscriberList represents a unsubscriber list.
type UnsubscriberList struct {
	Page      int             `json:"page"`
	Limit     int             `json:"limit"`
	PageCount int             `json:"pageCount"`
	Resources []*Unsubscriber `json:"resources"`
}

// GetUnsubscriber returns an unsubscribed contact by ID.
func (c *Client) GetUnsubscriber(id int) (*Unsubscriber, error) {
	var u *Unsubscriber

	return u, c.get(unsubscriberURI+"/"+strconv.Itoa(id), nil, nil, &u)
}

// UnsubscribePhone unsubscribes a contact by phone number.
func (c *Client) UnsubscribePhone(phone string) (*NewUnsubscriber, error) {
	var u *NewUnsubscriber

	return u, c.post(unsubscriberURI, nil, NewParams("phone", phone), &u)
}

// GetUnsubscriberList returns all contacts that
// have unsubscribed from your communication.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetUnsubscriberList(p Params) (*UnsubscriberList, error) {
	var l *UnsubscriberList

	return l, c.get(unsubscriberURI, p, nil, &l)
}
