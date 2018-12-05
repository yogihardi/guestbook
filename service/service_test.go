package service

import (
	"reflect"
	"testing"
	"time"

	"github.com/yogihardi/guestbook/model/servicemodel"
	"github.com/yogihardi/guestbook/service/daomock"
	"golang.org/x/net/context"
)

var serviceTest Service

func init() {
	serviceTest, _ = NewService(context.Background(), daomock.DaoMock{})
}

func Test_service_Add(t *testing.T) {
	type args struct {
		guest servicemodel.GuestBook
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test add",
			args: args{
				guest: servicemodel.GuestBook{
					ID:        "123",
					Timestamp: time.Now().Unix(),
					Comment:   "test comment",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := serviceTest
			if err := s.Add(tt.args.guest); (err != nil) != tt.wantErr {
				t.Errorf("service.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_List(t *testing.T) {
	tests := []struct {
		name    string
		want    []servicemodel.GuestBook
		wantErr bool
	}{
		{
			name: "test list",
			want: []servicemodel.GuestBook{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := serviceTest
			got, err := s.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("service.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Delete(t *testing.T) {
	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test delete",
			args:    args{ID: "123"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := serviceTest
			if err := s.Delete(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("service.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
