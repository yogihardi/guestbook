package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// swagger
var _ = Resource("swagger", func() {
	NoSecurity()
	Origin("*", func() {
		Methods("GET")
	})
	Files("/guestbook/swagger.json", "rest/swagger/swagger.json")
})

// version
var VersionType = Type("version", func() {
	Attribute("version", String, "Application version", func() {
		Example("1.0")
	})
	Attribute("git", String, "Git commit hash", func() {
		Example("000000")
	})

	Required("version")
})

var _ = Resource("version", func() {
	Action("version", func() {
		Routing(GET("version"))
		Response(OK, VersionMedia)
		Metadata("swagger:summary", "Return application's version and commit hash")
	})
})

// GuestbookType guest book payload
var GuestbookType = Type("guestbookType", func() {
	Attribute("comment", String, "Comment", func() {
		Example("Lorem ipsum")
	})

	Required("comment")
})

var _ = Resource("guestbook", func() {
	BasePath("/")

	Action("add", func() {
		Routing(POST("/"))
		Payload(GuestbookType)
		Response(NoContent)
		Response(BadRequest, GuestBookErrorMedia)
		Response(InternalServerError, GuestBookErrorMedia)
		Metadata("swagger:summary", "Add an entry")
	})

	Action("list", func() {
		Routing(GET("/"))
		Response(OK, CollectionOf(GuestMedia))
		Response(BadRequest, GuestBookErrorMedia)
		Response(InternalServerError, GuestBookErrorMedia)
		Metadata("swagger:summary", "List of guests")
	})

	Action("delete", func() {
		Routing(DELETE("/:id"))
		Params(func() {
			Param("id", String, "Entry ID", func() {
				Example("c2d9ecce-f6a3-49cf-8c71-7d4415beb3b8")
			})
		})
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, GuestBookErrorMedia)
		Response(InternalServerError, GuestBookErrorMedia)
		Metadata("swagger:summary", "Delete an entry")
	})
})
