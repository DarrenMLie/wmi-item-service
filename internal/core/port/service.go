package port

import "wmi-item-service/internal/core/domain"

type AuthService interface {
	SignUp(domain.SignUpRequest) error
	SignIn(domain.SignInRequest) (*domain.User, error)
}

type ResidenceService interface {
	CreateResidence(domain.CreateResidenceRequest) (*domain.Residence, error)
	UpdateResidence(domain.UpdateResidenceRequest) (*domain.Residence, error)
	GetResidence(domain.GetResidenceRequest) (*domain.Residence, error)
	GetResidenceList(domain.GetResidenceListRequest) (*domain.MetaResidences, error)
	DeleteResidence(domain.DeleteResidenceRequest) (error)
}

type ItemService interface {
	CreateItem(domain.CreateItemRequest) (*domain.Item, error)
	UpdateItem(domain.UpdateItemRequest) (*domain.Item, error)
	GetItem(domain.GetItemRequest) (*domain.Item, error)
	GetItemList(domain.GetItemListRequest) (*domain.MetaItems, error)
	DeleteItem(domain.DeleteItemRequest) (error)
}