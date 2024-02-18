package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type User struct {
	Id int
	Name       string
	Registered string
	Age        int
	Weight     int
	Height     int
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

var schemaDB = `CREATE TABLE IF NOT EXISTS users(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name VARCHAR(32),
	registered TETX,
	age INTEGER,
	weight INTEGER,
	height INTEGER

);
CREATE TABLE IF NOT EXISTS diaries(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	created TIMESTAMP,
	user_id INTEGER,
	FOREIGN KEY (user_id)  REFERENCES users (id)
	);

CREATE TABLE IF NOT EXISTS meals(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	meal_time TEXT,
	meal_ccal INTEGER,
	meal_name TEXT,
	diary_id INTEGET,
	FOREIGN KEY (diary_id)  REFERENCES diaries (id)

);`

type DB struct {
	sql    *sql.DB
	stmt   *sql.Stmt
	buffer []User
}

func NewDB(dbFile string) (*DB, error) {
	sqlDB, err := sql.Open("sqlite3", dbFile)

	check_err(err)
	if _, err := sqlDB.Exec(schemaDB); err != nil {
		return nil, err
	}

	db := DB{
		sql:    sqlDB,
		buffer: make([]User, 0, 10),
	}

	return &db, nil
}

var insertUserSQL = `INSERT INTO users (name, registered, age, weight, height) VALUES (?, ?, ?, ?, ?)`

func (db *DB) AddUser(user User) error {
	tx, err := db.sql.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(insertUserSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Registered, user.Age, user.Weight, user.Height)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (db *DB) GetUser() (*[]User, error) {
	tx, err := db.sql.Begin()
	if err != nil {
		return nil, err
	}
	getUserSQL := "SELECT id, name FROM users"
	rows, err := tx.Query(getUserSQL)
	if err != nil{
		return nil, err
	}
	var users []User
	user := User{}
	for rows.Next(){
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}
	return &users, nil
	}


		


