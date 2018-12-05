package boltdb

import (
	"io/ioutil"
	"os"

	"github.com/boltdb/bolt"
	"github.com/inconshreveable/log15"
)

var logger log15.Logger

func init() {
	logger = log15.New("module", "test")
}

func NewTestDB() *BoltDBCtx {
	dir, err := ioutil.TempDir("", "guestbook")
	if err != nil {
		logger.Error("failed to get temp dir", "err", err)
	}
	dbPath := dir + "/test-guestbook.boltdb"

	if _, err := os.Stat(dbPath); os.IsExist(err) {
		os.Remove(dbPath)
	}

	db, err := bolt.Open(dbPath, 0600, nil)
	if err != nil {
		logger.Error("failed to open dbfile", "err", err)
	}

	// Return wrapped type.
	ctx, err := NewBoltDBCtx(db)
	if err != nil {
		logger.Error("failed to open dbfile", "err", err)
	}

	return ctx
}

func getDaoTest() *BoltDB {

	boltdbCtxTest := NewTestDB()
	dao := &BoltDB{
		ctx:    boltdbCtxTest,
		logger: logger,
	}
	//defer boltdbCtxTest.close()

	return dao
}
