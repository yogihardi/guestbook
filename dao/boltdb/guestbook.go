package boltdb

import (
	"encoding/json"

	"github.com/goadesign/goa/uuid"
	"github.com/yogihardi/guestbook/model/daomodel"
)

// TableGuestbook guestbook table name
const TableGuestbook = "guestbook"

// Add add an entry
func (dao BoltDB) Add(guest daomodel.GuestBook) error {
	if guest.ID == "" {
		guest.ID = uuid.NewV4().String()
	}

	if err := dao.ctx.Put(TableGuestbook, guest.ID, guest); err != nil {
		dao.logger.Error("failed to add an entry", "err", err)
		return err
	}

	return nil
}

// List list of guestes
func (dao BoltDB) List() ([]daomodel.GuestBook, error) {
	rows, err := dao.ctx.View(TableGuestbook)
	if err != nil {
		dao.logger.Error("failed to get guestbooks", "err", err)
		return nil, err
	}

	result := []daomodel.GuestBook{}
	for _, row := range rows {
		var gb daomodel.GuestBook
		if err := json.Unmarshal(row, &gb); err != nil {
			dao.logger.Error("failed unmarshal json", "err", err, "data", string(row))
			continue
		}

		result = append(result, gb)
	}

	return result, nil
}

// Delete delete an entry
func (dao BoltDB) Delete(ID string) error {
	if err := dao.ctx.Delete(TableGuestbook, ID); err != nil {
		dao.logger.Error("failed to delete", "err", err, "id", ID)
		return err
	}
	return nil
}
