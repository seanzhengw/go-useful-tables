package session

import (
	"database/sql"
)

// CreateStatement of the session table with (id, user_id, ip_address, last_activity, user_agent)
const CreateStatement string = `
CREATE TABLE session (
	id BINARY(64) NOT NULL,
	user_id INT UNSIGNED NOT NULL, 
	ip_address BINARY(16) NOT NULL,
	last_activity DATETIME NOT NULL,
	user_agent TEXT NOT NULL,
	PRIMARY KEY (id),
	INDEX session_user_id (user_id),
	CONSTRAINT session_user_id FOREIGN KEY (user_id) REFERENCES user (id) ON UPDATE CASCADE ON DELETE CASCADE
)`

// CreateTable create session table with (id, user_id, ip_address, last_activity, user_agent)
func CreateTable(sqldb *sql.DB) (sql.Result, error) {
	return sqldb.Exec(CreateStatement)
}
