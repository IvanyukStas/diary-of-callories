package utilsfunc

import (
	"docc/db"
	"errors"
	"fmt"
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
	for strMeal
	 := slcstr{

		// slcmeals = append(slcmeals, db.Meal{MealTime: CheckMealTime(), })

	}
	

	fmt.Println(slcstr)
	return nil, nil

}
