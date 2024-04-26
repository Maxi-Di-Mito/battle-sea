package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var client *sql.DB

func InitDb() *sql.DB {

	if client != nil {
		return client
	}

	var err error

	client, err = sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err)
	}

	fmt.Println("CONECTED")
	const create string = `
  CREATE TABLE IF NOT EXISTS activities (
  id INTEGER NOT NULL PRIMARY KEY,
  time DATETIME NOT NULL,
  description TEXT
  );`

	client.Exec(create)

	return client

}
