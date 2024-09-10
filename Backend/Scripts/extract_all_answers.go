/*
Author: Keyvan Sadeghi

Short: This file will extract all answers from the task files and write them to the sql file.
It uses the utils package to abstract away some of the minor implementation
*/

package main

import (
	utils "SQL_Learning_Framework/Scripts/Utils"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("Starting up...")
	root := "Static/Problems"

	// list of all queries from all files
	list := extract_from_files(root)
	if list == nil {
		fmt.Printf("List is nil!!!!\n")
	}

	// write all answers to sql file
	

}

/*
This function will path traverse all sub directories in Static/Problems and extract all the answers
a student has written.
*/
func extract_from_files(path string) []string {
	var list []string
	// num_tasks := preprocessing_tasks(root_dir)
	// fmt.Println(num_tasks)
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue // skip iteration
		}

		abs_dir_path := filepath.Join(path, dir.Name())
		files, _ := os.ReadDir(abs_dir_path)
		
		// Get the asnwer in files
		for _, file := range files {
			file_name := filepath.Join(abs_dir_path, file.Name())
			file_answer := utils.Get_answer(file_name)
			list = append(list, file_answer)
			fmt.Printf("File name: %s\n, File Content: %s\n", file.Name(), file_answer)
		}
	}

	return list
}

// 	return list
// }

func add_answer(file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	// Get answer from whole markdown content
	answer := utils.Get_answer(string(content))

	// Write files to my_queries.sql
	my_queries_file := "SQL_Setup/my_queries.sql"
	err = utils.Write_answer(answer, my_queries_file)

	if err != nil {
		return err
	}

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
