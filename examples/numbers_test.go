package textmagic

import (
	".."
	"github.com/stretchr/testify/assert"
	// "strconv"
	"testing"
	"time"
)

func TestNumbers(t *testing.T) {
	username := "xxx"
	token := "xxx"

	interval := 500 * time.Millisecond
	client := textmagic.NewClient(username, token)

	time.Sleep(interval)
	// Get available numbers

	available, _ := client.GetAvailableNumbers(
		map[string]string{
			"country": "US",
		},
	)

	assert.NotEmpty(t, len(available.Numbers))
	assert.NotEmpty(t, available.Price)

	time.Sleep(interval)
	// Get User

	currentUser, _ := client.GetUser()

	/* !! Please be careful if you want to uncomment this because it will spend the money from your account !! */

	/*
		time.Sleep(interval)
		// Buy number

		buyNumberData := map[string]string{
			"phone":   available.Numbers[0],
			"country": "US",
			"userId":  strconv.Itoa(int(currentUser.Id)),
		}
		numberNew, _ := client.BuyNumber(buyNumberData)

		assert.NotEmpty(t, numberNew.Id)
		assert.NotEmpty(t, numberNew.Href)
	*/

	time.Sleep(interval)
	// Get Numbers List

	numbers, _ := client.GetNumberList(map[string]string{})

	assert.NotEmpty(t, numbers.Page)
	assert.NotEmpty(t, numbers.Limit)
	assert.NotEmpty(t, numbers.PageCount)
	assert.NotEqual(t, 0, len(numbers.Numbers))

	time.Sleep(interval)
	// Get Number by id

	number, _ := client.GetNumber(numbers.Numbers[0].Id)

	assert.NotEmpty(t, number.Id)
	assert.NotEmpty(t, number.PurchasedAt)
	assert.NotEmpty(t, number.ExpireAt)
	assert.NotEmpty(t, number.Phone)
	assert.NotEmpty(t, number.Status)

	assert.Equal(t, numbers.Numbers[0].Id, number.Id)
	assert.Equal(t, numbers.Numbers[0].Phone, number.Phone)

	country := number.Country

	assert.Equal(t, country.Id, "US")
	assert.Equal(t, country.Name, "United States")

	user := number.User

	assert.Equal(t, currentUser.Id, user.Id)
	assert.Equal(t, currentUser.Username, user.Username)
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

	/* !! Please be careful if you want to uncomment this because it will close one of your subaccount !! */

	/*
		time.Sleep(interval)
		// Cancel dedicated number

		r, _ := client.CancelNumber(number.Id)

		assert.True(t, r)
	*/
}
