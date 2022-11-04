package adapter

import (
	"context"
	"database/sql"
	"fmt"
)

type txKey struct{}

type Transactor struct {
	Db *sql.DB
}

func NewTransactor(db *sql.DB) *Transactor {
	return &Transactor{Db: db}
}

func (t Transactor) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx, err := t.Db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback()

	err = tFunc(context.WithValue(ctx, txKey{}, tx))
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (t Transactor) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.QueryContext(ctx, query, args...)
	}

	return t.Db.QueryContext(ctx, query, args...)
}

func (t Transactor) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return t.Db.QueryRowContext(ctx, query, args...)
}

func (t Transactor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}

	return t.Db.ExecContext(ctx, query, args...)
}

func (t Transactor) extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}
