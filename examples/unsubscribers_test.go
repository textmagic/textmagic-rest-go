package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUnsubscribers(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := 500 * time.Millisecond
	client := textmagic.NewClient(username, token)

	phone := "9993243232"

	time.Sleep(interval)
	// Unsubscribe phone number

	uNew, err := client.UnsubscribePhone(phone)

	assert.Nil(t, err)
	assert.NotEmpty(t, uNew.Id)
	assert.NotEmpty(t, uNew.Href)

	time.Sleep(interval)
	// Get unsubscriber

	u, err := client.GetUnsubscriber(uNew.Id)

	assert.Nil(t, err)
	assert.NotEmpty(t, u.Id)
	assert.Equal(t, phone, u.Phone)
	assert.NotEmpty(t, u.UnsubscribeTime)
	assert.NotEmpty(t, u.FirstName)
	assert.NotEmpty(t, u.LastName)

	time.Sleep(interval)
	// Get unsubscriber list

	list, err := client.GetUnsubscriberList(map[string]string{})

	assert.Nil(t, err)
	assert.NotEmpty(t, list.Page)
	assert.NotEmpty(t, list.Limit)
	assert.NotEqual(t, len(list.Unsubscribers), 0)
}
