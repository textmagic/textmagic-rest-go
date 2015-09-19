package textmagic

import "strconv"

const (
	messageURI   = "messages"
	bulkURI      = "bulks"
	chatURI      = "chats"
	replyURI     = "replies"
	scheduledURI = "schedules"
	sessionURI   = "sessions"
)

// NewMessage represents a new message.
type NewMessage struct {
	ID         int    `json:"id"`
	Href       string `json:"href"`
	Type       string `json:"type"`
	SessionID  int    `json:"sessionId"`
	BulkID     int    `json:"bulkId"`
	MessageID  int    `json:"messageId"`
	ScheduleID int    `json:"scheduleId"`
}

// Message represents a message.
type Message struct {
	ID          int     `json:"id"`
	Receiver    string  `json:"receiver"`
	MessageTime string  `json:"messageTime"`
	Status      string  `json:"status"`
	Text        string  `json:"text"`
	Charset     string  `json:"charset"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Country     string  `json:"country"`
	Sender      string  `json:"sender"`
	Price       float64 `json:"price"`
	PartsCount  int     `json:"partsCount"`
}

// MessageList represents a message list.
type MessageList struct {
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
	PageCount int        `json:"pageCount"`
	Resources []*Message `json:"resources"`
}

// Session represents a session.
type Session struct {
	ID           int     `json:"id"`
	StartTime    string  `json:"startTime"`
	Text         string  `json:"text"`
	Source       string  `json:"source"`
	ReferenceID  string  `json:"referenceId"`
	Price        float64 `json:"price"`
	NumbersCount int     `json:"numbersCount"`
}

// SessionList represents a session list.
type SessionList struct {
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
	PageCount int        `json:"pageCount"`
	Resources []*Session `json:"resources"`
}

// BulkSession represents a bulk session.
type BulkSession struct {
	ID             int      `json:"id"`
	Status         string   `json:"status"`
	ItemsProcessed int      `json:"itemsProcessed"`
	ItemsTotal     int      `json:"itemsTotal"`
	CreatedAt      string   `json:"createdAt"`
	Session        *Session `json:"session"`
	Text           string   `json:"text"`
}

// BulkSessionList represents a bulk session list.
type BulkSessionList struct {
	Page      int            `json:"page"`
	Limit     int            `json:"limit"`
	PageCount int            `json:"pageCount"`
	Resources []*BulkSession `json:"resources"`
}

// Chat represents a chat item.
type Chat struct {
	ID        int      `json:"id"`
	Phone     string   `json:"phone"`
	Contact   *Contact `json:"contact"`
	Unread    int      `json:"unread"`
	UpdatedAt string   `json:"updatedAt"`
}

// ChatList represents a chat item list.
type ChatList struct {
	Page      int     `json:"page"`
	Limit     int     `json:"limit"`
	PageCount int     `json:"pageCount"`
	Resources []*Chat `json:"resources"`
}

// ChatMessage represents a chat message.
type ChatMessage struct {
	ID          int    `json:"id"`
	Sender      string `json:"sender"`
	MessageTime string `json:"messageTime"`
	Text        string `json:"text"`
	Receiver    string `json:"receiver"`
	Status      string `json:"status"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
}

// ChatMessageList represents a chat message list.
type ChatMessageList struct {
	Page      int            `json:"page"`
	Limit     int            `json:"limit"`
	PageCount int            `json:"pageCount"`
	Resources []*ChatMessage `json:"resources"`
}

// CountryPrice represents a country price.
type CountryPrice struct {
	Country string  `json:"country"`
	Count   int     `json:"count"`
	Max     float64 `json:"max"`
}

// MessagePrice represents a message price.
type MessagePrice struct {
	Total     float64                  `json:"total"`
	Parts     int                      `json:"parts"`
	Countries map[string]*CountryPrice `json:"countries"`
}

// Reply represents a reply.
type Reply struct {
	ID          int    `json:"id"`
	Sender      string `json:"sender"`
	MessageTime string `json:"messageTime"`
	Text        string `json:"text"`
	Receiver    string `json:"receiver"`
}

// ReplyList represents a reply list.
type ReplyList struct {
	Page      int      `json:"page"`
	Limit     int      `json:"limit"`
	PageCount int      `json:"pageCount"`
	Resources []*Reply `json:"resources"`
}

// Scheduled represents a scheduled item.
type Scheduled struct {
	ID       int      `json:"id"`
	NextSend string   `json:"nextSend"`
	Rrule    string   `json:"rrule"`
	Session  *Session `json:"session"`
}

// ScheduledList represents a scheduled item list.
type ScheduledList struct {
	Page      int          `json:"page"`
	Limit     int          `json:"limit"`
	PageCount int          `json:"pageCount"`
	Resources []*Scheduled `json:"resources"`
}

// CreateMessage creates and sends a new outbound
// message with the corresponding POST DATA.
//
// The data payload includes:
// - text:			Message text. Required if template_id is not set.
// - templateId:	Template used instead of message text. Required if text is not set.
// - sendingTime:	Message sending time in Unix timestamp format. Default is now. Optional (required with recurrency_rule set).
// - contacts:		Contacts ids, separated by comma, message will be sent to.
// - lists:			Lists ids, separated by comma, message will be sent to.
// - phones:		Phone numbers, separated by comma, message will be sent to.
// - cutExtra:		Should sending method cut extra characters which not fit supplied parts_count or return 400 Bad request response instead.
// - partsCount:	Maximum message parts count (TextMagic allows sending 1 to 6 message parts).
// - referenceId:	Custom message reference id which can be used in your application infrastructure.
// - from:			One of allowed Sender ID (phone number or alphanumeric sender ID).
// - rrule:			iCal RRULE parameter to create recurrent scheduled messages. When used, sending_time is mandatory as start point of sending.
func (c *Client) CreateMessage(d Params) (*NewMessage, error) {
	var m *NewMessage

	return m, c.post(messageURI, nil, d, &m)
}

// GetMessage returns a single outgoing message by ID.
func (c *Client) GetMessage(id int) (*Message, error) {
	var m *Message

	return m, c.get(messageURI+"/"+strconv.Itoa(id), nil, nil, &m)
}

// GetMessageList returns all user outbound messages.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetMessageList(p Params, search bool) (*MessageList, error) {
	var l *MessageList

	return l, c.get(messageURI, p, nil, &l)
}

// SearchMessageList returns all user outbound messages
// for the given search
//
// The parameter payload includes:
// - page:		Fetch specified results page.
// - limit:		How many results on page.
// - ids:		Find message by ID(s).
// - sessionId:	Find messages by session ID.
// - query:		Find messages by specified search query.
func (c *Client) SearchMessageList(p Params) (*MessageList, error) {
	var l *MessageList

	return l, c.get(messageURI+"/search", p, nil, &l)
}

// GetBulkSession returns the bulk message
// session by ID.
func (c *Client) GetBulkSession(id int) (*BulkSession, error) {
	var b *BulkSession

	return b, c.get(bulkURI, nil, nil, &b)
}

// GetBulkSessionList returns all bulk sending sessions.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetBulkSessionList(p Params) (*BulkSessionList, error) {
	var l *BulkSessionList

	return l, c.get(bulkURI, p, nil, &l)
}

// GetChatMessageList returns all messages from
// chat with specified phone number.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetChatMessageList(phone string, p Params) (*ChatMessageList, error) {
	var l *ChatMessageList

	return l, c.get(chatURI, p, nil, &l)
}

// GetChatList returns all user chats.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetChatList(p Params) (*ChatList, error) {
	var l *ChatList

	return l, c.get(chatURI, p, nil, &l)
}

// GetMessagePrice checks pricing for a
// new outbound message.
//
// The parameter payload includes:
// - text:         	Message text. Required if template_id is not set.
// - templateId:   	Template used instead of message text. Required if text is not set.
// - sendingTime:	Message sending time in Unix timestamp format. Default is now. Optional (required with recurrency_rule set).
// - contacts:		Contacts ids, separated by comma, message will be sent to.
// - lists:			Lists ids, separated by comma, message will be sent to.
// - phones:		Phone numbers, separated by comma, message will be sent to.
// - cutExtra:     	Should sending method cut extra characters which not fit supplied parts_count or return 400 Bad request response instead.
// - partsCount:   	Maximum message parts count (TextMagic allows sending 1 to 6 message parts).
// - referenceId:	Custom message reference id which can be used in your application infrastructure.
// - from:         	One of allowed Sender ID (phone number or alphanumeric sender ID).
// - rrule:        	iCal RRULE parameter to create recurrent scheduled messages. When used, sending_time is mandatory as start point of sending.
func (c *Client) GetMessagePrice(p Params) (*MessagePrice, error) {
	var m *MessagePrice

	return m, c.get(messageURI+"/price", p, nil, &m)
}

// DeleteMessage deletes the message with the given ID.
func (c *Client) DeleteMessage(id int) error {
	return c.delete(messageURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// GetReply returns a single inbound message by ID.
func (c *Client) GetReply(id int) (*Reply, error) {
	var r *Reply

	return r, c.get(replyURI+"/"+strconv.Itoa(id), nil, nil, &r)
}

// GetReplyList returns all user inbound messages.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetReplyList(p Params, search bool) (*ReplyList, error) {
	var l *ReplyList

	return l, c.get(replyURI, p, nil, &l)
}

// SearchReplyList returns all user chats.
//
// The parameter payload includes:
// - page:		Fetch specified results page.
// - limit:		How many results on page.
// - ids:		Find replies by ID(s).
// - query:		Find replies by specified search query.
func (c *Client) SearchReplyList(p Params) (*ReplyList, error) {
	var l *ReplyList

	return l, c.get(replyURI+"/search", p, nil, &l)
}

// DeleteReply deletes the reply with the given ID.
func (c *Client) DeleteReply(id int) error {
	return c.delete(replyURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// GetScheduled returns the single scheduled item
// with the given ID.
func (c *Client) GetScheduled(id int) (*Scheduled, error) {
	var s *Scheduled

	return s, c.get(scheduledURI+"/"+strconv.Itoa(id), nil, nil, &s)
}

// GetScheduledList returns all user scheduled messages.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetScheduledList(p Params) (*ScheduledList, error) {
	var l *ScheduledList

	return l, c.get(scheduledURI, p, nil, &l)
}

// DeleteScheduled deletes the scheduled message
// with the given ID.
func (c *Client) DeleteScheduled(id int) error {
	return c.delete(scheduledURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// GetSession returns the single session
// with the given ID.
func (c *Client) GetSession(id int) (*Session, error) {
	var s *Session

	return s, c.get(sessionURI+"/"+strconv.Itoa(id), nil, nil, &s)
}

// GetSessionList returns all user message sessions.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetSessionList(p Params) (*SessionList, error) {
	var l *SessionList

	return l, c.get(sessionURI, p, nil, &l)
}

// DeleteSession deletes the message session with
// the given ID.
func (c *Client) DeleteSession(id int) error {
	return c.delete(sessionURI+"/"+strconv.Itoa(id), nil, nil, nil)
}

// GetSessionMessages fetches messages bound by
// the given session ID.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetSessionMessages(id int, p Params) (*MessageList, error) {
	var l *MessageList

	return l, c.get(sessionURI+"/"+strconv.Itoa(id)+"/messages", p, nil, &l)
}
