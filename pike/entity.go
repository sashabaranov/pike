package pike

import (
	"errors"
	"fmt"
)

type Entity struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`

	SQLTableOverride string `yaml:"sql_table_name"`
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

func (e Entity) SQLTableName() string {
	if e.SQLTableOverride != "" {
		return e.SQLTableOverride
	}

	return fmt.Sprintf("%ss", e.Name)
}
