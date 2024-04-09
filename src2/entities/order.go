package entities

type Shelf struct {
	Id   int
	Name string
	Main bool
}

type Product struct {
	Id       int
	Name     string
	Quantity int
	Shelving []Shelf
}

type Order struct {
	Number   int
	Products []Product
}
