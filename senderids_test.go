package textmagic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSenderids(t *testing.T) {
	// Requires support approval
	t.SkipNow()
	time.Sleep(interval)

	// Create sender ID

	senderIDData := Params{
		"senderId":    "testapi",
		"explanation": "need for testing",
	}
	sidNew, err := client.CreateSenderID(senderIDData)

	assert.NotEmpty(t, sidNew.ID)
	assert.NotEmpty(t, sidNew.Href)

	time.Sleep(interval)
	// Get Sender ID by id

	sid, _ := client.GetSenderID(sidNew.ID)

	assert.NotEmpty(t, sid.ID)
	assert.NotEmpty(t, sid.SenderID)
	assert.NotEmpty(t, sid.Status)

	user := sid.User

	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Username)
	assert.NotEmpty(t, user.FirstName)
	assert.NotEmpty(t, user.LastName)
	assert.NotEmpty(t, user.Status)
	assert.NotEmpty(t, user.Balance)
	assert.NotEmpty(t, user.SubaccountType)

	currency := user.Currency

	assert.NotEmpty(t, currency.ID)
	assert.NotEmpty(t, currency.HTMLSymbol)

	tz := user.Timezone

	assert.NotEmpty(t, tz.ID)
	assert.NotEmpty(t, tz.Area)
	assert.NotEmpty(t, tz.Timezone)

	time.Sleep(interval)
	// Get sources

	sources, _ := client.GetSources(nil)

	assert.NotEmpty(t, sources.Dedicated)
	assert.NotEmpty(t, sources.Shared)
	assert.NotEmpty(t, sources.SenderIDs)

	time.Sleep(interval)
	// Get Sender IDs List

	senderIDs, err := client.GetSenderIDList(nil)

	assert.NotEmpty(t, senderIDs.Page)
	assert.NotEmpty(t, senderIDs.Limit)
	assert.NotEmpty(t, senderIDs.PageCount)
	assert.NotEqual(t, 0, len(senderIDs.Resources))

	time.Sleep(interval)
	// Delete sender ID

	err = client.DeleteSenderID(sidNew.ID)
	assert.Nil(t, err)
}
