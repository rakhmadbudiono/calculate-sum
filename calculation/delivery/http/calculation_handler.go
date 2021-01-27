package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"

	"github.com/rakhmadbudiono/calculate-sum/domain"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// CalculationHandler represent the httphandler for sum
type CalculationHandler struct {
	CUsecase domain.CalculationUsecase
}

// NewCalculationHandler will initialize the calculation/ resources endpoint
func NewCalculationHandler(e *echo.Echo, us domain.CalculationUsecase) {
	handler := &CalculationHandler{
		CUsecase: us,
	}
	e.GET("/sum", handler.Sum)
}

// Sum will give the result based on given params
func (cal *CalculationHandler) Sum(c echo.Context) error {
	aS := c.QueryParam("a")
	a, _ := strconv.Atoi(aS)
	bS := c.QueryParam("b")
	b, _ := strconv.Atoi(bS)
	ctx := c.Request().Context()

	result, err := cal.CUsecase.Add(ctx, int64(a), int64(b))
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
