// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "guest-book": Application User Types
//
// Command:
// $ goagen
// --design=github.com/yogihardi/guestbook/rest/design
// --out=$(GOPATH)/src/github.com/yogihardi/guestbook/rest
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// guestbookType user type.
type guestbookType struct {
	// Comment
	Comment *string `form:"comment,omitempty" json:"comment,omitempty" xml:"comment,omitempty"`
}

// Validate validates the guestbookType type instance.
func (ut *guestbookType) Validate() (err error) {
	if ut.Comment == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "comment"))
	}
	return
}

// Publicize creates GuestbookType from guestbookType
func (ut *guestbookType) Publicize() *GuestbookType {
	var pub GuestbookType
	if ut.Comment != nil {
		pub.Comment = *ut.Comment
	}
	return &pub
}

// GuestbookType user type.
type GuestbookType struct {
	// Comment
	Comment string `form:"comment" json:"comment" xml:"comment"`
}

// Validate validates the GuestbookType type instance.
func (ut *GuestbookType) Validate() (err error) {
	if ut.Comment == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "comment"))
	}
	return
}

// version user type.
type version struct {
	// Git commit hash
	Git *string `form:"git,omitempty" json:"git,omitempty" xml:"git,omitempty"`
	// Application version
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
}

// Validate validates the version type instance.
func (ut *version) Validate() (err error) {
	if ut.Version == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "version"))
	}
	return
}

// Publicize creates Version from version
func (ut *version) Publicize() *Version {
	var pub Version
	if ut.Git != nil {
		pub.Git = ut.Git
	}
	if ut.Version != nil {
		pub.Version = *ut.Version
	}
	return &pub
}

// Version user type.
type Version struct {
	// Git commit hash
	Git *string `form:"git,omitempty" json:"git,omitempty" xml:"git,omitempty"`
	// Application version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the Version type instance.
func (ut *Version) Validate() (err error) {
	if ut.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "version"))
	}
	return
}