package user

import (
	"database/sql"
)

// CreateStatement of the user table with (id, name, pass)
const CreateStatement string = `
	CREATE TABLE IF NOT EXISTS user (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(60) NOT NULL,
		pass BINARY(60) NOT NULL,
		PRIMARY KEY (id),
		UNIQUE INDEX name (name)
	)`

// CreateTable create user table with (id, name, pass) if not exist
func CreateTable(sqldb *sql.DB) (sql.Result, error) {
	return sqldb.Exec(CreateStatement)
}
