package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

func New(id uint, product string, value float64) Item {
	return Item{
		id:      id,
		product: product,
		value:   value,
	}
}

func (i Item) ID() uint {
	return i.id
}

func (i Item) Product() string {
	return i.product
}

func (i Item) Value() float64 {
	return i.value
}
