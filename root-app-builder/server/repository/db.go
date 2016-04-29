package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

/*
 * Initiates connections to the database and initializes the
 * schema.
 */
func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	/*
	 * Initializes the schema.
	 */
	return GetRepository().CreateSchema()
}

/*
 * Sets the used datasource to the given one.
 */
func SetDB(dbObject *sql.DB) {
	db = dbObject
}
