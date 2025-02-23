package main

import "fmt"

type ProductComponent interface {
	ShowDetails(indent string)
	GetPrice() float64
}

type Product struct {
	Name  string
	Price float64
}

func (p *Product) ShowDetails(indent string) {
	fmt.Println(indent, "- Product: ", p.Name, " Price: ($", p.Price, ")")
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

type Bundle struct {
	Name     string
	Products []ProductComponent
}

func (b *Bundle) AddProduct(p ProductComponent) {
	b.Products = append(b.Products, p)
}

func (b *Bundle) ShowDetails(indent string) {
	fmt.Println(indent+"+ Bundle:", b.Name)
	for _, p := range b.Products {
		p.ShowDetails(indent)
	}
}

func (b *Bundle) GetPrice() float64 {
	total := 0.0
	for _, p := range b.Products {
		total += p.GetPrice()
	}
	return total
}

func main() {
	laptop := &Product{Name: "Laptop", Price: 1000.0}
	phone := &Product{Name: "Phone", Price: 1100.0}
	headPhone := &Product{Name: "Headphone", Price: 300.0}

	bundle1 := &Bundle{Name: "Tech Started Pack"}
	bundle2 := &Bundle{Name: "Premium Pack"}

	bundle1.AddProduct(laptop)
	bundle1.AddProduct(headPhone)

	bundle2.AddProduct(bundle1)
	bundle2.AddProduct(phone)

	bundle2.ShowDetails("")
}
