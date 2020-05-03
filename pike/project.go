package pike

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"text/template"
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

func (p Project) GenerateProto() {
	t, err := template.New("project_proto.tmplt").Funcs(templateFuncMap).ParseFiles("templates/project_proto.tmplt")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
