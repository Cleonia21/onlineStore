package entities

type Shelf struct {
	Id   int
	Name string
}

type Product struct {
	Id               int
	Name             string
	Quantity         int
	Shelf            Shelf
	OptionalShelving []Shelf
}

type Order struct {
	Number   int
	Products []Product
}
