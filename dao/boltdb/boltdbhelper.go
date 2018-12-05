package boltdb

import (
	"encoding/json"
	"errors"

	"github.com/boltdb/bolt"
)

// BoltDBCtx is opaque structure containing BoltDB context
type (
	BoltDBCtx struct {
		Db *bolt.DB
	}

	Filter struct {
		FilterKey   string
		FilterValue string
	}
)

var (
	Tables = []string{
		TableGuestbook,
	}
)

// NewBoltDBCtx creates an instance of BoltDB backed Dao
func NewBoltDBCtx(db *bolt.DB) (*BoltDBCtx, error) {
	boltdbCtx := &BoltDBCtx{
		Db: db,
	}

	// create buckets
	boltdbCtx.CreateBuckets(Tables)

	return boltdbCtx, nil
}

type BoltdbCtxInterface interface {
	IsBucketExists(bucketName string) bool
	CreateBuckets(bucketNames []string) error
	Get(bucketName string, key string) ([]byte, error)
	Put(bucketName string, key string, value []byte) error
	Delete(bucketName string, key string) error
	View(bucketName string, filter Filter) ([][]byte, error)
	DeleteBucket(bucketName string) error
	DeleteWithoutChecking(bucketName string, key string) error
	RecreateBucket(bucketName string) error
}

func (db *BoltDBCtx) IsBucketExists(bucketName string) bool {
	tx, err := db.Db.Begin(true)
	if err != nil {
		return false
	}
	defer tx.Rollback()

	if bucket := tx.Bucket([]byte(bucketName)); bucket != nil {
		return true
	}

	return false
}

func (db *BoltDBCtx) CreateBuckets(bucketNames []string) error {
	tx, err := db.Db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	for _, tbl := range bucketNames {
		bucket := tx.Bucket([]byte(tbl))
		if bucket == nil {
			if _, err := tx.CreateBucketIfNotExists([]byte(tbl)); err != nil {
				return err
			}
		}
	}

	tx.Commit()

	return nil
}

func (db *BoltDBCtx) Get(bucketName string, key string) ([]byte, error) {
	var value []byte
	err := db.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(key))
		if v != nil {
			value = append(value, b.Get([]byte(key))...)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, errors.New("Data not found")
	}

	return value, nil
}

func (db *BoltDBCtx) Put(bucketName string, key string, value interface{}) error {
	if key == "" {
		return bolt.ErrKeyRequired
	}

	err := db.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))

		encodedValue, err := json.Marshal(value)
		if err != nil {
			return nil
		}

		err = b.Put([]byte(key), encodedValue)
		return err
	})

	return err
}

func (db *BoltDBCtx) Delete(bucketName string, key string) error {
	if _, err := db.Get(bucketName, key); err != nil {
		return err
	}

	err := db.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Delete([]byte(key))
		return err
	})

	return err
}

func (db *BoltDBCtx) DeleteWithoutChecking(bucketName string, key string) error {
	err := db.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Delete([]byte(key))
		return err
	})

	return err
}

func (db *BoltDBCtx) View(bucketName string) ([][]byte, error) {
	var result [][]byte
	err := db.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		b.ForEach(func(k, v []byte) error {
			// Due to
			// Byte slices returned from Bolt are only valid during a transaction. Once the transaction has been committed or rolled back then the memory they point to can be reused by a new page or can be unmapped from virtual memory and you'll see an unexpected fault address panic when accessing it.
			// We copy the slice to retain it
			dstv := make([]byte, len(v))
			copy(dstv, v)

			result = append(result, dstv)
			return nil
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (db *BoltDBCtx) DeleteBucket(bucketName string) error {
	tx, err := db.Db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	tx.DeleteBucket([]byte(bucketName))
	tx.Commit()

	return nil
}

func (db *BoltDBCtx) RecreateBucket(bucketName string) error {
	tx, err := db.Db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	tx.DeleteBucket([]byte(bucketName))
	tx.CreateBucket([]byte(bucketName))
	tx.Commit()

	return nil
}
