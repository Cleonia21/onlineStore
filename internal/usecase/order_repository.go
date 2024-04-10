package usecase

import "onlineStore/internal/entities"

type OrderRepository interface {
	SelectOrder(id int) (entities.Order, error)
}
