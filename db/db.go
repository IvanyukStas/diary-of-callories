package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Name       string
	registered time.Time
	age        int
	wieght     int
	height     int
}

type Diary struct {
	Created time.Time
	UserId  User
}

type Meal struct {
	MealTime string
	MealName []string
	MealCcal int
	created  time.Time
	DiaryId  Diary
}
// created TIMESTAMP,
var schemaDB = `CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(32),
	
	age INTEGER,
	weight INTEGER,
	height INTEGER

);
CREATE TABLE IF NOT EXISTS diaries(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created TIMESTAMP,
	user_id INTEGER,
	FOREIGN KEY (user_id)  REFERENCES users (id)

);`

type DB struct {
	sql    *sql.DB
	stmt   *sql.Stmt
	buffer []User
}

func NewDB(dbFile string) (*DB, error) {
	sqlDB, err := sql.Open("sqlite3", dbFile)
	fmt.Println("1111111111")

	check_err(err)
	if _, err := sqlDB.Exec(schemaDB); err != nil {
		return nil, err
	}

	insertSQL := `
	INSERT INTO users (
	name, age, weight, height
	 ) VALUES (
	     ?, ?, ?, ?
	 )`

	stmt, err := sqlDB.Prepare(insertSQL)
	if err != nil {
		return nil, err
	}

	var db DB

	db = DB{
		sql:    sqlDB,
		stmt:   stmt,
		buffer: make([]User, 0, 10),
	}

	return &db, nil
}
