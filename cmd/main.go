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
	defer db.CLose()
	// user := cdb.User{Name: "stas", Registered: time.Now().Format("02-01-2006"), Age: 30}
	// err = db.AddUser(user)
	// if err != nil{
	// 	log.Fatal(err)
	// }
	us, err:= db.GetUser()
	if err !=nil{
		log.Fatal(err)
	}
	for _, user := range *us{
		fmt.Println(user.Id, user.Name)
	}
	

	fmt.Println("Привет, давай считать каллории вместе!")
	fmt.Printf("Сегодня %v\n", time.Now().Format("02-01-2006"))
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
