# CLI Text Editor (Go vanilla)

A minimal command-line text editor written in Go, focused on basic file operations and simple text manipulation (find/replace).

Features:
- Open an existing file or start with an empty buffer
- View current content
- Find occurrences of a substring
- Replace first or all occurrences
- Save and Save As
- Simple, friendly CLI prompts

Requirements:
- Go 1.20+ (module declares 1.24, any recent Go should work)

Quick start:
```bash
go run . [optional-path-to-file]
```

CLI usage examples:
```bash
# Start with an empty buffer
go run .

# Open an existing file
go run . document.txt
```

Main menu actions:
- 1) New Document (start fresh)
- 2) Edit Document (append or replace content)
- 3) Save (auto-adds .txt extension)
- 4) Save As (choose new filename, auto-adds .txt)
- 5) Find text
- 6) Replace text (first or all)
- 7) View content
- 8) Exit

Human-friendly features:
- Press Enter twice to finish editing
- Automatic .txt extension when saving
- Clear, intuitive prompts
- New Document option for fresh start

Developer notes:
- Core functions: `readFilePath`, `writeFile`, `findOccurrences`, `replaceText`
- Variadic logger: `logInfo(format string, args ...any)`
- Input helper: `readLine(reader, prompt)`

License: MIT

Documentation in Brazilian Portuguese: see [`README.pt-BR.md`](./README.pt-BR.md)
