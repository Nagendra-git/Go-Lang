package main

import "fmt"

func main() {

	//type inference
	var product = struct {
		id   int
		name string
		cost float32
	}{
		id:   100,
		name: "pen",
		cost: 10,
	}

	// fmt.Println(product)
	// fmt.Printf("%#v\n", product)
	// fmt.Printf("%+v\n", product)
	fmt.Println(Format(product)) //=> should print "id = 100, name = "Pen", cost = 10"
	ApplyDiscount(&product, 10)
	fmt.Println(Format(product)) //=> should print "id = 100, name = "Pen", cost = 9"
}

func Format(p struct {
	id   int
	name string
	cost float32
}) string {
	return fmt.Sprintf("id = %d, name = %q, cost = %0.2f", p.id, p.name, p.cost)
}

func ApplyDiscount(p *struct {
	id   int
	name string
	cost float32
}, discountPercentage float32) {
	p.cost = p.cost * ((100 - discountPercentage) / 100)
}
