package invoice

type Invoice struct {
	country string
	city    string
	total   string
	client  *Customer.Customer
	items   []*Item.Item
}

// New return new Invoice
func New(country, city, client *Customer.Customer, items []*Item.Item) *Invoice {
	return &Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of invoice.go
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}
