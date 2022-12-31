package controller

import (
	"context"
	"encoding/json"
	"logger/internal/domain"
	"logger/internal/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type errorresponse struct {
	Error   bool `json:"error"`
	Message any  `json:"message"`
}

// interface
type LogController interface {
	Insert(ec echo.Context) error
	InsertWorker(ctx context.Context, data string) error
}

// implement interface
type LogControllerImpl struct {
	LogUsecase usecase.LogUseCase
}

type successresponse struct {
	Error   bool `json:"error"`
	Message any  `json:"message"`
}

func NewLogController(logUsecase usecase.LogUseCase) LogController {
	return &LogControllerImpl{
		LogUsecase: logUsecase,
	}
}

func (uc *LogControllerImpl) Insert(c echo.Context) error {
	// Convert Echo Context
	con := c.Request().Context()
	ctx, cancel := context.WithTimeout(con, 10000*time.Second)
	defer cancel()

	// Validation
	u := new(domain.Log)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Registering User
	err := uc.LogUsecase.InsertLog(ctx, u)

	if err != nil {
		response := errorresponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := successresponse{
		Error:   false,
		Message: "Berhasil mendaftar",
	}

	return c.JSON(http.StatusOK, response)
}

func (uc *LogControllerImpl) InsertWorker(c context.Context, data string) error {
	logDomain := &domain.Log{}
	var jsonData = []byte(data)

	var _ = json.Unmarshal(jsonData, &logDomain)

	// Registering User
	err := uc.LogUsecase.InsertLog(c, logDomain)

	return err
}
