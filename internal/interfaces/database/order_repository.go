package database

import (
	"onlineStore/internal/entities"
)

type OrderRepository struct {
	SqlHandler
}

type Order struct {
	ProductId int
	Quantity  int
}

type Shelf struct {
	Id   int
	Name string
}

type OptionalShelving struct {
	ProductId int
	ShelfId   int
}

type Product struct {
	Id      int
	Name    string
	ShelfId int
}

type parser struct {
	productsMap map[int]entities.Product
	shelvingMap map[int]entities.Shelf
}

func initParser() parser {
	var p parser
	p.productsMap = make(map[int]entities.Product)
	p.shelvingMap = make(map[int]entities.Shelf)
	return p
}

func (or *OrderRepository) SelectOrder(number int) (entities.Order, error) {
	advancedOrder := entities.Order{}
	advancedOrder.Number = number

	p := initParser()

	orders, err := or.GetOrders(number)
	if err != nil {
		return advancedOrder, err
	}
	p.saveOrders(orders)

	products, err := or.GetProducts(keys(p.productsMap))
	if err != nil {
		return advancedOrder, err
	}
	p.saveProducts(products)

	optionalShelving, err := or.GetOptionalShelving(keys(p.productsMap))
	if err != nil {
		return advancedOrder, err
	}
	p.saveOptShelving(optionalShelving)

	shelving, err := or.GetShelving(keys(p.shelvingMap))
	if err != nil {
		return advancedOrder, err
	}
	p.saveShelving(shelving)

	advancedOrder.Products = p.parseData()
	return advancedOrder, nil
}

func (s *parser) saveOrders(orders []Order) {
	for _, order := range orders {
		product := entities.Product{Id: order.ProductId, Quantity: order.Quantity}
		s.productsMap[order.ProductId] = product
	}
}

func (s *parser) saveProducts(products []Product) {
	for _, product := range products {
		p := s.productsMap[product.Id]
		p.Name = product.Name
		p.Shelf.Id = product.ShelfId
		s.shelvingMap[product.ShelfId] = entities.Shelf{Id: product.ShelfId}
	}
}

func (s *parser) saveOptShelving(optionalShelving []OptionalShelving) {
	for _, optionalShelf := range optionalShelving {
		p := s.productsMap[optionalShelf.ProductId]
		p.OptionalShelving = append(p.OptionalShelving, entities.Shelf{Id: optionalShelf.ShelfId})
		s.shelvingMap[optionalShelf.ShelfId] = entities.Shelf{Id: optionalShelf.ShelfId}
	}
}

func (s *parser) saveShelving(shelving []Shelf) {
	for _, shelf := range shelving {
		s.shelvingMap[shelf.Id] = entities.Shelf{Id: shelf.Id, Name: shelf.Name}
	}
}

func (s *parser) parseData() (products []entities.Product) {
	for _, product := range s.productsMap {
		for i, shelf := range product.OptionalShelving {
			product.OptionalShelving[i].Name = s.shelvingMap[shelf.Id].Name
		}
		product.Shelf.Name = s.shelvingMap[product.Shelf.Id].Name
		products = append(products, product)
	}
	return
}

func keys[T any](m map[int]T) (keys []int) {
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
