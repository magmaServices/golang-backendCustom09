package database

import (
	"github.com/gocraft/dbr"
	"gitlab.com/oiacow/fesl3/backend/model"
)

type Adapter interface {
	TxAdapter
	model.QueriesAdapter
}

type TxAdapter interface {
	NewSession() *dbr.Session

	Begin(sess *dbr.Session) (*dbr.Tx, error)
	Commit(tx *dbr.Tx) error
	Rollback(tx *dbr.Tx) error
}
