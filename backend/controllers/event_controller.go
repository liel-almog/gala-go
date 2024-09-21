package controllers

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/liel-almog/gala-go/backend/configs"
	"github.com/liel-almog/gala-go/backend/models"
	"github.com/liel-almog/gala-go/backend/services"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EventController interface {
	GetAllEvents(c echo.Context) error
	GetEventById(c echo.Context) error
	CreateEvent(c echo.Context) error
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

func (c *eventControllerImpl) GetEventById(ctx echo.Context) error {
	id, err := bson.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "Invalid ID",
		})
	}

	event, err := c.eventService.GetById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"event": event,
	})
}

func (c *eventControllerImpl) CreateEvent(ctx echo.Context) error {
	event := models.NewEvent()
	if err := ctx.Bind(&event); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "Invalid request body",
		})
	}

	if err := configs.GetValidator().Struct(event); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	err := c.eventService.Create(ctx.Request().Context(), event)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(201, echo.Map{})
}
