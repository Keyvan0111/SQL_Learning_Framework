package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting up...")
	file := "Static/Problems/1_0_Easy/1_1.md"

	content, err := os.ReadFile(file)
	if err != nil { panic(err) }
	
	answer := get_answer(string(content))

	fmt.Println(string(content))
	fmt.Printf("Answer:%s, Length: %d", answer, len(answer))

	my_queries_file := "SQL_Setup/my_queries.sql"
	err = write_answer(answer, my_queries_file)
	if err != nil { fmt.Printf("Error when writing to sql file: %s", err)}
	
}

func write_answer(query string, file_name string) error {
	data := []byte(query)
	err := os.WriteFile(file_name, data, 0777)

	if err != nil { return err }
	return nil
}
	
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