package types

type Computer struct {
	name  string
	stock int
	price float64
}

//NAME
func (c *Computer) GetName() string {
	return c.name
}
func (c *Computer) SetName(name string) {
	c.name = name
}

//PRICE
func (c *Computer) GetPrice() float64 {
	return c.price
}
func (c *Computer) SetPrice(price float64) {
	c.price = price
}

//STOCK
func (c *Computer) GetStock() int {
	return c.stock
}
func (c *Computer) SetStock(stock int) {
	c.stock = stock
}
