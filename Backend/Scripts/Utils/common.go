package utils

import (
	"fmt"
	"os"
)

func Usage() {
	fmt.Println("Error parsing CLI arguments. Please check usage for running the program.")
	fmt.Printf("Usage: ./exec <Problems_path> <Sql_path>")
	os.Exit(1)
}
