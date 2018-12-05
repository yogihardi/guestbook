package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// GuestBookErrorMedia ..
var GuestBookErrorMedia = MediaType("application/vnd.guestbook.error+json", func() {
	ContentType("application/json")
	Attribute("code", String, "Message ID", func() {
		Example("UNAUTHORIZED")
	})
	Attribute("msg", String, "Localized message", func() {
		Example("Unauthorized access")
	})

	View("default", func() {
		Attribute("code")
		Attribute("msg")
	})

	Required("code", "msg")
})

// VersionMedia ..
var VersionMedia = MediaType("application/vnd.guestbook.version+json", func() {
	ContentType("application/json")
	Attribute("version", String, "Application version", func() {
		Example("1.0")
	})
	Attribute("git", String, "Git commit hash", func() {
		Example("000000")
	})
	View("default", func() {
		Attribute("version")
		Attribute("git")
	})
	Required("version")
})

// GuestMedia ..
var GuestMedia = MediaType("application/vnd.guestbook.guest+json", func() {
	ContentType("application/json")
	Attribute("id", String, "Guestbook ID", func() {
		Example("c2d9ecce-f6a3-49cf-8c71-7d4415beb3b8")
	})
	Attribute("timestamp", Number, "Entry timestamp", func() {
		Example(1543981510)
	})
	Attribute("comment", String, "Comment", func() {
		Example("Lorem ipsum")
	})

	View("default", func() {
		Attribute("id")
		Attribute("timestamp")
		Attribute("comment")
	})

	Required("id", "timestamp", "comment")
})
