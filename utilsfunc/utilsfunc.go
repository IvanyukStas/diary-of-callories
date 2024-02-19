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
		return nil, errors.New("пустая строка! паполните строку")
	}
	slcstr := strings.Split(r, ",")
	for _, strMeal := range slcstr{
		slcMeal := convStrToSlcMeal(strMeal)
		if len(slcMeal) > 1{
			ccalInt, err := strconv.Atoi(slcMeal[1])
			if err != nil {
				fmt.Println("Несмог записать ", slcMeal[0])
				fmt.Println("После блюда "+ slcMeal[0] + " укажите граммовку цифрами")
			}
			slcmeals = append(slcmeals, db.Meal{MealTime: CheckMealTime(), 
				MealName: slcMeal[0], 
				MealCcal: ccalInt, 
				MealDate: time.Now().Format("02-01-2006") })
		
		}else {
			fmt.Println("Несмог записать ", strMeal)
				fmt.Println("После блюда "+ strMeal + " укажите граммовку цифрами")
				fmt.Println("Если указываете больше 1 блюда, разделяите их запятой")
		}
		
	}
	

	return slcmeals, nil

}

// Convert raw string Meal to slc full name of meal  and ccal
func convStrToSlcMeal(s string) []string{
	var d []string
	var mealName string
	for i, v := range s{
		_, err := strconv.Atoi(string(v))
		if err != nil {
			mealName = mealName + string(v)
		}else{
			d = append(d, mealName)
			ccalInt := s[i:]
			d = append(d, ccalInt)
			break
		}
	}
	return d
}
