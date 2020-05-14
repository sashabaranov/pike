package pike

import (
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
