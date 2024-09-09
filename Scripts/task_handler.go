package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("Starting up...")

	file := "Static/Problems/1_0_Easy/1_1.md"
	content, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

func write_answer() {
	
}

func get_answer() {

}