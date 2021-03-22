package factory

import (
	"fmt"

	. "github.com/dyden/go-avanced/04.Factory/product/types"
)

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return NewLaptop(), nil
	}

	if computerType == "desktop" {
		return NewDesktop(), nil
	}

	return nil, fmt.Errorf("invalid computer type")
}
func PrintNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d and price %f\n", p.GetName(), p.GetStock(), p.GetPrice())
}
