package utilsfunc

import (
	
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
