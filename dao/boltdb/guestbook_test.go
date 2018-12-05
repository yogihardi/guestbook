package boltdb

import (
	"reflect"
	"testing"
	"time"

	"github.com/yogihardi/guestbook/model/daomodel"
)

func TestBoltDB_Add(t *testing.T) {
	type args struct {
		guest daomodel.GuestBook
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test add",
			args: args{
				guest: daomodel.GuestBook{
					ID:        "123",
					Timestamp: time.Now().Unix(),
					Comment:   "test comment",
				},
			},
			wantErr: false,
		},
		{
			name: "test add 2",
			args: args{
				guest: daomodel.GuestBook{
					ID:        "123456",
					Timestamp: time.Now().Unix(),
					Comment:   "test comment #2",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := getDaoTest()
			if err := dao.Add(tt.args.guest); (err != nil) != tt.wantErr {
				t.Errorf("BoltDB.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBoltDB_List(t *testing.T) {
	daotest := getDaoTest()
	daotest.Add(
		daomodel.GuestBook{
			ID:        "123",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment",
		},
	)

	daotest.Add(
		daomodel.GuestBook{
			ID:        "123456",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment #2",
		},
	)

	tests := []struct {
		name    string
		want    []daomodel.GuestBook
		wantErr bool
	}{
		{
			name: "test list",
			want: []daomodel.GuestBook{
				daomodel.GuestBook{
					ID:        "123",
					Timestamp: time.Now().Unix(),
					Comment:   "test comment",
				},
				daomodel.GuestBook{
					ID:        "123456",
					Timestamp: time.Now().Unix(),
					Comment:   "test comment #2",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := daotest
			got, err := dao.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("BoltDB.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BoltDB.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBoltDB_Delete(t *testing.T) {
	daotest := getDaoTest()
	daotest.Add(
		daomodel.GuestBook{
			ID:        "123",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment",
		},
	)

	daotest.Add(
		daomodel.GuestBook{
			ID:        "123456",
			Timestamp: time.Now().Unix(),
			Comment:   "test comment #2",
		},
	)

	type args struct {
		ID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test delete 1",
			args:    args{ID: "123"},
			wantErr: false,
		},
		{
			name:    "test delete non exists",
			args:    args{ID: "no"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := daotest
			if err := dao.Delete(tt.args.ID); (err != nil) != tt.wantErr {
				t.Errorf("BoltDB.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
