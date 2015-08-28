package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestMessages(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := time.Second
	client := textmagic.NewClient(username, token)

	newMessageText := "API GOLANG TEST"

	time.Sleep(interval)
	// Send message

	data := map[string]string{
		"text":   newMessageText,
		"phones": "999123345",
	}
	m, err := client.CreateMessage(data)

	assert.Nil(t, err)
	assert.NotEmpty(t, m.Id)
	assert.NotEmpty(t, m.Href)
	assert.NotEmpty(t, m.Type)
	assert.NotEmpty(t, m.SessionId)
	assert.Empty(t, m.BulkId)
	assert.NotEmpty(t, m.MessageId)
	assert.Empty(t, m.ScheduleId)
	assert.Equal(t, "message", m.Type)

	time.Sleep(interval)
	// Get Session List

	sessions, err := client.GetSessionList(map[string]string{})

	assert.Nil(t, err)
	assert.NotEmpty(t, sessions.Page)
	assert.NotEmpty(t, sessions.Limit)
	assert.NotEmpty(t, sessions.PageCount)
	assert.NotEmpty(t, sessions.Resources[0].Id)
	assert.NotEmpty(t, sessions.Resources[0].Text)
	assert.NotEmpty(t, sessions.Resources[0].StartTime)
	assert.NotEmpty(t, sessions.Resources[0].Source)
	assert.NotEmpty(t, sessions.Resources[0].ReferenceId)
	assert.NotEmpty(t, sessions.Resources[0].NumbersCount)

	time.Sleep(interval)
	// Get session by id

	session, _ := client.GetSession(sessions.Resources[0].Id)

	assert.Equal(t, sessions.Resources[0].Id, session.Id)
	assert.Equal(t, sessions.Resources[0].Text, session.Text)
	time.Sleep(interval)
	// Get session's messages

	sessionMessages, err := client.GetSessionMessages(session.Id, map[string]string{})

	assert.Nil(t, err)
	assert.NotEmpty(t, sessionMessages.Page)
	assert.NotEmpty(t, sessionMessages.Limit)
	assert.NotEmpty(t, sessionMessages.PageCount)
	assert.Equal(t, newMessageText, sessionMessages.Resources[0].Text)

	time.Sleep(interval)
	// Get Bulk Session List

	bulks, err := client.GetBulkSessionList(map[string]string{})

	assert.Nil(t, err)
	assert.NotEmpty(t, bulks.Page)
	assert.NotEmpty(t, bulks.Limit)
	assert.NotEmpty(t, bulks.PageCount)
	assert.NotEmpty(t, bulks.Resources[0].Id)
	assert.Equal(t, uint32(1), bulks.Page)
	assert.Equal(t, uint8(10), bulks.Limit)

	time.Sleep(interval)
	// Get Bulk Session by id

	bulk, err := client.GetBulkSession(bulks.Resources[0].Id)

	assert.Nil(t, err)
	assert.NotEmpty(t, bulk.Id)
	assert.NotEmpty(t, bulk.Status)
	assert.NotEmpty(t, bulk.ItemsProcessed)
	assert.NotEmpty(t, bulk.ItemsTotal)
	assert.NotEmpty(t, bulk.CreatedAt)
	assert.NotEmpty(t, bulk.Text)

	time.Sleep(interval)
	// Get chats

	chatList, err := client.GetChatList(map[string]string{})

	assert.Nil(t, err)
	assert.NotEmpty(t, chatList.Page)
	assert.NotEmpty(t, chatList.Limit)
	assert.NotEmpty(t, chatList.PageCount)
	assert.NotEqual(t, len(chatList.Resources), 0)

	chat := chatList.Resources[0]

	assert.NotEmpty(t, chat.Id)
	assert.NotEmpty(t, chat.Phone)
	assert.NotEmpty(t, chat.UpdatedAt)

	time.Sleep(interval)
	// Get Chat messages

	chatMessageList, err := client.GetChatMessageList(chat.Phone, map[string]string{})

	assert.Nil(t, err)
	assert.NotEmpty(t, chatMessageList.Page)
	assert.NotEmpty(t, chatMessageList.Limit)
	assert.NotEmpty(t, chatMessageList.PageCount)
	assert.NotEqual(t, len(chatMessageList.Resources), 0)

	chatMessage := chatMessageList.Resources[0]

	assert.NotEmpty(t, chatMessage.Id)
	assert.NotEmpty(t, chatMessage.Sender)
	assert.NotEmpty(t, chatMessage.MessageTime)
	assert.NotEmpty(t, chatMessage.Text)
	assert.NotEmpty(t, chatMessage.Receiver)
	assert.NotEmpty(t, chatMessage.Status)
	assert.NotEmpty(t, chatMessage.FirstName)
	assert.NotEmpty(t, chatMessage.LastName)

	time.Sleep(interval)
	// Get messages list

	listMessages, err := client.GetMessageList(map[string]string{}, false)

	assert.Nil(t, err)
	assert.NotEmpty(t, listMessages.Page)
	assert.NotEmpty(t, listMessages.Limit)
	assert.NotEmpty(t, listMessages.PageCount)
	assert.NotEqual(t, len(listMessages.Resources), 0)

	time.Sleep(interval)
	// Get messages price

	messagePriceData := map[string]string{
		"text":   "Go Api Test",
		"phones": "19025555555",
	}

	price, err := client.GetMessagePrice(messagePriceData)

	assert.Nil(t, err)
	assert.NotEmpty(t, price.Total)
	assert.NotEmpty(t, price.Parts)
	assert.NotEmpty(t, price.Countries)

	time.Sleep(interval)
	// Get message by id

	message, err := client.GetMessage(m.Id)

	assert.Nil(t, err)
	assert.NotEmpty(t, message.Id)
	assert.NotEmpty(t, message.Receiver)
	assert.NotEmpty(t, message.MessageTime)
	assert.NotEmpty(t, message.Status)
	assert.NotEmpty(t, message.Text)
	assert.NotEmpty(t, message.Charset)
	assert.NotEmpty(t, message.Country)
	assert.NotEmpty(t, message.Sender)
	assert.NotEmpty(t, message.PartsCount)

	time.Sleep(interval)
	// Delete single message

	r, err := client.DeleteMessage(m.Id)

	assert.Nil(t, err)
	assert.True(t, r)

	time.Sleep(interval)
	// Delete session

	r, err = client.DeleteSession(session.Id)

	assert.Nil(t, err)
	assert.True(t, r)

	time.Sleep(interval)
	// Get deleted message

	del, err := client.GetMessage(m.Id)

	t.Log(del)
	assert.NotNil(t, err)

	time.Sleep(interval)
	// Get deleted session

	delSess, err := client.GetSession(session.Id)

	t.Log(delSess)
	assert.NotNil(t, err)

	time.Sleep(interval)
	// Get inbox messages

	replies, _ := client.GetReplyList(map[string]string{}, false)

	assert.NotEmpty(t, replies.Page)
	assert.NotEmpty(t, replies.Limit)
	assert.NotEmpty(t, replies.PageCount)

	time.Sleep(interval)
	// Get inbox message by id

	reply, _ := client.GetReply(replies.Resources[0].Id)

	assert.Equal(t, replies.Resources[0].Text, reply.Text)

	time.Sleep(interval)
	// Create scheduled message

	scheduledData := map[string]string{
		"text":        "Scheduled Go Test",
		"phones":      "99900000",
		"sendingTime": strconv.Itoa(int(uint32(time.Now().Unix()) + 7200)),
	}
	scheduledNew, err := client.CreateMessage(scheduledData)

	assert.Nil(t, err)
	assert.NotEmpty(t, scheduledNew.Id)
	assert.NotEmpty(t, scheduledNew.Href)
	assert.NotEmpty(t, scheduledNew.Type)
	assert.Empty(t, scheduledNew.SessionId)
	assert.Empty(t, scheduledNew.BulkId)
	assert.Empty(t, scheduledNew.MessageId)
	assert.NotEmpty(t, scheduledNew.ScheduleId)
	assert.Equal(t, "schedule", scheduledNew.Type)

	time.Sleep(interval)
	// Get scheduled message by id

	scheduled, _ := client.GetScheduled(scheduledNew.Id)

	assert.Equal(t, scheduledNew.Id, scheduled.Id)
	assert.NotEmpty(t, scheduled.NextSend)
	assert.NotEmpty(t, scheduled.Session)
	assert.NotEmpty(t, scheduled.Session.Id)
	assert.NotEmpty(t, scheduled.Session.StartTime)
	assert.Equal(t, "Scheduled Go Test", scheduled.Session.Text)
	assert.Equal(t, "A", scheduled.Session.Source)
	assert.NotEmpty(t, scheduled.Session.ReferenceId)
	assert.Equal(t, float32(0), scheduled.Session.Price)
	assert.Equal(t, uint32(1), scheduled.Session.NumbersCount)

	time.Sleep(interval)
	// Get scheduled list

	listScheduled, _ := client.GetScheduledList(map[string]string{})

	assert.NotEmpty(t, listScheduled.Page)
	assert.NotEmpty(t, listScheduled.Limit)
	assert.NotEmpty(t, listScheduled.PageCount)

	time.Sleep(interval)
	// Delete scheduled message

	r, _ = client.DeleteScheduled(scheduledNew.Id)

	assert.True(t, r)

}
