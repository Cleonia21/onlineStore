package database

type SqlHandler interface {
	GetOrders(id int) ([]Order, error)
	GetProducts(id []int) ([]Product, error)
	GetShelving(id []int) ([]Shelf, error)
	GetOptionalShelving(ProductId []int) ([]OptionalShelving, error)
}
