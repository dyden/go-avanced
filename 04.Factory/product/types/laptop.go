package types

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop",
			stock: 10,
			price: 999.99,
		},
	}
}
