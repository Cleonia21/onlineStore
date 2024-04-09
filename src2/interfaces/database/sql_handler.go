package database

type SqlHandler interface {
	GetOrders(id int) []Order
	GetShelf(id int) Shelf
	GetProducts(id []int) []Product
	GetOptionalShelving(ProductId int) []Shelf
}
