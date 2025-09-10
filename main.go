package main

import (
	"fmt"
	"os"
	"strings"
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

func findOccurrences(text, sub string) []int {
	var positions []int
	offset := 0
	for {
		idx := strings.Index(text[offset:], sub)
		if idx == -1 {
			break
		}
		positions = append(positions, offset+idx)
		offset += idx + len(sub)
	}
	return positions
}

func replaceText(text, old, new string, replaceAll bool) string {
	if replaceAll {
		return strings.ReplaceAll(text, old, new)
	}
	return strings.Replace(text, old, new, 1)
}

func main() {
	fmt.Println("CLI Text Editor - Go vanilla!")

	filePath := "./document.txt"
	fileData, err := readFilePath(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Arquivo:", fileData)

	// Teste da busca
	occ := findOccurrences(fileData, "Lucas")
	fmt.Println("Posições de 'Lucas':", occ)

	// Teste substituição
	newContent := replaceText(fileData, "Lucas", "João", true)
	fmt.Println("Novo conteúdo:", newContent)

	// Grava o novo texto
	if err := writeFile(filePath, newContent); err != nil {
		fmt.Println("Erro ao salvar:", err)
	}
}
