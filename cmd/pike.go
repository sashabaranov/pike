package main

import (
	"fmt"
	"github.com/sashabaranov/pike/pike"
	"log"
	"os"
	"path/filepath"
)

const testYAML = `
name: backend
go_import_path: github.com/sashabaranov/testbackend
entities:
  - name: animal
    fields:
      - {name: id, type: uint32, primary_key: true}
      - {name: name, type: string}
      - {name: password_hash, type: string}
      - {name: age, type: int32}
      - {name: userpic_url, type: string}
`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pike <project path inside Go path>")
		return
	}

	goPath := os.Getenv("GOPATH")
	projectDir := filepath.Join(goPath, "src", os.Args[1])

	proj, err := pike.ProjectFromYAMLString(testYAML)
	if err != nil {
		log.Fatalf("Error unmarshalling yaml: %v", err)
	}

	proj.GenerateProto(filepath.Join(projectDir, "proto/project.proto"))
	proj.GenerateSQLMigrations(filepath.Join(projectDir, "sql/migrations"))
	proj.GenerateGoFiles(filepath.Join(projectDir, proj.Name))
	proj.GenerateConfigFiles(filepath.Join(projectDir, "configs"))
	proj.GenerateLauncher(filepath.Join(projectDir, "cli"))
}
