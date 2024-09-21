package services

import (
	"context"
	"sync"

	"github.com/liel-almog/gala-go/backend/models"
	"github.com/liel-almog/gala-go/backend/repositories"
)

type EventService interface {
	GetAll(ctx context.Context) ([]models.Event, error)
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
