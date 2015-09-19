package textmagic

import "strconv"

const numberURI = "numbers"

// Country represents a country.
type Country struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// NewNumber represents a new number.
type NewNumber struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// Number represents a number.
type Number struct {
	ID          int      `json:"id"`
	User        *User    `json:"user"`
	PurchasedAt string   `json:"purchasedAt"`
	ExpireAt    string   `json:"expireAt"`
	Phone       string   `json:"phone"`
	Country     *Country `json:"country"`
	Status      string   `json:"status"`
}

// NumberList represents a number list.
type NumberList struct {
	Page      int       `json:"page"`
	Limit     int       `json:"limit"`
	PageCount int       `json:"pageCount"`
	Resources []*Number `json:"resources"`
}

// AvailableNumbers represents available numbers.
type AvailableNumbers struct {
	Numbers []string `json:"numbers"`
	Price   float64  `json:"price"`
}

// GetNumber gets a single dedicated number
// for the given ID.
func (c *Client) GetNumber(id int) (*Number, error) {
	var n *Number

	return n, c.get(numberURI+"/"+strconv.Itoa(id), nil, nil, &n)
}

// GetNumberList returns all user dedicated numbers.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetNumberList(p Params) (*NumberList, error) {
	var l *NumberList

	return l, c.get(numberURI, p, nil, &l)
}

// BuyNumber buys a dedicated number and assigns
// it to the specified account.
//
// The data payload includes:
// - phone:   Desired dedicated phone number in international E.164 format. Required.
// - country: Dedicated number country. Required.
// - userId:  User ID this number will be assigned to. Required.
func (c *Client) BuyNumber(d Params) (*NewNumber, error) {
	var n *NewNumber

	return n, c.post(numberURI, nil, d, &n)
}

// GetAvailableNumbers finds available dedicated
// numbers to buy.
//
// The parameter payload includes:
// - country: Dedicated number country. Required.
// - prefix:  Desired number prefix. Should include country code (i.e. 447 for GB)
func (c *Client) GetAvailableNumbers(p Params) (*AvailableNumbers, error) {
	var n *AvailableNumbers

	return n, c.get(numberURI+"/available", p, nil, &n)
}

// CancelNumber cancels the dedicated number
// subscription for the given ID.
func (c *Client) CancelNumber(id int) error {
	return c.delete(numberURI+"/"+strconv.Itoa(id), nil, nil, nil)
}
