package textmagic

import "strconv"

const (
	statURI       = "stats"
	tokenURI      = "tokens"
	userURI       = "user"
	subAccountURI = "subaccounts"
)

// MessagingStat represents messaging statistics.
type MessagingStat struct {
	ReplyRate             float64 `json:"replyRate"`
	Date                  string  `json:"date"`
	DeliveryRate          float64 `json:"deliveryRate"`
	Costs                 float64 `json:"costs"`
	MessagesReceived      int     `json:"messagesReceived"`
	MessagesSentDelivered int     `json:"messagesSentDelivered"`
	MessagesSentAccepted  int     `json:"messagesSentAccepted"`
	MessagesSentBuffered  int     `json:"messagesSentBuffered"`
	MessagesSentFailed    int     `json:"messagesSentFailed"`
	MessagesSentRejected  int     `json:"messagesSentRejected"`
	MessagesSentParts     int     `json:"messagesSentParts"`
}

// SpendingStat represents spending statistics.
type SpendingStat struct {
	ID      int     `json:"id"`
	UserID  int     `json:"userId"`
	Date    string  `json:"date"`
	Balance float64 `json:"balance"`
	Delta   float64 `json:"delta"`
	Type    string  `json:"type"`
	Value   string  `json:"value"`
	Comment string  `json:"comment"`
}

// SpendingStatList represents a spending statistics list.
type SpendingStatList struct {
	Page      int             `json:"page"`
	Limit     int             `json:"limit"`
	PageCount int             `json:"pageCount"`
	Resources []*SpendingStat `json:"resources"`
}

// NewToken represents a new token.
type NewToken struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	Expires  string `json:"string"`
}

// Currency represents a currency
type Currency struct {
	ID         string `json:"id"`
	HTMLSymbol string `json:"htmlSymbol"`
}

// Timezone represents a timezone.
type Timezone struct {
	ID       int    `json:"id"`
	Area     string `json:"area"`
	Dst      int    `json:"dst"`
	Offset   int    `json:"offset"`
	Timezone string `json:"timezone"`
}

// User represents a user.
type User struct {
	ID             int       `json:"id"`
	Username       string    `json:"username"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Status         string    `json:"status"`
	Balance        float64   `json:"balance"`
	Company        string    `json:"company"`
	Currency       *Currency `json:"currency"`
	Timezone       *Timezone `json:"timezone"`
	SubaccountType string    `json:"subaccountType"`
}

// UserList represents a user list.
type UserList struct {
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	PageCount int     `json:"pageCount"`
	Resources []*User `json:"resources"`
}

// GetMessagingStat returns messaging statistics.
//
// The parameter payload includes:
// - by:    Group results by specified period: `off`, `day`, `month` or `year`. Default is `off`.
// - start:	Start date in Unix timestamp format. Default is 7 days ago.
// - end:	End date in Unix timestamp format. Default is now.
func (c *Client) GetMessagingStat(p Params) ([]*MessagingStat, error) {
	var l []*MessagingStat

	return l, c.get(statURI+"/messaging", p, nil, &l)
}

// GetSpendingStat returns account spending statistics.
//
// The parameter payload includes:
// - page:  Fetch specified results page. Default=1
// - limit: How many results on page. Default=10
// - start: Start date in Unix timestamp format. Default is 7 days ago.
// - end:   End date in Unix timestamp format. Default is now.
func (c *Client) GetSpendingStat(p Params) (*SpendingStatList, error) {
	var s *SpendingStatList

	return s, c.get(statURI+"/spending", p, nil, &s)
}

// GetUser returns the current user.
func (c *Client) GetUser() (*User, error) {
	var u *User

	return u, c.get(userURI, nil, nil, &u)
}

// UpdateUser updates the current user.
//
// The data payload includes:
// - firstName: User first name. Required.
// - lastName:  User last name. Required.
// - company:   User company. Required.
func (c *Client) UpdateUser(d Params) (map[string]string, error) {
	result := make(map[string]string)

	return result, c.put(userURI, nil, d, &result)
}

// GetSubaccount gets a subaccount by the given ID.
func (c *Client) GetSubaccount(id int) (*User, error) {
	var u *User

	return u, c.get(subAccountURI+"/"+strconv.Itoa(id), nil, nil, &u)
}

// GetSubaccountList returns all user subaccounts.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetSubaccountList(p Params) (*UserList, error) {
	var l *UserList

	return l, c.get(subAccountURI, p, nil, &l)
}

// SendInvite sends an invite for a new subaccount.
//
// The data payload includes:
// - email: Subaccount email. Required.
// - role:  Subaccount role: `A` for administrator or `U` for regular user. Required.
func (c *Client) SendInvite(d Params) error {
	return c.post(subAccountURI, nil, d, nil)
}

// CloseSubaccount closes the subaccount for the given ID.
func (c *Client) CloseSubaccount(id int) error {
	return c.delete(subAccountURI+"/"+strconv.Itoa(id), nil, nil, nil)
}
