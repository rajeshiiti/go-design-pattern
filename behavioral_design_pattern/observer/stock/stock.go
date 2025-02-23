package main

import "fmt"

// observer
// invester

// subject
// stock

type Observer interface {
	Update(amount float64)
}

type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	Notify()
}

type Stock struct {
	observers []Observer
	price     float64
}

func (s *Stock) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Stock) Deregister(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			fmt.Println("Deregister the observer")
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *Stock) Notify() {
	for _, o := range s.observers {
		o.Update(s.price)
	}
}

func (s *Stock) SetPrice(price float64) {
	fmt.Println("Updated share price to : ", price)
	s.price = price
	s.Notify()
}

type Invester struct {
	Name string
}

func (i *Invester) Update(amount float64) {
	fmt.Println("Name: ", i.Name, " updated price of stock: ", amount)
}

func main() {

	invester1 := &Invester{Name: "Ram"}
	invester2 := &Invester{Name: "Sita"}
	appleStock := Stock{}

	appleStock.Register(invester1)
	appleStock.Register(invester2)

	appleStock.SetPrice(110.0)
	appleStock.SetPrice(115.0)

	appleStock.Deregister(invester1)

	appleStock.SetPrice(114.0)

}
