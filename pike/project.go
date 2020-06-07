package pike

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	executeTemplate("project_proto.tmplt", path, p)
}

func (p Project) GenerateSQLMigrations(path string) {
	version := timeVersion()
	pathPrefix := filepath.Join(
		path,
		fmt.Sprintf("%s_initial", version),
	)

	executeTemplate("initial_migration.up.sql.tmplt", fmt.Sprintf("%s.up.sql", pathPrefix), p)
	executeTemplate("initial_migration.down.sql.tmplt", fmt.Sprintf("%s.down.sql", pathPrefix), p)
	fmt.Printf("🌿  Generated SQL migrations into %s directory\n", Bold("sql/migrations"))
}

func (p Project) GenerateGoFiles(path string) {
	files := []string{
		"storage.go",
		"server.go",
		"report_error.go",
		"run.go",
		"auth.go",
	}

	for _, filename := range files {
		executeTemplate(
			fmt.Sprintf("%s.tmplt", filename),
			filepath.Join(path, filename),
			p,
		)
	}

	for _, entity := range p.Entities {
		perEntityFiles := []string{
			fmt.Sprintf("server_%s.go", entity.Name),
			fmt.Sprintf("storage_%s.go", entity.Name),
		}
		tmpProject := Project{
			Name:     p.Name,
			Entities: []Entity{entity},
		}

		for _, filename := range perEntityFiles {
			templateName := strings.ReplaceAll(filename, entity.Name, "entity")
			executeTemplate(
				fmt.Sprintf("%s.tmplt", templateName),
				filepath.Join(path, filename),
				tmpProject,
			)
		}
	}

	fmt.Printf("🌿  Generated Go implementation into %s directory\n", Bold(p.Name))
}

func (p Project) GenerateConfigFiles(path string) {
	executeTemplate(
		"config.yaml.tmplt",
		filepath.Join(path, "dev.yaml"),
		p,
	)

	fmt.Printf("🌿  Generated %s config file\n", Bold("configs/dev.yaml"))
}

func (p Project) GenerateLauncher(path string) {
	executeTemplate(
		"launcher.go.tmplt",
		filepath.Join(path, "main.go"),
		p,
	)
}

func (p Project) GenerateBinScripts(path string) {
	executeTemplate(
		"run.sh.tmplt",
		filepath.Join(path, "run.sh"),
		p,
	)

	executeTemplate(
		"compile_proto.sh.tmplt",
		filepath.Join(path, "compile_proto.sh"),
		p,
	)

	os.Chmod(filepath.Join(path, "run.sh"), 0755)
	os.Chmod(filepath.Join(path, "compile_proto.sh"), 0755)

	fmt.Printf("🌿  Generated %s scripts\n", Bold("bin/"))
}

func (p Project) CheckDirectoryNotPresent() {
	_, err := os.Stat(p.AbsolutePath())
	if !os.IsNotExist(err) {
		fmt.Printf("Project directory must not exist. Try rm -r %s\n", p.AbsolutePath())
		os.Exit(0)
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

	var dirStr string
	for ix, dir := range dirs {
		path := filepath.Join(p.AbsolutePath(), dir)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("Error creating directory %s: %v", dir, err)
		}
		dirStr += dir
		if ix != len(dirs)-1 {
			dirStr += ", "
		}
	}
}

func (p Project) CompileProtobuf() {
	err := p.compileProto()
	if err != nil {
		fmt.Printf("❌  Failed to compile protobuf. Try %s yourself\n", Bold(Magenta("bin/compile_proto.sh")))
	}

	fmt.Printf("🌸  Compiled protobuf. Use %s in the future!\n", Bold(Magenta("bin/compile_proto.sh")))
}

func (p Project) AbsolutePath() string {
	goPath := os.Getenv("GOPATH")
	return filepath.Join(goPath, "src", p.GoImportPath)
}

func timeVersion() string {
	return time.Now().Format("20060102150405")
}
