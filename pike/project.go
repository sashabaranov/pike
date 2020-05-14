package pike

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type Project struct {
	Name     string   `yaml:"name"`
	Entities []Entity `yaml:"entities"`
}

func ProjectFromYAMLString(yamlStr string) (proj Project, err error) {
	err = yaml.Unmarshal([]byte(yamlStr), &proj)
	proj.Validate()
	return
}

func (p Project) ProtoCapsName() string {
	return GoCamelCase(p.Name)
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
	fmt.Println("Path prefix: ", pathPrefix)

	p.executeTemplate("initial_migration.up.sql.tmplt", fmt.Sprintf("%s.up.sql", pathPrefix))
	p.executeTemplate("initial_migration.down.sql.tmplt", fmt.Sprintf("%s.down.sql", pathPrefix))

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

func timeVersion() string {
	return time.Now().Format("20060102150405")
}
