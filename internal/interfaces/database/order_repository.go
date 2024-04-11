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
	p.saveOptionalShelving(optionalShelving)

	shelving, err := or.GetShelving(keys(p.shelvingMap))
	if err != nil {
		return advancedOrder, err
	}
	p.saveShelving(shelving)

	advancedOrder.Products = p.parseData()
	return advancedOrder, nil
}
