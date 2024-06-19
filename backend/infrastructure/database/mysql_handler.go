package database

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/yosuke7040/dine-out-discoveries/adapter/repository"
)

type mySqlHandler struct {
	db *sql.DB
}

func NewMySqlHandler(c *config) (*mySqlHandler, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	conf := mysql.Config{
		DBName:               c.database,
		User:                 c.user,
		Passwd:               c.password,
		Addr:                 c.host + ":" + c.port,
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		slog.Error("cloud not open db:", c.database, err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		slog.Error("cloud not ping db:", c.database, err)
		return nil, err
	}

	return &mySqlHandler{db: db}, nil
}

func (m mySqlHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := m.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return mysqlTx{}, err
	}

	return newMySQLTx(tx), nil
}

func (m mySqlHandler) ExecuteContext(ctx context.Context, query string, args ...any) error {
	_, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m mySqlHandler) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newMySQLRows(rows)

	return row, nil
}

func (m mySqlHandler) QueryRowContext(ctx context.Context, query string, args ...any) repository.Row {
	row := m.db.QueryRowContext(ctx, query, args...)

	return newMySQLRow(row)
}

type mysqlRow struct {
	row *sql.Row
}

func newMySQLRow(row *sql.Row) mysqlRow {
	return mysqlRow{row: row}
}

func (mr mysqlRow) Scan(dest ...any) error {
	if err := mr.row.Scan(dest...); err != nil {
		return err
	}

	return nil
}

type mysqlRows struct {
	rows *sql.Rows
}

func newMySQLRows(rows *sql.Rows) mysqlRows {
	return mysqlRows{rows: rows}
}

func (mr mysqlRows) Scan(dest ...any) error {
	if err := mr.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

func (mr mysqlRows) Next() bool {
	return mr.rows.Next()
}

func (mr mysqlRows) Err() error {
	return mr.rows.Err()
}

func (mr mysqlRows) Close() error {
	return mr.rows.Close()
}

type mysqlTx struct {
	tx *sql.Tx
}

func newMySQLTx(tx *sql.Tx) mysqlTx {
	return mysqlTx{tx: tx}
}

func (m mysqlTx) ExecuteContext(ctx context.Context, query string, args ...any) error {
	_, err := m.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (m mysqlTx) QueryContext(ctx context.Context, query string, args ...any) (repository.Rows, error) {
	rows, err := m.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newMySQLRows(rows)

	return row, nil
}

func (m mysqlTx) QueryRowContext(ctx context.Context, query string, args ...any) repository.Row {
	row := m.tx.QueryRowContext(ctx, query, args...)

	return newMySQLRow(row)
}

func (m mysqlTx) Commit() error {
	return m.tx.Commit()
}

func (m mysqlTx) Rollback() error {
	return m.tx.Rollback()
}
