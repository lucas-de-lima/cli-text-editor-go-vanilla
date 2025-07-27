package main

import (
	"fmt"
	"os"
)


func writeFile(fp string, ct string) (error) {
	 err := os.WriteFile(fp, []byte(ct), 0666)
	 return err
}

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
	filePath := "./document.txt"

	// fileData, err := readFilePath(filePath)
	// if err != nil {
	// 	fmt.Println("Fail to read file")
	// }
	// if fileData == "" {
	// 	fmt.Println("Empty File!")
	// }

	// fmt.Println("Archive content: ", fileData)
	fmt.Println("________________________________")

	newContentToWrite := "Hello, my name is Lucas!!!"
	err := writeFile(filePath, newContentToWrite)
	if err != nil {
		fmt.Println("Fail to write new content")
	}
	fmt.Println("Success to write new content!")
}
