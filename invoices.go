package textmagic

import (
	"encoding/json"
	"fmt"
)

const (
	INVOICE_RES = "invoices"
)

type Invoice struct {
	Id            uint32 `json:"id"`
	Bundle        int    `json:"bundle"`
	Currency      string `json:"currency"`
	Vat           int    `json:"vat"`
	PaymentMethod string `json:"paymentMethod"`
}

type InvoiceList struct {
	Page      uint32    `json:"page"`
	Limit     uint8     `json:"limit"`
	PageCount uint32    `json:"pageCount"`
	Invoices  []Invoice `json:"resources"`
}

/*
Get all user invoices.

	Parameters:

		Var `data` may contains next keys:

			page:  Fetch specified results page.
			limit: How many results on page.
*/
func (client *TextmagicRestClient) GetInvoiceList(data map[string]string) (*InvoiceList, error) {
	var invoiceList = new(InvoiceList)

	method := "GET"
	params := transformGetParams(data)

	uri := fmt.Sprintf("%s/%s", client.BaseUrl(), INVOICE_RES)

	response, err := client.Request(method, uri, params, nil)
	if err != nil {
		return invoiceList, err
	}

	err = json.Unmarshal(response, invoiceList)

	return invoiceList, err
}
