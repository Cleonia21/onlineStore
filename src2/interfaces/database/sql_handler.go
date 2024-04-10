package database

type SqlHandler interface {
	GetOrders(id int) []Order
	GetProducts(id []int) []Product
	GetShelving(id []int) []Shelf
	GetOptionalShelving(ProductId []int) []OptionalShelving
}
