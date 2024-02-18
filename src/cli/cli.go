package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFromSttin() string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		return scanner.Text()

	}
	return ""
}


func main() {
	for  {

	 r := ReadFromSttin()
	 if r == "q"{
		break
	 }
	 fmt.Println(r)
	}
}