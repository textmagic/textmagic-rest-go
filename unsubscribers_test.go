package textmagic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnsubscribers(t *testing.T) {
	phone := "9993243232"

	time.Sleep(interval)
	// Unsubscribe phone number

	uNew, err := client.UnsubscribePhone(phone)

	assert.Nil(t, err)
	assert.NotEmpty(t, uNew.ID)
	assert.NotEmpty(t, uNew.Href)

	time.Sleep(interval)
	// Get unsubscriber

	u, err := client.GetUnsubscriber(uNew.ID)

	assert.Nil(t, err)
	assert.NotEmpty(t, u.ID)
	assert.Equal(t, phone, u.Phone)
	assert.NotEmpty(t, u.UnsubscribeTime)
	assert.NotEmpty(t, u.FirstName)
	assert.NotEmpty(t, u.LastName)

	time.Sleep(interval)
	// Get unsubscriber list

	list, err := client.GetUnsubscriberList(nil)

	assert.Nil(t, err)
	assert.NotEmpty(t, list.Page)
	assert.NotEmpty(t, list.Limit)
	assert.NotEqual(t, len(list.Resources), 0)
}
