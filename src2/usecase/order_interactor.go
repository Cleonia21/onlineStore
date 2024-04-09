package usecase

import "onlineStore/src2/entities"

type OrderInteractor struct {
	OrderRepository OrderRepository
}

func (interactor *OrderInteractor) GetOrder(id int) entities.Order {
	return interactor.OrderRepository.Select(id)
}
