# Editor de Texto em Linha de Comando (Go vanilla)

Um editor de texto minimalista em linha de comando escrito em Go, focado em operações básicas de arquivo e manipulação simples de texto (buscar/substituir).

Recursos:
- Abrir um arquivo existente ou iniciar com um buffer vazio
- Visualizar o conteúdo atual
- Buscar ocorrências de uma substring
- Substituir a primeira ou todas as ocorrências
- Salvar e Salvar Como
- Prompts simples e amigáveis no terminal

Requisitos:
- Go 1.20+ (o módulo declara 1.24; qualquer Go recente deve funcionar)

Início rápido:
```bash
go run . [caminho-opcional-para-arquivo]
```

Exemplos de uso:
```bash
# Iniciar com buffer vazio
go run .

# Abrir um arquivo existente
go run . document.txt
```

Ações do menu principal:
- 1) Novo Documento (começar do zero)
- 2) Editar Documento (acrescentar ou substituir conteúdo)
- 3) Salvar (adiciona automaticamente extensão .txt)
- 4) Salvar Como (escolher novo nome, adiciona .txt automaticamente)
- 5) Buscar texto
- 6) Substituir texto (primeira ou todas)
- 7) Visualizar conteúdo
- 8) Sair

Recursos amigáveis ao usuário:
- Pressione Enter duas vezes para finalizar a edição
- Extensão .txt automática ao salvar
- Prompts claros e intuitivos
- Opção Novo Documento para começar do zero

Notas para desenvolvimento:
- Funções principais: `readFilePath`, `writeFile`, `findOccurrences`, `replaceText`
- Logger variádico: `logInfo(format string, args ...any)`
- Auxiliar de entrada: `readLine(reader, prompt)`

Licença: MIT

Documentação em inglês: veja `README.md`
