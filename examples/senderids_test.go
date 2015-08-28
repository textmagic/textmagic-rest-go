package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSenderids(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := time.Second
	client := textmagic.NewClient(username, token)

	time.Sleep(interval)
	// Create sender ID

	senderIdData := map[string]string{
		"senderId":    "GOLANGTEST",
		"explanation": "need for testing",
	}
	sidNew, _ := client.CreateSenderId(senderIdData)

	assert.NotEmpty(t, sidNew.Id)
	assert.NotEmpty(t, sidNew.Href)

	time.Sleep(interval)
	// Get Sender ID by id

	sid, _ := client.GetSenderId(sidNew.Id)

	assert.NotEmpty(t, sid.Id)
	assert.NotEmpty(t, sid.SenderId)
	assert.NotEmpty(t, sid.Status)

	user := sid.User

	assert.NotEmpty(t, user.Id)
	assert.NotEmpty(t, user.Username)
	assert.NotEmpty(t, user.FirstName)
	assert.NotEmpty(t, user.LastName)
	assert.NotEmpty(t, user.Status)
	assert.NotEmpty(t, user.Balance)
	assert.NotEmpty(t, user.SubaccountType)

	currency := user.Currency

	assert.NotEmpty(t, currency.Id)
	assert.NotEmpty(t, currency.HtmlSymbol)

	tz := user.Timezone

	assert.NotEmpty(t, tz.Id)
	assert.NotEmpty(t, tz.Area)
	assert.NotEmpty(t, tz.Timezone)

	time.Sleep(interval)
	// Get sources

	sources, _ := client.GetSources(map[string]string{})

	assert.NotEmpty(t, sources.Dedicated)
	assert.NotEmpty(t, sources.Shared)
	assert.NotEmpty(t, sources.SenderIds)

	time.Sleep(interval)
	// Get Sender IDs List

	senderIds, _ := client.GetSenderIdList(map[string]string{})

	assert.NotEmpty(t, senderIds.Page)
	assert.NotEmpty(t, senderIds.Limit)
	assert.NotEmpty(t, senderIds.PageCount)
	assert.NotEqual(t, 0, len(senderIds.Resources))

	time.Sleep(interval)
	// Delete sender ID

	r, _ := client.DeleteSenderId(sidNew.Id)
	assert.True(t, r)
}
