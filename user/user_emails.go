package user

import (
	"database/sql"
)

// CreateEmailsStatement of the user_emails table with (id, user_id, email)
//
// for the case with user may have multiple emails, and there is another table for users primary email
const CreateEmailsStatement string = `
CREATE TABLE user_emails (
	id INT UNSIGNED NOT NULL AUTO_INCREMENT, 
	user_id INT UNSIGNED NOT NULL, 
	email VARCHAR(100) NOT NULL,
	PRIMARY KEY (id),
	CONSTRAINT user_emails_id FOREIGN KEY (user_id) REFERENCES user (id) ON UPDATE CASCADE ON DELETE CASCADE
)`

// CreateEmailsStatementWithoutID of the user_emails table with (id, email)
//
// for the case with user may have multiple emails
const CreateEmailsStatementWithoutID string = `
CREATE TABLE user_emails (
	id INT UNSIGNED NOT NULL, 
	email VARCHAR(100) NOT NULL,
	CONSTRAINT user_emails_id FOREIGN KEY (id) REFERENCES user (id) ON UPDATE CASCADE ON DELETE CASCADE
)`

// CreateEmails create user_emails table with (id, user_id, email)
//
// for the case with user may have multiple emails, and there is another table for users primary email
func CreateEmails(sqldb *sql.DB) (sql.Result, error) {
	return sqldb.Exec(CreateEmailsStatement)
}

// CreateEmailsWithoutID create user_emails table with (id, email)
//
// for the case with user may have multiple emails
func CreateEmailsWithoutID(sqldb *sql.DB) (sql.Result, error) {
	return sqldb.Exec(CreateEmailsStatementWithoutID)
}
