package textmagic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNumbers(t *testing.T) {
	// Skip integration test if account has no funds
	t.SkipNow()

	time.Sleep(interval)
	// Get available numbers

	available, err := client.GetAvailableNumbers(Params{"country": "US"})

	assert.NotEmpty(t, len(available.Numbers))
	assert.NotEmpty(t, available.Price)

	time.Sleep(interval)
	// Get User

	currentUser, err := client.GetUser()

	assert.Nil(t, err)

	time.Sleep(interval)

	numbers, err := client.GetNumberList(nil)

	if err != nil {
		t.Fatal(err)
	}

	time.Sleep(interval)

	var numID int

	if len(numbers.Resources) == 0 {
		data := NewParams("userId", currentUser.ID)
		data.Set("country", "US")
		data.Set("phone", available.Numbers[0])

		newNumber, err := client.BuyNumber(data)

		if err != nil {
			t.Fatalf("%#v", err)
		}

		if newNumber == nil || newNumber.ID == 0 {
			t.Fatal("unable to buy number")
		}

		numID = newNumber.ID
	} else {
		numID = numbers.Resources[0].ID
	}

	number, err := client.GetNumber(numID)

	assert.Nil(t, err)
	assert.NotNil(t, numbers)

	time.Sleep(interval)

	assert.NotEmpty(t, number.ID)
	assert.NotEmpty(t, number.PurchasedAt)
	assert.NotEmpty(t, number.ExpireAt)
	assert.NotEmpty(t, number.Phone)
	assert.NotEmpty(t, number.Status)

	assert.Equal(t, numbers.Resources[0].ID, number.ID)
	assert.Equal(t, numbers.Resources[0].Phone, number.Phone)

	country := number.Country

	assert.Equal(t, country.ID, "US")
	assert.Equal(t, country.Name, "United States")

	user := number.User

	assert.Equal(t, currentUser.ID, user.ID)
	assert.Equal(t, currentUser.Username, user.Username)
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
}
