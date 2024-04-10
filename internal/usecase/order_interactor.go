package usecase

import "onlineStore/internal/entities"

type OrderInteractor struct {
	OrderRepository OrderRepository
}

func (interactor *OrderInteractor) GetOrder(id int) (entities.Order, error) {
	return interactor.OrderRepository.SelectOrder(id)
}
