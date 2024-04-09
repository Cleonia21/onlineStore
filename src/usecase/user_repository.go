package usecase

import "onlineStore/src/domain"

type UserRepository interface {
	Store(domain.User)
	Select() []domain.User
	Delete(id string)
}
