package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelDate = input.Date{}

	dataDate = map[string]struct {
		inputData string
		result    bool
	}{
		"correcto ":                  {"2002-12-03", true},
		"caracter de mas icorrecto ": {"2002-12-03-", false},
		"formato incorrecto ":        {"21/12/1998", false},
		"fecha incorrecta ":          {"2020-31-01", false},
		"data incorrecta ":           {"0000-00-00", false},
	}
)

func Test_InputDate(t *testing.T) {
	for prueba, data := range dataDate {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelDate.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
