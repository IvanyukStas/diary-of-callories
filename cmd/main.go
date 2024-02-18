package main

import (
	"docc/logger"
	"log"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

}

func run() error {
	log.Println("Start program!")
	logg := logger.Loging1{}
	logg.New()
	logg.I.Println("fdfsdfg")

	// logdir := filepath.Join("../", "log", "1.log")
	// f, err := os.OpenFile(logdir,
	// 	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer f.Close()

	// logger := log.New(f, "префикс: ", log.LstdFlags)
	// logger.Println("текст для добавления")
	// logger.Println("еще текст для добавления")
	time.Sleep(5 * time.Second)
	defer log.Println("program is closed!")
	return nil
}
