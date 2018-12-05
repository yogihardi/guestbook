package boltdb

import (
	"github.com/inconshreveable/log15"
)

// BoltDB ..
type BoltDB struct {
	ctx    *BoltDBCtx
	logger log15.Logger
}

// New creates new BoltDB instance
func New(ctx *BoltDBCtx) BoltDB {
	return BoltDB{
		ctx:    ctx,
		logger: log15.New("module", "dao"),
	}
}
