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
	
func Get_answer(data string) string {
	answer_header := "### Answer here:"
	header_index := strings.Index(data, answer_header)

	// Get the index of header to procees extracting answer
	if header_index == -1 {
		fmt.Printf("File does not contain the given header: %s", answer_header)
		return ""
	}

	// Get the answer
	start_index := header_index + len(answer_header)
	content := data[start_index:]

	return content
}

/*
This function will setup the needed items for path traversing

Return:
	- array containing number of files per difficulty
*/
func Preprocessing_tasks(root_dir string) *[3]int {
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

/*
This function will path traverse all sub directories in Static/Problems and extract all the answers
a student has written.
*/
func Extract_from_files(root_dir string) {
	
	// num_tasks := preprocessing_tasks(root_dir)
	// fmt.Println(num_tasks)

}
