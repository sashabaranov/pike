package pike

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type Project struct {
	Name     string   `yaml:"name"`
	Entities []Entity `yaml:"entities"`

	GoImportPath string `yaml:"go_import_path"`

	OverrideConfigEnvVar string `yaml:"config_env_var"`
}

func ProjectFromYAMLString(yamlStr string) (proj Project, err error) {
	err = yaml.Unmarshal([]byte(yamlStr), &proj)
	proj.Validate()
	return
}

func (p Project) ProtoCapsName() string {
	return GoCamelCase(p.Name)
}

func (p Project) ConfigEnvVariable() string {
	if p.OverrideConfigEnvVar != "" {
		return p.OverrideConfigEnvVar
	}
	return fmt.Sprintf("%s_CONFIG", strings.ToUpper(p.Name))
}

func (p Project) Validate() {
	for _, entity := range p.Entities {
		err := entity.Validate()
		if err != nil {
			log.Fatalf("Error validating %s: %v", entity.Name, err)
		}
	}
}

func (p Project) GenerateProto(path string) {
	t, err := template.New("project_proto.tmplt").Funcs(templateFuncMap).ParseFiles("templates/project_proto.tmplt")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	outputFile, err := os.Create(path)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	err = t.Execute(outputFile, p)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func (p Project) GenerateSQLMigrations(path string) {
	version := timeVersion()
	pathPrefix := filepath.Join(
		path,
		fmt.Sprintf("%s_initial", version),
	)

	p.executeTemplate("initial_migration.up.sql.tmplt", fmt.Sprintf("%s.up.sql", pathPrefix))
	p.executeTemplate("initial_migration.down.sql.tmplt", fmt.Sprintf("%s.down.sql", pathPrefix))
}

func (p Project) GenerateGoFiles(path string) {
	files := []string{
		"storage.go",
		"server.go",
		"report_error.go",
		"config.go",
		"run.go",
	}

	for _, filename := range files {
		fmt.Printf("⚙️  Generating %s\n", filename)
		p.executeTemplate(
			fmt.Sprintf("%s.tmplt", filename),
			filepath.Join(path, filename),
		)
	}
}

func (p Project) GenerateConfigFiles(path string) {
	fmt.Println("⚙️  Generating config file")
	p.executeTemplate(
		"config.yaml.tmplt",
		filepath.Join(path, "dev.yaml"),
	)
}

func (p Project) GenerateLauncher(path string) {
	fmt.Println("⚙️  Generating launch file")
	p.executeTemplate(
		"launcher.go.tmplt",
		filepath.Join(path, "main.go"),
	)
}

func (p Project) executeTemplate(templateName, outputPath string) {
	templatePath := fmt.Sprintf("templates/%s", templateName)
	t, err := template.New(templateName).Funcs(templateFuncMap).ParseFiles(templatePath)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	err = t.Execute(outputFile, p)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func (p Project) CreateDirectories() {
	dirs := []string{
		"proto",
		p.Name,
		"sql/migrations",
		"configs",
		"cli",
		"bin",
	}

	for _, dir := range dirs {
		path := filepath.Join(p.AbsolutePath(), dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("Error creating directory %s: %v", dir, err)
		}
		fmt.Printf("📂 Created directory %s\n", path)
	}
}

func (p Project) AbsolutePath() string {
	goPath := os.Getenv("GOPATH")
	return filepath.Join(goPath, "src", p.GoImportPath)
}

func timeVersion() string {
	return time.Now().Format("20060102150405")
}
