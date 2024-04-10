package controller

import (
	"onlineStore/internal/entities"
	"onlineStore/internal/interfaces/database"
	"onlineStore/internal/usecase"
)

type OrderController struct {
	Interactor usecase.OrderInteractor
}

func NewOrderController(sqlHandler database.SqlHandler) *OrderController {
	return &OrderController{
		Interactor: usecase.OrderInteractor{
			OrderRepository: &database.OrderRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

type Order struct {
	Number        int
	ProductName   string
	ProductId     int
	Quantity      int
	MainShelf     string
	OptionalShelf []string
}

func (controller *OrderController) GetOrders(ids []int) (map[string][]Order, error) {
	entOrders, err := controller.collectOrders(ids)
	if err != nil {
		return nil, err
	}
	orders := convertOrders(entOrders)
	orderSortMap := sortOrdersByShelving(orders)
	return orderSortMap, nil
}

func (controller *OrderController) collectOrders(ids []int) (orders []entities.Order, err error) {
	for _, id := range ids {
		order, err := controller.Interactor.GetOrder(id)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return
}

func convertOrders(entOrders []entities.Order) (orders []Order) {
	for _, entOrder := range entOrders {
		for _, entProduct := range entOrder.Products {
			order := Order{
				Number:        entOrder.Number,
				ProductName:   entProduct.Name,
				ProductId:     entProduct.Id,
				Quantity:      entProduct.Quantity,
				MainShelf:     entProduct.Shelf.Name,
				OptionalShelf: nil,
			}
			for _, entShelf := range entProduct.OptionalShelving {
				order.OptionalShelf = append(order.OptionalShelf, entShelf.Name)
			}
		}
	}
	return orders
}

func sortOrdersByShelving(orders []Order) map[string][]Order {
	sortMap := make(map[string][]Order)
	for _, o := range orders {
		sortMap[o.MainShelf] = append(sortMap[o.MainShelf], o)
	}
	return sortMap
}
