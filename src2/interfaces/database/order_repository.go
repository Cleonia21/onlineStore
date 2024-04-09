package database

import "onlineStore/src2/entities"

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

type Product struct {
	Id      int
	Name    string
	ShelfId int
}

func (or *OrderRepository) Select(number int) entities.Order {
	advancedOrder := entities.Order{}
	advancedOrder.Number = number

	orders := or.GetOrders(number)
	var productsId []int
	for _, order := range orders {
		productsId = append(productsId, order.ProductId)

		var advancedProduct entities.Product
		advancedProduct.Quantity = order.Quantity
		advancedOrder.Products = append(advancedOrder.Products, advancedProduct)
	}

	products := or.GetProducts(productsId)
	shelvingDictionary := make(map[int]string)
	for i, product := range products {
		shelvingId = append(shelvingId, product.ShelfId)

		advancedOrder.Products[i].Name = product.Name
		advancedOrder.Products[i].Id = product.Id
	}

}
