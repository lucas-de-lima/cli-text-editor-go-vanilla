package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func writeFile(fp string, ct string) error {
	err := os.WriteFile(fp, []byte(ct), 0666)
	return err
}

func readFilePath(fp string) (string, error) {
	data, err := os.ReadFile(fp)
	if os.IsNotExist(err) {
		fmt.Println("File not found")
		return "", err
	}
	if err != nil {
		fmt.Println("Failed to read file")
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

// variadic logger with optional tags
func logInfo(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

func readLine(reader *bufio.Reader, prompt string) (string, error) {
	if prompt != "" {
		fmt.Print(prompt)
	}
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(line, "\r\n"), nil
}

func showMenu() {
	fmt.Println()
	fmt.Println("=== CLI Text Editor ===")
	fmt.Println("1) New Document")
	fmt.Println("2) Edit Document")
	fmt.Println("3) Save")
	fmt.Println("4) Save As")
	fmt.Println("5) Find text")
	fmt.Println("6) Replace text")
	fmt.Println("7) View content")
	fmt.Println("8) Exit")
}

func main() {
	fmt.Println("CLI Text Editor - Go vanilla!")

	reader := bufio.NewReader(os.Stdin)

	var filePath string
	var content string

	// CLI arg: file path
	if len(os.Args) > 1 {
		filePath = os.Args[1]
		data, err := readFilePath(filePath)
		if err != nil {
			if os.IsNotExist(err) {
				logInfo("Starting with a new file: %s", filePath)
				content = ""
			} else {
				logInfo("Error opening file: %v", err)
				return
			}
		} else {
			content = data
			logInfo("Opened file: %s (%d bytes)", filePath, len(content))
		}
	} else {
		logInfo("No file provided. Starting with an empty buffer.")
		filePath = ""
		content = ""
	}

	for {
		showMenu()
		choice, err := readLine(reader, "Select an option [1-8]: ")
		if err != nil {
			logInfo("Input error: %v", err)
			continue
		}

		switch strings.TrimSpace(choice) {
		case "1": // New Document
			fmt.Println()
			fmt.Println("Starting a new document...")
			content = ""
			filePath = ""
			logInfo("New document created. Current content cleared.")

		case "2": // Edit Document
			fmt.Println()
			fmt.Println("Edit modes:")
			fmt.Println("a) Append to current content")
			fmt.Println("b) Replace entire content")
			mode, err := readLine(reader, "Choose [a/b]: ")
			if err != nil {
				logInfo("Input error: %v", err)
				continue
			}

			switch strings.ToLower(strings.TrimSpace(mode)) {
			case "a":
				fmt.Println("Enter text to append (press Enter twice to finish):")
				added, err := readMultilineHuman(reader)
				if err != nil {
					logInfo("Input error: %v", err)
					continue
				}
				if len(added) > 0 {
					if len(content) > 0 && !strings.HasSuffix(content, "\n") {
						content += "\n"
					}
					content += added
					logInfo("Text appended successfully.")
				}
			case "b":
				fmt.Println("Enter new content (press Enter twice to finish):")
				newBody, err := readMultilineHuman(reader)
				if err != nil {
					logInfo("Input error: %v", err)
					continue
				}
				content = newBody
				logInfo("Content replaced successfully.")
			default:
				logInfo("Unknown edit mode: %s", mode)
			}

		case "3": // Save
			if filePath == "" {
				path, err := readLine(reader, "Enter filename (without extension): ")
				if err != nil {
					logInfo("Input error: %v", err)
					continue
				}
				path = strings.TrimSpace(path)
				if path == "" {
					logInfo("Invalid filename.")
					continue
				}
				filePath = addTxtExtension(path)
			}
			if err := writeFile(filePath, content); err != nil {
				logInfo("Failed to save: %v", err)
			} else {
				logInfo("Document saved to %s (%d bytes)", filePath, len(content))
			}

		case "4": // Save As
			path, err := readLine(reader, "Enter filename (without extension): ")
			if err != nil {
				logInfo("Input error: %v", err)
				continue
			}
			path = strings.TrimSpace(path)
			if path == "" {
				logInfo("Invalid filename.")
				continue
			}
			path = addTxtExtension(path)
			if err := writeFile(path, content); err != nil {
				logInfo("Failed to save: %v", err)
			} else {
				filePath = path
				logInfo("Document saved to %s (%d bytes)", filePath, len(content))
			}

		case "5": // Find text
			query, err := readLine(reader, "Search for: ")
			if err != nil {
				logInfo("Input error: %v", err)
				continue
			}
			if query == "" {
				logInfo("Nothing to search.")
				continue
			}
			positions := findOccurrences(content, query)
			if len(positions) == 0 {
				logInfo("Text not found.")
			} else {
				logInfo("Found %d occurrence(s) at positions: %v", len(positions), positions)
			}

		case "6": // Replace text
			oldText, err := readLine(reader, "Find: ")
			if err != nil {
				logInfo("Input error: %v", err)
				continue
			}
			if oldText == "" {
				logInfo("Nothing to replace.")
				continue
			}
			newText, err := readLine(reader, "Replace with: ")
			if err != nil {
				logInfo("Input error: %v", err)
				continue
			}
			mode, err := readLine(reader, "Replace all occurrences? (y/N): ")
			if err != nil {
				logInfo("Input error: %v", err)
				continue
			}
			replaceAll := strings.EqualFold(strings.TrimSpace(mode), "y")
			updated := replaceText(content, oldText, newText, replaceAll)
			if updated == content {
				logInfo("No changes made.")
			} else {
				content = updated
				if replaceAll {
					logInfo("All occurrences replaced.")
				} else {
					logInfo("First occurrence replaced.")
				}
			}

		case "7": // View content
			fmt.Println()
			if filePath == "" {
				fmt.Println("[Unsaved document]")
			} else {
				fmt.Println("[File:", filePath+"]")
			}
			fmt.Println("----- BEGIN -----")
			if content == "" {
				fmt.Println("(empty)")
			} else {
				fmt.Print(content)
			}
			fmt.Println("------ END ------")

		case "8": // Exit
			logInfo("Goodbye!")
			return

		default:
			logInfo("Unknown option: %s", choice)
		}
	}
}

// read multiple lines until user presses Enter twice (human-friendly)
func readMultilineHuman(reader *bufio.Reader) (string, error) {
	var b strings.Builder
	emptyLineCount := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		trimmed := strings.TrimRight(line, "\r\n")

		if trimmed == "" {
			emptyLineCount++
			if emptyLineCount >= 2 {
				break
			}
		} else {
			emptyLineCount = 0
		}

		b.WriteString(trimmed)
		b.WriteString("\n")
	}

	result := b.String()
	// Remove trailing empty lines
	result = strings.TrimRight(result, "\n")
	if result != "" {
		result += "\n"
	}
	return result, nil
}

// automatically add .txt extension if not present
func addTxtExtension(filename string) string {
	if strings.HasSuffix(strings.ToLower(filename), ".txt") {
		return filename
	}
	return filename + ".txt"
}
