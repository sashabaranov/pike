package pike

import (
	"strings"
)

type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`

	SQLTypeOverride string `yaml:"sql_type"`

	IsPrimaryKey bool `yaml:"primary_key"`
}

func (f Field) GoName() string {
	return GoCamelCase(f.Name)
}

func (f Field) SQLType() string {
	if f.SQLTypeOverride != "" {
		return strings.ToUpper(f.SQLTypeOverride)
	}

	if strings.HasPrefix(f.Type, "int") && f.IsPrimaryKey {
		if f.Type == "int64" {
			return "BIGSERIAL PRIMARY KEY"
		}
		return "SERIAL PRIMARY KEY"
	}

	sqlType := map[string]string{
		"int32":  "INTEGER",
		"int64":  "BIGINT",
		"string": "TEXT",
		"float":  "REAL",
	}[f.Type]

	if sqlType == "" {
		sqlType = "TEXT"
	}
	if f.IsPrimaryKey {
		sqlType += " PRIMARY KEY"
	}
	return sqlType
}

// Following code imported from package google.golang.org/protobuf/internal/strs
// Due to package being internal it is impossible to import it

// GoCamelCase camel-cases a protobuf name for use as a Go identifier.
//
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
func GoCamelCase(s string) string {
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c == '.' && i+1 < len(s) && isASCIILower(s[i+1]):
			// Skip over '.' in ".{{lowercase}}".
		case c == '.':
			b = append(b, '_') // convert '.' to '_'
		case c == '_' && (i == 0 || s[i-1] == '.'):
			// Convert initial '_' to ensure we start with a capital letter.
			// Do the same for '_' after '.' to match historic behavior.
			b = append(b, 'X') // convert '_' to 'X'
		case c == '_' && i+1 < len(s) && isASCIILower(s[i+1]):
			// Skip over '_' in "_{{lowercase}}".
		case isASCIIDigit(c):
			b = append(b, c)
		default:
			// Assume we have a letter now - if not, it's a bogus identifier.
			// The next word is a sequence of characters that must start upper case.
			if isASCIILower(c) {
				c -= 'a' - 'A' // convert lowercase to uppercase
			}
			b = append(b, c)

			// Accept lower case sequence that follows.
			for ; i+1 < len(s) && isASCIILower(s[i+1]); i++ {
				b = append(b, s[i+1])
			}
		}
	}
	return string(b)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}
func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}
func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}
