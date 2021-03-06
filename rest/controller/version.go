package controller

import (
	"github.com/goadesign/goa"
	"github.com/yogihardi/guestbook/rest/app"
	"github.com/yogihardi/guestbook/version"
)

// VersionController implements the version resource.
type VersionController struct {
	*goa.Controller
}

// NewVersionController creates a version controller.
func NewVersionController(service *goa.Service) *VersionController {
	return &VersionController{Controller: service.NewController("VersionController")}
}

// Version runs the version action.
func (c *VersionController) Version(ctx *app.VersionVersionContext) error {
	// VersionController_Version: start_implement
	res := &app.GuestbookVersion{
		Git:     &version.GitCommit,
		Version: version.Version,
	}
	return ctx.OK(res)
	// VersionController_Version: end_implement
}
