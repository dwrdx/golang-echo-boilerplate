package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var mySQLPool *sql.DB

// Init inits the database
func Init() error {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URI"))
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	mySQLPool = db
	return nil
}

// GetMySQLPool gets the mysql pool
func GetMySQLPool() *sql.DB {
	return mySQLPool
}

// Insert inserts data to database and return the id
func Insert(db *sql.DB, sqlString string, args ...interface{}) (int64, error) {

	stmt, err := db.Prepare(sqlString)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
