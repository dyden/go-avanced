package main

import "fmt"

func main() {
	nvidiaItem := NewItem("RTX 3090")
	firstObserver := &EmailClient{
		id: "12345-ID",
	}
	secondObserver := &EmailClient{
		id: "67890-ID",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.UpdateAvailable()
}

/***************************************************************
*                          INTERFACES                          *
****************************************************************/

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> Not available
// Item -> Notify -> Items available

/***************************************************************
*                          STRUCTS                              *
****************************************************************/
type Item struct {
	observers []Observer
	name      string
	available bool
}
type EmailClient struct {
	id string
}

/***************************************************************
*                          FUNCTIONS                           *
****************************************************************/

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}
func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s\n", value, eC.id)
}

func (eC *EmailClient) getId() string {
	return eC.id
}
