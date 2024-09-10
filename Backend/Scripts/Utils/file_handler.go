/*
Author: Keyvan Sadeghi

Short: This file contains handlers for file manipulation
*/

package utils

import (
	"fmt"
	"os"
	"strings"

)

func Write_answer(query string, file_name string) error {
	data := []byte(query)

	err := os.WriteFile(file_name, data, 0777)

	if err != nil { return err }
	return nil
}

func Get_answer(file_path string) string {
	data_bytes, err := os.ReadFile(file_path)
	data_str := string(data_bytes)
	if err != nil {
		panic(err)
	}

	answer_header := "### Answer here:"
	header_index := strings.Index(data_str, answer_header)

	// Get the index of header to procees extracting answer
	if header_index == -1 {
		fmt.Printf("File does not contain the given header: %s", answer_header)
		return ""
	}

	// Get the answer
	start_index := header_index + len(answer_header)
	content := data_str[start_index:]

	return content
}

/*
This function will setup the needed items for path traversing

Return:
	- array containing number of files per difficulty
*/
func preprocessing_tasks(root_dir string) *[3]int {
	// array of zeros with len 3
	var num_questions [3]int
	paths := [3]string{"/1_0_Easy", "/2_0_Medium", "/3_0_Hard"}

	for i := 0; i < len(paths); i++ {
		// Get list of files in folder
		inner_path := root_dir + paths[i]
		files, _ := os.ReadDir(inner_path)

		for j := 0; j < len(files); j++ {
			num_questions[j] = len(files)
		}
	}
	
	return &num_questions
}

func get_num_files(folder_abs_path string) int {
	files, _ := os.ReadDir(folder_abs_path)
	length := len(files)
	return length
}


