package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	MESSAGE_RES   = "messages"
	BULK_RES      = "bulks"
	CHAT_RES      = "chats"
	REPLY_RES     = "replies"
	SCHEDULED_RES = "schedules"
	SESSION_RES   = "sessions"
)

type NewMessage struct {
	Id         uint32 `json:"id"`
	Href       string `json:"href"`
	Type       string `json:"type"`
	SessionId  uint32 `json:"sessionId"`
	BulkId     uint32 `json:"bulkId"`
	MessageId  uint32 `json:"messageId"`
	ScheduleId uint32 `json:"scheduleId"`
}

type Message struct {
	Id          uint32  `json:"id"`
	Receiver    string  `json:"receiver"`
	MessageTime string  `json:"messageTime"`
	Status      string  `json:"status"`
	Text        string  `json:"text"`
	Charset     string  `json:"charset"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Country     string  `json:"country"`
	Sender      string  `json:"sender"`
	Price       float32 `json:"price"`
	PartsCount  uint8   `json:"partsCount"`
}

type MessageList struct {
	Page      uint32    `json:"page"`
	Limit     uint8     `json:"limit"`
	PageCount uint32    `json:"pageCount"`
	Resources []Message `json:"resources"`
}

type Session struct {
	Id           uint32  `json:"id"`
	StartTime    string  `json:"startTime"`
	Text         string  `json:"text"`
	Source       string  `json:"source"`
	ReferenceId  string  `json:"referenceId"`
	Price        float32 `json:"price"`
	NumbersCount uint32  `json:"numbersCount"`
}

type SessionList struct {
	Page      uint32    `json:"page"`
	Limit     uint8     `json:"limit"`
	PageCount uint32    `json:"pageCount"`
	Resources []Session `json:"resources"`
}

type BulkSession struct {
	Id             uint32  `json:"id"`
	Status         string  `json:"status"`
	ItemsProcessed uint32  `json:"itemsProcessed"`
	ItemsTotal     uint32  `json:"itemsTotal"`
	CreatedAt      string  `json:"createdAt"`
	Session        Session `json:"session"`
	Text           string  `json:"text"`
}

type BulkSessionList struct {
	Page      uint32        `json:"page"`
	Limit     uint8         `json:"limit"`
	PageCount uint32        `json:"pageCount"`
	Resources []BulkSession `json:"resources"`
}

type Chat struct {
	Id        uint32  `json:"id"`
	Phone     string  `json:"phone"`
	Contact   Contact `json:"contact"`
	Unread    uint32  `json:"unread"`
	UpdatedAt string  `json:"updatedAt"`
}

type ChatList struct {
	Page      uint32 `json:"page"`
	Limit     uint8  `json:"limit"`
	PageCount uint32 `json:"pageCount"`
	Resources []Chat `json:"resources"`
}

type ChatMessage struct {
	Id          uint32 `json:"id"`
	Sender      string `json:"sender"`
	MessageTime string `json:"messageTime"`
	Text        string `json:"text"`
	Receiver    string `json:"receiver"`
	Status      string `json:"status"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
}

type ChatMessageList struct {
	Page      uint32        `json:"page"`
	Limit     uint8         `json:"limit"`
	PageCount uint32        `json:"pageCount"`
	Resources []ChatMessage `json:"resources"`
}

type CountryPrice struct {
	Country string  `json:"country"`
	Count   uint32  `json:"count"`
	Max     float32 `json:"max"`
}

type MessagePrice struct {
	Total     float32                 `json:"total"`
	Parts     uint8                   `json:"parts"`
	Countries map[string]CountryPrice `json:"countries"`
}

type Reply struct {
	Id          uint32 `json:"id"`
	Sender      string `json:"sender"`
	MessageTime string `json:"messageTime"`
	Text        string `json:"text"`
	Receiver    string `json:"receiver"`
}

type ReplyList struct {
	Page      uint32  `json:"page"`
	Limit     uint8   `json:"limit"`
	PageCount uint32  `json:"pageCount"`
	Resources []Reply `json:"resources"`
}

type Scheduled struct {
	Id       uint32  `json:"id"`
	NextSend string  `json:"nextSend"`
	Rrule    string  `json:"rrule"`
	Session  Session `json:"session"`
}

type ScheduledList struct {
	Page      uint32      `json:"page"`
	Limit     uint8       `json:"limit"`
	PageCount uint32      `json:"pageCount"`
	Resources []Scheduled `json:"resources"`
}

/*
Create and send a new outbound message.

    Parameters:

        Var `data` may contain next keys:

			text:         Message text. Required if template_id is not set.
			templateId:   Template used instead of message text. Required if text is not set.
			sendingTime:  Message sending time in unix timestamp format. Default is now. Optional (required with recurrency_rule set).
			contacts:     Contacts ids, separated by comma, message will be sent to.
			lists:        Lists ids, separated by comma, message will be sent to.
			phones:       Phone numbers, separated by comma, message will be sent to.
			cutExtra:     Should sending method cut extra characters which not fit supplied parts_count or return 400 Bad request response instead.
			partsCount:   Maximum message parts count (TextMagic allows sending 1 to 6 message parts).
			referenceId:  Custom message reference id which can be used in your application infrastructure.
			from:        One of allowed Sender ID (phone number or alphanumeric sender ID).
			rrule:        iCal RRULE parameter to create recurrent scheduled messages. When used, sending_time is mandatory as start point of sending.
*/
func (client *TextmagicRestClient) CreateMessage(data map[string]string) (*NewMessage, error) {
	var message = new(NewMessage)

	method := "POST"

	params := preparePostParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), MESSAGE_RES)

	response, err := client.Request(method, uri, nil, params)
	if err != nil {
		return message, err
	}

	err = json.Unmarshal(response, message)

	return message, err
}

/*
Get a single outgoing message.

    Parameters:

        id: Message id.
*/
func (client *TextmagicRestClient) GetMessage(id uint32) (*Message, error) {
	var message = new(Message)

	method := "GET"

	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), MESSAGE_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return message, err
	}

	err = json.Unmarshal(response, message)

	return message, err
}

/*
Get all user oubound messages.

    Parameters:

        Var `data` may contain next keys:

            page:   	Fetch specified results page.
            limit:  	How many results on page.
	        ids:        Find message by ID(s). Using with `search`=true.
	        sessionId:  Find messages by session ID. Using with `search`=true.
	        query:      Find messages by specified search query. Using with `search`=true.

        search: If true then search messages using `ids`, `sessionId`, and/or `query`.
*/
func (client *TextmagicRestClient) GetMessageList(data map[string]string, search bool) (*MessageList, error) {
	var messageList = new(MessageList)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), MESSAGE_RES)
	if search {
		uri += "/search"
	}

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return messageList, err
	}

	err = json.Unmarshal(response, messageList)

	return messageList, err
}

/*
Get bulk message session.

    Parameters:

        id: Bulk session id.
*/
func (client *TextmagicRestClient) GetBulkSession(id uint32) (*BulkSession, error) {
	bulk := new(BulkSession)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), BULK_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return bulk, err
	}

	err = json.Unmarshal(response, bulk)

	return bulk, err
}

/*
Get all bulk sending sessions.

	Parameters:

		Var `data` may contain next keys:

			page:  Fetch specified results page.
			limit: How many results on page.
*/
func (client *TextmagicRestClient) GetBulkSessionList(data map[string]string) (*BulkSessionList, error) {
	var bulkSessionList = new(BulkSessionList)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), BULK_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return bulkSessionList, err
	}

	err = json.Unmarshal(response, bulkSessionList)

	return bulkSessionList, err
}

/*
Fetch messages from chat with specified phone number.

	Parameters:

		phone: Phone number in E.164 format.

		Var `data` may contain next keys:

	        page:  Fetch specified results page.
	        limit: How many results on page.
*/
func (client *TextmagicRestClient) GetChatMessageList(phone string, data map[string]string) (*ChatMessageList, error) {
	chatMessageList := new(ChatMessageList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%s", client.BaseUrl(), CHAT_RES, phone)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return chatMessageList, err
	}

	err = json.Unmarshal(response, chatMessageList)

	return chatMessageList, err
}

/*
Get all user chats.

	Parameters:

		Var `data` may contain next keys:

	        page:  Fetch specified results page.
	        limit: How many results on page.
*/
func (client *TextmagicRestClient) GetChatList(data map[string]string) (*ChatList, error) {
	chatList := new(ChatList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), CHAT_RES)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return chatList, err
	}

	err = json.Unmarshal(response, chatList)

	return chatList, err
}

/*
Check pricing for a new outbound message.

    Parameters:

        Var `data` may contain next keys:

			text:         Message text. Required if template_id is not set.
			templateId:   Template used instead of message text. Required if text is not set.
			sendingTime:  Message sending time in unix timestamp format. Default is now. Optional (required with recurrency_rule set).
			contacts:     Contacts ids, separated by comma, message will be sent to.
			lists:        Lists ids, separated by comma, message will be sent to.
			phones:       Phone numbers, separated by comma, message will be sent to.
			cutExtra:     Should sending method cut extra characters which not fit supplied parts_count or return 400 Bad request response instead.
			partsCount:   Maximum message parts count (TextMagic allows sending 1 to 6 message parts).
			referenceId:  Custom message reference id which can be used in your application infrastructure.
			from:         One of allowed Sender ID (phone number or alphanumeric sender ID).
			rrule:        iCal RRULE parameter to create recurrent scheduled messages. When used, sending_time is mandatory as start point of sending.
*/
func (client *TextmagicRestClient) GetMessagePrice(data map[string]string) (*MessagePrice, error) {
	price := new(MessagePrice)

	method := "GET"
	params := transformGetParams(data)
	uri := fmt.Sprintf("%s/%s/price", client.BaseUrl(), MESSAGE_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return price, err
	}

	err = json.Unmarshal(response, price)

	return price, err
}

/*
Delete the specified Message from TextMagic.

	Parameters:

		id: The unique id of the Message. Required.
*/
func (client *TextmagicRestClient) DeleteMessage(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), MESSAGE_RES, id)

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
Get a single inbound message

	Parameters:

		id: Inbound message id.
*/
func (client *TextmagicRestClient) GetReply(id uint32) (*Reply, error) {
	reply := new(Reply)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), REPLY_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return reply, err
	}

	err = json.Unmarshal(response, reply)

	return reply, err
}

/*
Get all user inbound messages.

    Parameters:

        Var `data` may contain next keys:

            page:   	Fetch specified results page.
            limit:  	How many results on page.
	        ids:        Find replies by ID(s). Using with `search`=true.
	        query:      Find replies by specified search query. Using with `search`=true.

        search: If true then search messages using `ids` and/or `query`.
*/
func (client *TextmagicRestClient) GetReplyList(data map[string]string, search bool) (*ReplyList, error) {
	replyList := new(ReplyList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), REPLY_RES)
	if search {
		uri += "/search"
	}
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return replyList, err
	}

	err = json.Unmarshal(response, replyList)

	return replyList, err
}

/*
Delete the specified Reply from TextMagic.

	Parameters:

		id: The unique id of the Reply. Required.
*/
func (client *TextmagicRestClient) DeleteReply(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), REPLY_RES, id)

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
Get a single scheduled message.

	Parameters:

		id: Scheduled message id.
*/
func (client *TextmagicRestClient) GetScheduled(id uint32) (*Scheduled, error) {
	scheduled := new(Scheduled)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SCHEDULED_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return scheduled, err
	}

	err = json.Unmarshal(response, scheduled)

	return scheduled, err
}

/*
Get all user scheduled messages.

    Parameters:

        Var `data` may contain next keys:

            page:   	Fetch specified results page.
            limit:  	How many results on page.
*/
func (client *TextmagicRestClient) GetScheduledList(data map[string]string) (*ScheduledList, error) {
	scheduledList := new(ScheduledList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), SCHEDULED_RES)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return scheduledList, err
	}

	err = json.Unmarshal(response, scheduledList)

	return scheduledList, err
}

/*
Delete the specified Scheduled Message from TextMagic.

	Parameters:

		id: The unique id of the ScheduledMessage. Required.
*/
func (client *TextmagicRestClient) DeleteScheduled(id uint32) (bool, error) {
	var success bool

	method := "DELETE"

	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SCHEDULED_RES, id)

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
Get a single message session.

	Parameters:

		id: Message session id
*/
func (client *TextmagicRestClient) GetSession(id uint32) (*Session, error) {
	session := new(Session)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SESSION_RES, id)

	response, err := client.Request(method, uri, nil, nil)
	if err != nil {
		return session, err
	}

	err = json.Unmarshal(response, session)

	return session, err
}

/*
Get all user message session.

	Parameters:

		Var `data` may contains next keys:

			page:   	Fetch specified results page.
            limit:  	How many results on page.
*/
func (client *TextmagicRestClient) GetSessionList(data map[string]string) (*SessionList, error) {
	sessions := new(SessionList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), SESSION_RES)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return sessions, err
	}

	err = json.Unmarshal(response, sessions)

	return sessions, err
}

/*
Delete the specified Message Session from TextMagic.

	Parameters:

		id: The unique id of the Session. Required.
*/
func (client *TextmagicRestClient) DeleteSession(id uint32) (bool, error) {
	var success bool

	method := "DELETE"
	uri := fmt.Sprintf("%s/%s/%d", client.BaseUrl(), SESSION_RES, id)

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
Fetch messages by given session id.
An useful synonym for "messages/search" command with provided `sessionId` parameter.

	Parameters:

		id:   The unique id of the Session. Required.

		Var `data` may contains next keys:

			page:  Fetch specified results page.
			limit: How many results on page.
*/
func (client *TextmagicRestClient) GetSessionMessages(id uint32, data map[string]string) (*MessageList, error) {
	messageList := new(MessageList)

	method := "GET"
	uri := fmt.Sprintf("%s/%s/%d/messages", client.BaseUrl(), SESSION_RES, id)
	params := transformGetParams(data)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return messageList, err
	}

	err = json.Unmarshal(response, messageList)

	return messageList, err
}
