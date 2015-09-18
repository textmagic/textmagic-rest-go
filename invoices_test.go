package textmagic

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInvoices(t *testing.T) {
	time.Sleep(interval)

	invoices, err := client.GetInvoiceList(nil)

	if err != nil {
		t.Fatal(err)
	} else if len(invoices.Resources) == 0 {
		return
	}

	assert.NotEmpty(t, invoices.Page)
	assert.NotEmpty(t, invoices.Limit)
	assert.NotEmpty(t, invoices.PageCount)
	assert.NotEqual(t, 0, len(invoices.Resources))

	inv := invoices.Resources[0]

	assert.NotEmpty(t, inv.ID)
	assert.NotEmpty(t, inv.Bundle)
	assert.NotEmpty(t, inv.Currency)
	assert.NotEmpty(t, inv.PaymentMethod)
}
