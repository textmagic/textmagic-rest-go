package textmagic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"strconv"
)

func TestMessages(t *testing.T) {
	newMessageText := "API GOLANG TEST"

	time.Sleep(interval)
	// Send message

	data := Params{
		"text":   newMessageText,
		"phones": "999123345",
	}

	m, err := client.CreateMessage(data)

	assert.Nil(t, err)
	assert.NotEmpty(t, m.ID)
	assert.NotEmpty(t, m.Href)
	assert.NotEmpty(t, m.Type)
	assert.NotEmpty(t, m.SessionID)
	assert.Empty(t, m.BulkID)
	assert.NotEmpty(t, m.MessageID)
	assert.Empty(t, m.ScheduleID)
	assert.Equal(t, "message", m.Type)

	time.Sleep(interval)
	// Get Session List

	sessions, err := client.GetSessionList(nil)

	assert.Nil(t, err)
	assert.NotEmpty(t, sessions.Page)
	assert.NotEmpty(t, sessions.Limit)
	assert.NotEmpty(t, sessions.PageCount)
	assert.NotEmpty(t, sessions.Resources[0].ID)
	assert.NotEmpty(t, sessions.Resources[0].Text)
	assert.NotEmpty(t, sessions.Resources[0].StartTime)
	assert.NotEmpty(t, sessions.Resources[0].Source)
	assert.NotEmpty(t, sessions.Resources[0].ReferenceID)
	assert.NotEmpty(t, sessions.Resources[0].NumbersCount)

	time.Sleep(interval)
	// Get session by id

	session, _ := client.GetSession(sessions.Resources[0].ID)

	assert.Equal(t, sessions.Resources[0].ID, session.ID)
	assert.Equal(t, sessions.Resources[0].Text, session.Text)
	time.Sleep(interval)
	// Get session's messages

	sessionMessages, err := client.GetSessionMessages(session.ID, nil)

	assert.Nil(t, err)
	assert.NotEmpty(t, sessionMessages.Page)
	assert.NotEmpty(t, sessionMessages.Limit)
	assert.NotEmpty(t, sessionMessages.PageCount)
	assert.Equal(t, newMessageText, sessionMessages.Resources[0].Text)

	time.Sleep(interval)
	// Get Bulk Session List

	bulks, err := client.GetBulkSessionList(nil)

	assert.Nil(t, err)
	assert.NotNil(t, bulks)

	time.Sleep(interval)

	// Get chats
	chatList, err := client.GetChatList(nil)

	assert.Nil(t, err)
	assert.NotNil(t, chatList)

	time.Sleep(interval)
	// Get messages list

	listMessages, err := client.GetMessageList(nil, false)

	assert.Nil(t, err)
	assert.NotNil(t, listMessages)

	time.Sleep(interval)
	// Get messages price

	messagePriceData := Params{
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

	message, err := client.GetMessage(m.ID)

	assert.Nil(t, err)
	assert.NotEmpty(t, message.ID)
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

	err = client.DeleteMessage(m.ID)

	assert.Nil(t, err)

	time.Sleep(interval)
	// Delete session

	err = client.DeleteSession(session.ID)

	assert.Nil(t, err)

	time.Sleep(interval)
	// Get deleted message

	del, err := client.GetMessage(m.ID)

	t.Log(del)
	assert.NotNil(t, err)

	time.Sleep(interval)
	// Get deleted session

	delSess, err := client.GetSession(session.ID)

	t.Log(delSess)
	assert.NotNil(t, err)

	time.Sleep(interval)
	// Get inbox messages

	replies, err := client.GetReplyList(nil, false)

	assert.Nil(t, err)
	assert.NotNil(t, replies)

	time.Sleep(interval)

	scheduledData := Params{
		"text":   "Scheduled Go Test",
		"phones": "99900000",
	}
	scheduledData.Set("sendingTime", strconv.FormatInt(time.Now().Unix()+7200, 10))

	scheduledNew, err := client.CreateMessage(scheduledData)

	assert.Nil(t, err)
	assert.NotEmpty(t, scheduledNew.ID)
	assert.NotEmpty(t, scheduledNew.Href)
	assert.NotEmpty(t, scheduledNew.Type)
	assert.Empty(t, scheduledNew.SessionID)
	assert.Empty(t, scheduledNew.BulkID)
	assert.Empty(t, scheduledNew.MessageID)
	assert.NotEmpty(t, scheduledNew.ScheduleID)
	assert.Equal(t, "schedule", scheduledNew.Type)

	time.Sleep(interval)
	// Get scheduled message by id

	scheduled, _ := client.GetScheduled(scheduledNew.ID)

	assert.Equal(t, scheduledNew.ID, scheduled.ID)
	assert.NotEmpty(t, scheduled.NextSend)
	assert.NotEmpty(t, scheduled.Session)
	assert.NotEmpty(t, scheduled.Session.ID)
	assert.NotEmpty(t, scheduled.Session.StartTime)
	assert.Equal(t, "Scheduled Go Test", scheduled.Session.Text)
	assert.Equal(t, "A", scheduled.Session.Source)
	assert.NotEmpty(t, scheduled.Session.ReferenceID)
	assert.Equal(t, 0.0, scheduled.Session.Price)
	assert.Equal(t, 1, scheduled.Session.NumbersCount)

	time.Sleep(interval)
	// Get scheduled list

	listScheduled, _ := client.GetScheduledList(nil)

	assert.NotEmpty(t, listScheduled.Page)
	assert.NotEmpty(t, listScheduled.Limit)
	assert.NotEmpty(t, listScheduled.PageCount)

	time.Sleep(interval)
	// Delete scheduled message

	err = client.DeleteScheduled(scheduledNew.ID)

	assert.Nil(t, err)
}
