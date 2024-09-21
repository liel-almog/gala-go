package services

import (
	"context"
	"sync"

	"github.com/liel-almog/gala-go/backend/models"
	"github.com/liel-almog/gala-go/backend/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EventService interface {
	GetAll(ctx context.Context) ([]models.Event, error)
	GetById(ctx context.Context, id bson.ObjectID) (*models.Event, error)
	Create(ctx context.Context, event *models.Event) error
}

type eventServiceImpl struct {
	eventRepository repositories.EventRepository
}

var (
	initEventService sync.Once
	eventService     *eventServiceImpl
)

func newEventService() *eventServiceImpl {
	return &eventServiceImpl{
		eventRepository: repositories.GetEventController(),
	}
}

func GetEventService() EventService {
	initEventService.Do(func() {
		eventService = newEventService()
	})

	return eventService
}

func (s *eventServiceImpl) GetAll(ctx context.Context) ([]models.Event, error) {
	return s.eventRepository.FindAll(ctx)
}

func (s *eventServiceImpl) GetById(ctx context.Context, id bson.ObjectID) (*models.Event, error) {
	return s.eventRepository.FindById(ctx, id)
}

func (s *eventServiceImpl) Create(ctx context.Context, event *models.Event) error {
	return s.eventRepository.Create(ctx, event)
}
