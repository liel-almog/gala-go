package services

import (
	"context"
	"sync"

	"github.com/liel-almog/gala-go/backend/models"
	"github.com/liel-almog/gala-go/backend/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GuestService interface {
	GetAll(ctx context.Context) ([]models.Guest, error)
	GetById(ctx context.Context, id bson.ObjectID) (*models.Guest, error)
	Create(ctx context.Context, guest *models.Guest) error
}

type guestServiceImpl struct {
	guestRepository repositories.GuestRepository
}

var (
	initGuestService sync.Once
	guestService     *guestServiceImpl
)

func newGuestService() *guestServiceImpl {
	return &guestServiceImpl{
		guestRepository: repositories.GetGuestRepository(),
	}
}

func GetGuestService() GuestService {
	initGuestService.Do(func() {
		guestService = newGuestService()
	})

	return guestService
}

func (s *guestServiceImpl) GetAll(ctx context.Context) ([]models.Guest, error) {
	return s.guestRepository.FindAll(ctx)
}

func (s *guestServiceImpl) GetById(ctx context.Context, id bson.ObjectID) (*models.Guest, error) {
	return s.guestRepository.FindById(ctx, id)
}

func (s *guestServiceImpl) Create(ctx context.Context, guest *models.Guest) error {
	return s.guestRepository.Create(ctx, guest)
}
