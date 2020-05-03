package pike

import (
	"text/template"
)

var templateFuncMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
}
