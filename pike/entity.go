package pike

import "errors"

type Entity struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
}

func (e Entity) ProtoCapsName() string {
	return GoCamelCase(e.Name)
}

func (e Entity) Validate() error {
	nPrimaryKeys := 0
	for _, field := range e.Fields {
		if field.IsPrimaryKey {
			nPrimaryKeys += 1
		}
	}

	if nPrimaryKeys > 1 {
		return errors.New("Multiple primary keys")
	}
	return nil
}
