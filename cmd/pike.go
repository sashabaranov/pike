package main

import (
	// "fmt"
	"github.com/sashabaranov/pike/pike"
	"log"
)

const testYAML = `
name: backend
entities:
  - name: animal
    fields:
      - {name: id, type: int64, primary_key: true}
      - {name: name, type: string}
      - {name: password_hash, type: string}
      - {name: age, type: int32}
      - {name: userpic_url, type: string}
`

func main() {
	proj, err := pike.ProjectFromYAMLString(testYAML)
	if err != nil {
		log.Fatalf("Error unmarshalling yaml: %v", err)
	}

	proj.GenerateProto()
}
