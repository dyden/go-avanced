package types

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
			price: 1999.99,
		},
	}
}
