package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/liel-almog/gala-go/backend/controllers"
)

func NewEventRoute(router *echo.Group) {
	group := router.Group("/events")

	controller := controllers.GetEventController()
	group.GET("", controller.GetAllEvents)
	group.GET("/:id", controller.GetEventById)
	group.POST("", controller.CreateEvent)
}
