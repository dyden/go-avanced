package main

import (
	. "github.com/dyden/go-avanced/04.Factory/product"
)

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
