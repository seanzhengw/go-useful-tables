package user

import (
	"database/sql"
)

// CreateDetailStatement of the user_detail table with (id, fullname, email)
const CreateDetailStatement string = `
CREATE TABLE IF NOT EXISTS user_detail (
	id INT UNSIGNED NOT NULL, 
	fullname VARCHAR(100) NULL,
	email VARCHAR(100) NULL,
	PRIMARY KEY (id),
	CONSTRAINT user_detail_id FOREIGN KEY (id) REFERENCES user (id) ON UPDATE CASCADE ON DELETE CASCADE
)`

// CreateDetail create user_detail table with (id, fullname, email) if not exist
func CreateDetail(sqldb *sql.DB) (sql.Result, error) {
	return sqldb.Exec(CreateDetailStatement)
}
