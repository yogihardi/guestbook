package controller

import (
	"context"
	"testing"

	"github.com/goadesign/goa"
	"github.com/yogihardi/guestbook/rest/app"
	"github.com/yogihardi/guestbook/rest/app/test"
	"github.com/yogihardi/guestbook/rest/controller/servicemock"
)

var (
	serviceMock servicemock.ServiceMock
	serviceTest *goa.Service
	gbCtl       *GuestbookController
	ctx         context.Context
)

func init() {
	serviceMock = servicemock.ServiceMock{}
	serviceTest = goa.New("controller-test")
	gbCtl = NewGuestbookController(serviceTest, serviceMock)
	ctx = context.Background()
}

func TestGuestbookController_Add(t *testing.T) {
	test.AddGuestbookNoContent(t, ctx, serviceTest, gbCtl, &app.GuestbookType{
		Comment: "comment test",
	})
}

func TestGuestbookController_Delete(t *testing.T) {
	test.DeleteGuestbookNoContent(t, ctx, serviceTest, gbCtl, "123")
	test.DeleteGuestbookNotFound(t, ctx, serviceTest, gbCtl, "abc")
}

func TestGuestbookController_List(t *testing.T) {
	test.ListGuestbookOK(t, ctx, serviceTest, gbCtl)
}
