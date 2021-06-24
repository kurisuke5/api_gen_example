// Package health ...
// generated version: devel
package health

import (
	"net/http"

	"github.com/kurisuke5/api_gen_example/backend/domain/werror"
	"github.com/kurisuke5/api_gen_example/backend/interfaces/props"
	"github.com/kurisuke5/api_gen_example/backend/interfaces/wrapper"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

// GetHealthCheckController ...
type GetHealthCheckController struct {
	*props.ControllerProps
}

// NewGetHealthCheckController ...
func NewGetHealthCheckController(cp *props.ControllerProps) *GetHealthCheckController {
	g := &GetHealthCheckController{
		ControllerProps: cp,
	}
	return g
}

var itemCode = "itemCode"


// GetHealthCheck ...
// @Summary WIP
// @Description WIP
// @Accept json
// @Produce json
// @Success 200 {object} GetHealthCheckResponse
// @Failure 400 {object} wrapper.APIError
// @Failure 500 {object} wrapper.APIError
// @Router /api/health/health_check [GET]
func (g *GetHealthCheckController) GetHealthCheck(
	c echo.Context, req *GetHealthCheckRequest,
) (res *GetHealthCheckResponse, err error) {
	var werr *werror.ErrorResponse

	ctx := c.Request().Context()

	if err = g.HealthService.Check(ctx); err != nil {
		if xerrors.As(err, &werr) {
			return nil, wrapper.NewAPIError(werr.Status, werr).SetError(err)
		}
		return nil, wrapper.NewAPIError(http.StatusInternalServerError).SetError(err)
	}

	res = &GetHealthCheckResponse{
		Status: http.StatusOK,
	}

	return res, nil
}
