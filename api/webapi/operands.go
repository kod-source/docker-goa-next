package webapi

import (
	"strconv"

	"github.com/kod-source/docker-goa-next/webapi/app"
	goa "github.com/shogo82148/goa-v1"
)

// OperandsController implements the operands resource.
type OperandsController struct {
	*goa.Controller
}

// NewOperandsController creates a operands controller.
func NewOperandsController(service *goa.Service) *OperandsController {
	return &OperandsController{Controller: service.NewController("OperandsController")}
}

// Add runs the add action.
func (c *OperandsController) Add(ctx *app.AddOperandsContext) error {
	sum := ctx.Left + ctx.Right
	return ctx.OK([]byte(strconv.Itoa(sum)))
}
