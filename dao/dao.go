package dao

import (
	"github.com/yogihardi/guestbook/model/daomodel"
)

// Dao ..
type Dao interface {
	Add(guest daomodel.GuestBook) error
	List() ([]daomodel.GuestBook, error)
	Delete(ID string) error
}
