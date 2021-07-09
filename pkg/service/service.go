package service

import (
	"context"
	"gitlab.com/zharzhanov/region/models"
	"gitlab.com/zharzhanov/region/pkg/repository"
)

type Authentication interface {
	SignUp(ctx context.Context, user models.User) (string, error)
	SignIn(ctx context.Context, user models.User) (string, error)
}

type Adverts interface {
	CreateAdvert(ctx context.Context, advert models.Advert) (string, error)
	GetAllAdverts(ctx context.Context) ([]models.Advert, error)
	GetAdvertById(ctx context.Context, id string) (*models.Advert, error)
	UpdateAdvert(ctx context.Context, id string , advert models.UpdateAdvertInput) error
	DeleteAdvert(ctx context.Context, id string) error
}

type Users interface {

}

type Service struct {
	Authentication
	Adverts
	Users
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repos),
		Adverts: NewAdvertService(repos),
	}
}