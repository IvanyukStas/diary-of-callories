package main

import (
	"docc/src/cli"
	"fmt"
	"log"
	"time"
	cdb "docc/db"

)

func main() {
	dbFile := "./db-test.db"
	db, err := cdb.NewDB(dbFile)
	if err != nil {
		log.Fatal(err)
	}
	user := cdb.User{Name: "stas", Registered: time.Now(), Age: 30}
	err = db.AddUser(user)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Привет, давай считать каллории вместе!")
	fmt.Printf("Сегодня %02d-%02d-%d\n", time.Now().Day(), time.Now().Month(), time.Now().Year())
	fmt.Println("Выбирите какой прием пищи записываем завтрак, обед или ужин")

	for {
		r := cli.ReadFromSttin()
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
