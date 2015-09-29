package textmagic

const invoiceURI = "invoices"

// Invoice represents an invoice.
type Invoice struct {
	ID            int    `json:"id"`
	Bundle        int    `json:"bundle"`
	Currency      string `json:"currency"`
	Vat           int    `json:"vat"`
	PaymentMethod string `json:"paymentMethod"`
}

// InvoiceList represents a list of invoices
// and pagination information.
type InvoiceList struct {
	Page      int        `json:"page"`
	Limit     int        `json:"limit"`
	PageCount int        `json:"pageCount"`
	Resources []*Invoice `json:"resources"`
}

// GetInvoiceList returns all user invoices.
//
// The parameter payload includes:
// - page:	Fetch specified results page.
// - limit:	How many results on page.
func (c *Client) GetInvoiceList(p Params) (*InvoiceList, error) {
	var l *InvoiceList

	return l, c.get(invoiceURI, p, nil, &l)
}
