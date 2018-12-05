package controller

import (
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/uuid"
	"github.com/yogihardi/guestbook/model/servicemodel"
	"github.com/yogihardi/guestbook/rest/app"
	"github.com/yogihardi/guestbook/service"
)

// GuestbookController implements the guestbook resource.
type GuestbookController struct {
	*goa.Controller
	appService service.Service
}

// NewGuestbookController creates a guestbook controller.
func NewGuestbookController(service *goa.Service, appService service.Service) *GuestbookController {
	return &GuestbookController{
		Controller: service.NewController("GuestbookController"),
		appService: appService,
	}
}

// Add runs the add action.
func (c *GuestbookController) Add(ctx *app.AddGuestbookContext) error {
	// GuestbookController_Add: start_implement
	payload := ctx.Payload
	err := c.appService.Add(servicemodel.GuestBook{
		ID:        uuid.NewV4().String(),
		Timestamp: time.Now().Unix(),
		Comment:   payload.Comment,
	})
	if err != nil {
		return ctx.BadRequest(&app.GuestbookError{
			Code: "",
			Msg:  "Can't add an entry: " + err.Error(),
		})
	}

	return ctx.NoContent()
	// GuestbookController_Add: end_implement
}

// Delete runs the delete action.
func (c *GuestbookController) Delete(ctx *app.DeleteGuestbookContext) error {
	// GuestbookController_Delete: start_implement

	err := c.appService.Delete(ctx.ID)
	if err != nil {
		if err.Error() == "Data not found" {
			return ctx.NotFound()
		}
		return ctx.BadRequest(&app.GuestbookError{
			Code: "",
			Msg:  "Can't delete an entry: " + err.Error(),
		})
	}

	return ctx.NoContent()
	// GuestbookController_Delete: end_implement
}

// List runs the list action.
func (c *GuestbookController) List(ctx *app.ListGuestbookContext) error {
	// GuestbookController_List: start_implement

	rows, err := c.appService.List()
	if err != nil {
		return ctx.BadRequest(&app.GuestbookError{
			Code: "",
			Msg:  "Can't delete an entry: " + err.Error(),
		})
	}

	res := app.GuestbookGuestCollection{}
	for _, row := range rows {
		res = append(res, &app.GuestbookGuest{
			ID:        row.ID,
			Timestamp: float64(row.Timestamp),
			Comment:   row.Comment,
		})
	}

	return ctx.OK(res)
	// GuestbookController_List: end_implement
}
