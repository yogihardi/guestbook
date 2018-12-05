package servicemock

import (
	"errors"
	"time"

	"github.com/yogihardi/guestbook/model/servicemodel"
)

type ServiceMock struct{}

// Add an entry
func (s ServiceMock) Add(guest servicemodel.GuestBook) error {

	return nil
}

// List list of guestes
func (s ServiceMock) List() ([]servicemodel.GuestBook, error) {
	result := []servicemodel.GuestBook{
		servicemodel.GuestBook{
			ID:        "123",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment",
		},
		servicemodel.GuestBook{
			ID:        "456",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment #2",
		},
	}
	return result, nil
}

// Delete delete an entry
func (s ServiceMock) Delete(ID string) error {
	if ID == "123" {
		return nil
	}
	return errors.New("Data not found")
}
