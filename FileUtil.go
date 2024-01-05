package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeTextFile() {
	file, err := os.Create("Product.txt")
	if err != nil {
		fmt.Println("File could not be created.")
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	_, writeErr := writer.WriteString(getProduct().toString())
	if writeErr != nil {
		fmt.Println("There is an error when writing file.")
	}

	fmt.Println("File has been created.")
	wg.Done()
}
