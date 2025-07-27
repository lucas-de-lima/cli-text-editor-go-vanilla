package main

import (
	"fmt"
	"os"
)




func readFilePath(fp string) (string, error) {
	data, err := os.ReadFile(fp)
	if os.IsNotExist(err) {
		fmt.Println("Path not found")
		return "", err
	}
	if err != nil {
		fmt.Println("Fail to read file")
		return "", err
	}

	return string(data), nil
}

func main() {
	fmt.Println("CLI Text Editor - Go vanilla!")

	fileData, err := readFilePath("./document.txt")
	if err != nil {
		fmt.Println("Fail to read file")
	}
	if fileData == "" {
		fmt.Println("Empty File!")
	}

	fmt.Println("Archive content: ", fileData)
}
