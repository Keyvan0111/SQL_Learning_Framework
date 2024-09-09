/*
Author: Keyvan Sadeghi

Short: This file will extract all answers from the task files and write them to the sql file.
It uses the utils package to abstract away some of the minor implementation
*/

package main

import (
	"fmt"
	"os"
	"SQL_Learning_Framework/Scripts/Utils"
)

func main() {
	fmt.Println("Starting up...")

	add_all_answers()
	
}


func add_answer(file string) error {
	content, err := os.ReadFile(file)
	if err != nil { return err }
	
	// Get answer from whole markdown content
	answer := utils.Get_answer(string(content))
	
	fmt.Println(string(content))
	fmt.Printf("Answer:%s, Length: %d", answer, len(answer))
	
	// Write files to my_queries.sql
	my_queries_file := "SQL_Setup/my_queries.sql"
	err = utils.Write_answer(answer, my_queries_file)
	
	if err != nil { return err}
	
	return err
}

func add_all_answers() {
	// Use recursion to traverse all files
	file := "Static/Problems/1_0_Easy/1_1.md"
	
	err := add_answer(file)
	if err != nil { 
		fmt.Printf("Error occured when adding answers: %s", err)
		panic(err)
	}

}
