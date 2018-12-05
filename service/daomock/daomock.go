package daomock

import (
	"time"

	"github.com/yogihardi/guestbook/model/daomodel"
)

type DaoMock struct{}

func (m DaoMock) Add(guest daomodel.GuestBook) error {
	return nil
}
func (m DaoMock) List() ([]daomodel.GuestBook, error) {
	result := []daomodel.GuestBook{
		daomodel.GuestBook{
			ID:        "123",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment",
		},
		daomodel.GuestBook{
			ID:        "456",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment #2",
		},
	}
	return result, nil
}
func (m DaoMock) Delete(ID string) error {
	return nil
}
