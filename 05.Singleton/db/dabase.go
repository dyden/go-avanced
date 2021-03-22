package database

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (Database) CreateSingleConnection() {
	fmt.Println("Creating Singleton for database")
	time.Sleep(time.Second * 2)
	fmt.Println("Connection done")
}

var db *Database
var locker sync.Mutex

func GetDatabaseInstance() *Database {
	locker.Lock()
	defer locker.Unlock()
	if db == nil {
		fmt.Println("Creating DB Connection")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("DB Connection already exists")
	}
	return db
}
