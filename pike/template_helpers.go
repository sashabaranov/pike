package pike

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"
)

var templateFuncMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
	"last": func(x int, a interface{}) bool {
		return x == reflect.ValueOf(a).Len()-1
	},
}

func executeTemplate(templateName, outputPath string, data interface{}) {
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

	err = t.Execute(outputFile, data)
	if err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}
