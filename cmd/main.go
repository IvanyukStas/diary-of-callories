package main

import (
	cdb "docc/db"
	"fmt"
	// "docc/src/cli"
	"docc/utilsfunc"
	// "fmt"
	// "log"
	// "time"
)

var users *[]cdb.User
var db *cdb.DB

func main() {
	slcmeals, err := utilsfunc.AddMealToSlice("гречка 100, uhtxrf 200, dfdsfds sdffds, fdsfsfsd 3333")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(slcmeals)

	// var user cdb.User
	// user.Id = 1
	// user.Name = "Stas"

	// dbFile := "./db-test.db"
	// db, err := cdb.NewDB(dbFile)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // users, err := db.GetUsers()
	// // if err != nil {
	// // 	log.Fatal(err)
	// // }
	// // for _, u := range *users {
	// // 	user.Id, user.Name = u.Id, u.Name
	// // 	fmt.Printf("%v %v\t", u.Id, u.Name)
	// // }

	// defer db.CLose()

	// meal := cdb.Meal{
	// 	MealTime: utilsfunc.CheckMealTime(),
	// 	MealName: "авсянка",
	// 	MealCcal: 100,
	// 	MealDate: time.Now().Format("02-01-2006"),
	// 	UserId:   user.Id}

	// err = db.AddMeal(meal)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// user = cdb.User{Name: "stas", Registered: time.Now().Format("02-01-2006"), Age: 30}
	// err = db.AddUser(user)
	// if err != nil {
	// 	log.Fatal(err)
	//  }
	// fmt.Printf("Привет %v, начнем считать калории вместе!\n", user.Name)
	// fmt.Printf("Сегодня %v Время приема пищи: %v \n", time.Now().Format("02-01-2006"), utilsfunc.CheckMealTime())
	// fmt.Println("Давай заполним дневник калорий!")
	// fmt.Println("Заполни по одиночке каждое блюдо или через запятую")
	// fmt.Println("К примеру гречка 300 или гречка 300, свинная котлета 150, салат витаминный 150")
	// fmt.Println("Цифры после названия блюда это граммы")
	// fmt.Println()

	// fmt.Println("Выбирите какой прием пищи записываем завтрак, обед или ужин")

	// for {
	// 	r := cli.ReadFromSttin()
	// 	if r == "" {
	// 		println("Выберите значение, пустая строка не поддерживается")
	// 		continue

	// 	}
	// 	if r == "q" {
	// 		break
	// 	}
	// 	fmt.Println(r)
	// }

}

// func run() error {
// 	log.Println("Start program!")

// 	time.Sleep(5 * time.Second)
// 	defer log.Println("program is closed!")
// 	return nil
// }
