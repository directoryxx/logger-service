package api

import (
	"logger/internal/controller"

	"github.com/labstack/echo/v4"
)

func Routes(
	router *echo.Echo,
	LogController controller.LogController,
) {

	router.POST("/loginsert", LogController.Insert)
}
