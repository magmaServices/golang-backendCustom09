package database

import (
	"github.com/gocraft/dbr"
	"localhost/go-heroes/fesl-backend/backend/model"
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
