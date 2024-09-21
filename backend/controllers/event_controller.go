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
	events, err := c.eventService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"events": events,
	})
}
