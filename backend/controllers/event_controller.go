package controllers

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/liel-almog/gala-go/backend/services"
)

type EventController interface {
	GetAllEvents(c echo.Context) error
}

type eventControllerImpl struct {
	eventService services.EventService
}

var (
	initEventController sync.Once
	eventController     *eventControllerImpl
)

func newEventController() *eventControllerImpl {
	return &eventControllerImpl{
		eventService: services.GetEventService(),
	}
}

func GetEventController() EventController {
	initEventController.Do(func() {
		eventController = newEventController()
	})

	return eventController
}

func (c *eventControllerImpl) GetAllEvents(ctx echo.Context) error {
	return ctx.JSON(200, echo.Map{
		"message": "Hello from event controller",
	})
}
