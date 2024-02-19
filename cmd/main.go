package main

import (
	cdb "docc/db"
	"docc/src/cli"
	"fmt"
	"log"
	"time"
)

var users *[]cdb.User
var db *cdb.DB


func main() {	
	dbFile := "./db-test.db"
	db, err := cdb.NewDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	users, err := db.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CLose()
	// user := cdb.User{Name: "stas", Registered: time.Now().Format("02-01-2006"), Age: 30}
	// err = db.AddUser(user)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	
	for _, user := range *users {
		fmt.Printf("%v %v\t", user.Id, user.Name)
	}

	fmt.Println("Привет, давай считать каллории вместе!")
	fmt.Printf("Сегодня %v\n", time.Now().Format("02-01-2006"))
	fmt.Println("Выберите пользователя с которым будем работать!")

	// fmt.Println("Выбирите какой прием пищи записываем завтрак, обед или ужин")

	for {
		r := cli.ReadFromSttin()
		if r == ""{
			println(" Выберите значение, пустая строка не поддерживается")
			continue

		}
		if r == "q" {
			break
		}
		fmt.Println(r)
	}

}

// func run() error {
// 	log.Println("Start program!")

// 	time.Sleep(5 * time.Second)
// 	defer log.Println("program is closed!")
// 	return nil
// }
