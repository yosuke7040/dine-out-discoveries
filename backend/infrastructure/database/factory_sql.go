package database

import (
	"errors"

	"github.com/yosuke7040/dine-out-discoveries/adapter/repository"
)

var (
	errInvalidSqlDatabaseInstance = errors.New("invalid sql database instance")
)

const (
	InstanceMySql int = iota
	InstancePostgres
)

func NewDatabaseSqlFactory(instance int) (repository.Sql, error) {
	switch instance {
	case InstanceMySql:
		return NewMySqlHandler(newConfigMySql())
	case InstancePostgres:
		return NewPostgresHandler(newConfigPostgres())
	default:
		return nil, errInvalidSqlDatabaseInstance
	}
}
