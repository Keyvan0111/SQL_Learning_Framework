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