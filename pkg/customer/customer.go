package customer

// Customer is the customer struct
type Customer struct {
	name    string
	address string
	phone   string
}

// New return new Customer
func New(name, address, phone string) Customer {
	return Customer{
		name:    name,
		address: address,
		phone:   phone,
	}
}
