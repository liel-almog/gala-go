package controllers

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/liel-almog/gala-go/backend/models"
	"github.com/liel-almog/gala-go/backend/services"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GuestController interface {
	GetAllGuests(c echo.Context) error
	GetGuestById(c echo.Context) error
	CreateGuest(c echo.Context) error
}

type guestControllerImpl struct {
	guestService services.GuestService
}

var (
	initGuestController sync.Once
	guestController     *guestControllerImpl
)

func newGuestController() *guestControllerImpl {
	return &guestControllerImpl{
		guestService: services.GetGuestService(),
	}
}

func GetGuestController() GuestController {
	initGuestController.Do(func() {
		guestController = newGuestController()
	})

	return guestController
}

func (c *guestControllerImpl) GetAllGuests(ctx echo.Context) error {
	guests, err := c.guestService.GetAll(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"guests": guests,
	})
}

func (c *guestControllerImpl) GetGuestById(ctx echo.Context) error {
	id, err := bson.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "Invalid ID",
		})
	}

	guest, err := c.guestService.GetById(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"guest": guest,
	})
}

func (c *guestControllerImpl) CreateGuest(ctx echo.Context) error {
	var guest models.Guest
	if err := ctx.Bind(&guest); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	if err := c.guestService.Create(ctx.Request().Context(), &guest); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(201, echo.Map{
		"guest": guest,
	})
}
