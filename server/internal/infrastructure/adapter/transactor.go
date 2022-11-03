package adapter

import (
	"context"
	"database/sql"
	"fmt"
)

type txKey struct{}

type Transactor struct {
	db *sql.DB
}

func NewTransactor(db *sql.DB) *Transactor {
	return &Transactor{db: db}
}

func (t Transactor) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx, err := t.db.Begin()
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

	return t.db.QueryContext(ctx, query, args...)
}

func (t Transactor) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.QueryRowContext(ctx, query, args...)
	}

	return t.db.QueryRowContext(ctx, query, args...)
}

func (t Transactor) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	if tx := t.extractTx(ctx); tx != nil {
		return tx.ExecContext(ctx, query, args...)
	}

	return t.db.ExecContext(ctx, query, args...)
}

func (t Transactor) extractTx(ctx context.Context) *sql.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sql.Tx); ok {
		return tx
	}
	return nil
}
