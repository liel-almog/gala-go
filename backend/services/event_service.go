package services

import (
	"sync"
)

type EventService interface {
}

type eventServiceImpl struct{}

var (
	initEventService sync.Once
	eventService     *eventServiceImpl
)

func newEventService() *eventServiceImpl {
	return &eventServiceImpl{}
}

func GetEventService() EventService {
	initEventService.Do(func() {
		eventService = newEventService()
	})

	return eventService
}
