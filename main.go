package main

import (
	"fmt"

	"github.com/jorgemarquez2222/composition/pkg/invoice"
)

func main() {
	i := invoice.New(
		"Mexico",
		"Mexico City",
		customer.New("Jorge", "Mexico City", "5555555555"),
		[]*invoice.Item{
			invoiceitem.New(1, "Laptop", 1000),
			invoiceitem.New(2, "Mouse", 10),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v\n", i)
}
