package mssql

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/denisenkom/go-mssqldb" // justifying it
)

var db *sql.DB

func Connect() *sql.DB {
	accessLock := new(sync.RWMutex)

	accessLock.RLock()
	existing := db
	accessLock.RUnlock()
	if existing != nil {
		err := existing.Ping()
		if err != nil {
			existing = nil
		} else {
			return existing
		}
	}
	accessLock.Lock()
	conn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		"127.0.0.1",
		"sa",
		"Secret1234",
		1433,
		"chat_project",
	)
	db, err := sql.Open("sqlserver", conn)
	if err != nil {
		panic("Cannot connect to database" + err.Error())
	}
	accessLock.Unlock()
	return db
}
