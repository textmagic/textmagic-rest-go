package textmagic

import "strconv"

const (
	senderIDURI = "senderids"
	sourceURI   = "sources"
)

// NewSenderID represents a new sender ID.
type NewSenderID struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}

// SenderID represents a sender ID.
type SenderID struct {
	ID       int    `json:"id"`
	SenderID string `json:"senderId"`
	User     *User  `json:"user"`
	Status   string `json:"status"`
}

// SenderIDList  represents a sender ID list.
type SenderIDList struct {
	Page      int         `json:"page"`
	Limit     int         `json:"limit"`
	PageCount int         `json:"pageCount"`
	Resources []*SenderID `json:"resources"`
}

// Sources represents sources.
type Sources struct {
	Dedicated []string `json:"dedicated"`
	User      []string `json:"user"`
	Shared    []string `json:"shared"`
	SenderIDs []string `json:"senderIds"`
}

// GetSenderID returns a single sender ID
// for the give numeric ID.
func (c *Client) GetSenderID(id int) (*SenderID, error) {
	var s *SenderID

	return s, c.get(senderIDURI+"/"+strconv.Itoa(id), nil, nil, &s)
}

// GetSenderIDList returns all user sender IDs.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetSenderIDList(p Params) (*SenderIDList, error) {
	var l *SenderIDList

	return l, c.get(senderIDURI, p, nil, &l)
}

// CreateSenderID creates a new sender ID.
//
// The data payload includes:
// - senderId:		Alphanumeric Sender ID (maximum 11 characters). Required.
// - explanation:	Explain why do you need this Sender ID. Required.
func (c *Client) CreateSenderID(d Params) (*NewSenderID, error) {
	var s *NewSenderID

	return s, c.post(senderIDURI, nil, d, &s)
}

// DeleteSenderID deletes the given sender ID.
func (c *Client) DeleteSenderID(id int) error {
	return c.delete(senderIDURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// GetSources returns all available sender settings which
// could be used in "from" parameter of POST messages method.
//
// The parameter payload includes:
// - country:	Return sender settings available in specified country only. Optional.
func (c *Client) GetSources(p Params) (*Sources, error) {
	var s *Sources

	return s, c.get(sourceURI, p, nil, &s)
}
