package utilsfunc

import (
	"docc/db"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CheckMealTime() string {
	timer := time.Now().Hour()
	switch  {
	case (timer >= 6) && (timer < 13):
		return "Завтрак"
	case (timer >= 13) && (timer < 18):
		return "Обед"
	default:
		return "Ужин"

	}
}

func AddMealToSlice(r string)(slcmeals []db.Meal, err error){
	if len(r) == 0{
		return nil, errors.New("пустая строка! паполните строку!")
	}
	slcstr := strings.Split(r, ",")
	for _, strMeal := range slcstr{
		slcMeal := strings.Split(strings.TrimSpace(strMeal), " ")
		ccalInt, err := strconv.Atoi(slcMeal[1])
		if err != nil {
			fmt.Println("Не смог записать ", slcMeal[0])
			fmt.Println("После блюда "+ slcMeal[0] + " укажите граммовку цифрами")
		}
		slcmeals = append(slcmeals, db.Meal{MealTime: CheckMealTime(), 
											MealName: slcMeal[0], 
											MealCcal: ccalInt, 
											MealDate: time.Now().Format("02-01-2006") })
	}
	

	return slcmeals, nil

}
