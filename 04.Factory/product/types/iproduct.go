package types

type IProduct interface {
	//NAME
	SetName(name string)
	GetName() string
	//PRICE
	SetPrice(price float64)
	GetPrice() float64
	//STOCK
	SetStock(stock int)
	GetStock() int
}
