package main

import (
	"fmt"
	"github.com/sashabaranov/pike/pike"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: pike project.yaml")
		return
	}

	projFile := os.Args[1]
	content, err := ioutil.ReadFile(projFile)
	if err != nil {
		fmt.Printf("Error reading project file: %v\n", err)
		return
	}

	proj, err := pike.ProjectFromYAMLString(string(content))
	if err != nil {
		fmt.Printf("Error unmarshalling yaml: %v\n", err)
		return
	}

	proj.CheckDirectoryNotPresent()
	proj.CreateDirectories()

	projectDir := proj.AbsolutePath()

	protobufFile := fmt.Sprintf("proto/%s.proto", proj.Name)
	proj.GenerateProto(filepath.Join(projectDir, protobufFile))

	proj.GenerateSQLMigrations(filepath.Join(projectDir, "sql/migrations"))
	proj.GenerateGoFiles(filepath.Join(projectDir, proj.Name))
	proj.GenerateConfigFiles(filepath.Join(projectDir, "configs"))
	proj.GenerateLauncher(filepath.Join(projectDir, "cli"))
	proj.GenerateBinScripts(filepath.Join(projectDir, "bin"))

	proj.PrintOutro()
}
