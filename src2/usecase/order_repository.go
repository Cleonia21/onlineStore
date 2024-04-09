package usecase

import "onlineStore/src2/entities"

type OrderRepository interface {
	Select(id int) entities.Order
}
