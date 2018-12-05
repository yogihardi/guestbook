package service

import (
	"github.com/yogihardi/guestbook/dao"
	"github.com/yogihardi/guestbook/model/daomodel"
	"github.com/yogihardi/guestbook/model/servicemodel"
	"golang.org/x/net/context"

	"github.com/inconshreveable/log15"
)

type service struct {
	logger log15.Logger
	ctx    context.Context
	dao    dao.Dao
}

// Service list of service functions
type Service interface {
	// Add an entry
	Add(guest servicemodel.GuestBook) error
	// List list of guestes
	List() ([]servicemodel.GuestBook, error)
	// Delete delete an entry
	Delete(ID string) error
}

// NewService initiate service
func NewService(ctx context.Context, dao dao.Dao) (Service, error) {
	logger := log15.New("pkg", "service")

	return &service{
		logger: logger,
		ctx:    ctx,
		dao:    dao,
	}, nil
}

func (s *service) Add(guest servicemodel.GuestBook) error {
	return s.dao.Add(daomodel.GuestBook{
		ID:        guest.ID,
		Timestamp: guest.Timestamp,
		Comment:   guest.Comment,
	})
}

func (s *service) List() ([]servicemodel.GuestBook, error) {
	rows, err := s.dao.List()
	if err != nil {
		return nil, err
	}

	result := []servicemodel.GuestBook{}
	for _, row := range rows {
		result = append(result, servicemodel.GuestBook{
			ID:        row.ID,
			Comment:   row.Comment,
			Timestamp: row.Timestamp,
		})
	}

	return result, nil
}

func (s *service) Delete(ID string) error {
	return s.dao.Delete(ID)
}
