package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := time.Second
	client := textmagic.NewClient(username, token)

	time.Sleep(interval)
	// Get messaging stats

	msgStatData := map[string]string{
		"by":    "month",
		"start": "1 year ago",
		"end":   "now",
	}
	messagingStat, _ := client.GetMessagingStat(msgStatData)

	assert.NotEqual(t, 0, len(messagingStat))
	assert.NotEmpty(t, messagingStat[0].ReplyRate)
	assert.NotEmpty(t, messagingStat[0].Date)
	assert.NotEmpty(t, messagingStat[0].DeliveryRate)
	assert.NotEmpty(t, messagingStat[0].Costs)
	assert.NotEmpty(t, messagingStat[0].MessagesReceived)
	assert.NotEmpty(t, messagingStat[0].MessagesSentDelivered)
	assert.NotEmpty(t, messagingStat[0].MessagesSentAccepted)
	assert.NotEmpty(t, messagingStat[0].MessagesSentBuffered)
	assert.NotEmpty(t, messagingStat[0].MessagesSentFailed)
	assert.NotEmpty(t, messagingStat[0].MessagesSentRejected)
	assert.NotEmpty(t, messagingStat[0].MessagesSentParts)

	time.Sleep(interval)
	// Get spending stats

	spdgStatData := map[string]string{
		"start": "1 year ago",
	}

	spdgStat, _ := client.GetSpendingStat(spdgStatData)

	assert.NotEmpty(t, spdgStat.Page)
	assert.NotEmpty(t, spdgStat.Limit)
	assert.NotEmpty(t, spdgStat.PageCount)
	assert.NotEmpty(t, len(spdgStat.Resources))

	sstat := spdgStat.Resources[0]

	assert.NotEmpty(t, sstat.Id)
	assert.NotEmpty(t, sstat.Date)
	assert.NotEmpty(t, sstat.Balance)
	assert.NotEmpty(t, sstat.Type)
	assert.NotEmpty(t, sstat.Value)

	/*
		time.Sleep(interval)
		// Send invite to subaccount

		inviteData := map[string]string{
			"email": "golangtest15@mailinator.com",
			"role":  "U",
		}
		r, _ := client.SendInvite(inviteData)

		assert.True(t, r)
	*/

	time.Sleep(interval)
	// Get Subaccount List

	subaccounts, _ := client.GetSubaccountList(map[string]string{})

	assert.NotEmpty(t, subaccounts.Page)
	assert.NotEmpty(t, subaccounts.Limit)
	assert.NotEmpty(t, subaccounts.PageCount)
	assert.NotEmpty(t, len(subaccounts.Resources))

	sub1 := subaccounts.Resources[0]

	assert.NotEmpty(t, sub1.Id)
	assert.NotEmpty(t, sub1.Username)
	assert.NotEmpty(t, sub1.FirstName)
	assert.NotEmpty(t, sub1.LastName)
	assert.NotEmpty(t, sub1.Status)
	assert.NotEmpty(t, sub1.Balance)
	assert.NotEmpty(t, sub1.SubaccountType)

	currency := sub1.Currency

	assert.NotEmpty(t, currency.Id)
	assert.NotEmpty(t, currency.HtmlSymbol)

	tz := sub1.Timezone

	assert.NotEmpty(t, tz.Id)
	assert.NotEmpty(t, tz.Area)
	assert.NotEmpty(t, tz.Timezone)

	time.Sleep(interval)
	// Get subaccount by id

	sub2, _ := client.GetSubaccount(sub1.Id)

	assert.NotEmpty(t, sub2.Id)
	assert.NotEmpty(t, sub2.Username)
	assert.NotEmpty(t, sub2.FirstName)
	assert.NotEmpty(t, sub2.LastName)
	assert.NotEmpty(t, sub2.Status)
	assert.NotEmpty(t, sub2.Balance)
	assert.NotEmpty(t, sub2.SubaccountType)

	currency = sub2.Currency

	assert.NotEmpty(t, currency.Id)
	assert.NotEmpty(t, currency.HtmlSymbol)

	tz = sub2.Timezone

	assert.NotEmpty(t, tz.Id)
	assert.NotEmpty(t, tz.Area)
	assert.NotEmpty(t, tz.Timezone)

	/*
		time.Sleep(interval)
		// Delete subaccount by id

		r, _ := client.CloseSubaccount(sub2.Id)

		assert.True(t, r)
	*/

	time.Sleep(interval)
	// Get user info

	user, _ := client.GetUser()

	assert.NotEmpty(t, user.Id)
	assert.NotEmpty(t, user.Username)
	assert.NotEmpty(t, user.Status)
	assert.NotEmpty(t, user.Balance)
	assert.NotEmpty(t, user.SubaccountType)

	currency = user.Currency

	assert.NotEmpty(t, currency.Id)
	assert.NotEmpty(t, currency.HtmlSymbol)

	tz = user.Timezone

	assert.NotEmpty(t, tz.Id)
	assert.NotEmpty(t, tz.Area)
	assert.NotEmpty(t, tz.Timezone)

	time.Sleep(interval)
	// Update user info

	updateUserData := map[string]string{
		"firstName": "GO",
		"lastName":  "Test",
		"company":   "go test",
	}
	res, _ := client.UpdateUser(updateUserData)

	assert.NotEmpty(t, res["href"])

	time.Sleep(interval)

	userUpdated, err := client.GetUser()

	assert.Equal(t, "GO", userUpdated.FirstName)
	assert.Equal(t, "Test", userUpdated.LastName)
	assert.Equal(t, "go test", userUpdated.Company)

	time.Sleep(interval)
	// Update user info

	updateUserData = map[string]string{
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"company":   user.Company,
	}
	res, _ = client.UpdateUser(updateUserData)

	assert.NotEmpty(t, res["href"])

	time.Sleep(interval)

	user1, _ := client.GetUser()

	assert.Equal(t, user.FirstName, user1.FirstName)
	assert.Equal(t, user.LastName, user1.LastName)
	assert.Equal(t, user.Company, user1.Company)

	time.Sleep(interval)
	// Ping

	p, _ := client.Ping()

	assert.Equal(t, "pong", p)
}
