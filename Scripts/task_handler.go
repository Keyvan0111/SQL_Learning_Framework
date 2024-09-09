package main

import (
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Println("Starting up...")

	file := "Static/Problems/1_0_Easy/1_1.md"
	content, err := os.ReadFile(file)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
	answer := get_answer(string(content))
	fmt.Printf("Answer:%s, Length: %d", answer, len(answer))
	
}

// func write_answer(data string) {
	// 	answer_header := "### Answer here:"
	
	// }
	
func get_answer(data string) string {
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