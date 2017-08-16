package db

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

//UnitOfWork wrapper tx
type UnitOfWork interface {
	MustNamedExec(query string, arg interface{}) sql.Result

	Query(query string, args ...interface{}) (*sqlx.Rows, error)

	Select(dest interface{}, query string, args ...interface{}) error

	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)

	MustExec(query string, args ...interface{}) sql.Result

	Get(dest interface{}, query string, args ...interface{}) error

	InTransaction(contextOver func(db UnitOfWork) error)
}

type unitOfWork struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

//NewUnitOfWork factory method
func NewUnitOfWork(db *sqlx.DB, tx *sqlx.Tx) UnitOfWork {
	return &unitOfWork{db: db, tx: tx}
}

func (u *unitOfWork) InTransaction(contextOver func(db UnitOfWork) error) {
	u.begin()

	defer func() {
		if r := recover(); r != nil {
			u.rollback()
			panic(r)
		}
	}()

	err := contextOver(u)

	if err == nil {
		u.commit()
	} else {
		u.rollback()
	}
}

func (u *unitOfWork) MustNamedExec(query string, arg interface{}) sql.Result {
	if u.tx != nil {
		res, err := u.tx.NamedExec(query, arg)
		if err != nil {
			panic(err)
		}

		return res
	}

	res, err := u.db.NamedExec(query, arg)
	if err != nil {
		panic(err)
	}

	return res
}

func (u *unitOfWork) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	if u.tx != nil {
		return u.tx.Queryx(query, args...)
	}

	return u.db.Queryx(query, args...)
}

func (u *unitOfWork) Select(dest interface{}, query string, args ...interface{}) error {
	if u.tx != nil {
		return u.tx.Select(dest, query, args...)
	}

	return u.db.Select(dest, query, args...)
}

func (u *unitOfWork) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	if u.tx != nil {
		return u.tx.NamedQuery(query, arg)
	}

	return u.db.NamedQuery(query, arg)
}

func (u *unitOfWork) MustExec(query string, args ...interface{}) sql.Result {
	if u.tx != nil {
		return u.tx.MustExec(query, args...)
	}

	return u.db.MustExec(query, args...)
}

func (u *unitOfWork) Get(dest interface{}, query string, args ...interface{}) error {
	if u.tx != nil {
		return u.tx.Get(dest, query, args...)
	}

	return u.db.Get(dest, query, args...)
}

func (u *unitOfWork) begin() {
	u.tx = u.db.MustBegin()

	if u.tx == nil {
		panic(errors.New("Nenhuma transação foi iniciada."))
	}
}

func (u *unitOfWork) commit() {
	if u.tx == nil {
		panic(errors.New("Nenhuma transação foi iniciada."))
	}

	err := u.tx.Commit()
	if err != nil {
		panic(err)
	}

	u.tx = nil
}

func (u *unitOfWork) rollback() {
	if u.tx == nil {
		panic(errors.New("Nenhuma transação foi iniciada."))
	}

	err := u.tx.Rollback()
	if err != nil {
		panic(err)
	}

	u.tx = nil
}
