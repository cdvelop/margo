package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelTextNum = input.TextNum{}

	dataTextNum = map[string]struct {
		inputData string
		result    bool
	}{
		"guion bajo ":           {"son_24_botellas", true},
		"frase con guion bajo ": {"los_cuatro", true},
		"frase sin guion bajo ": {"los cuatro", false},
		"palabras guion bajo ":  {"son_2_cuadros", true},
		"palabras separadas ":   {"son 2 cuadros", false},
		"palabras guion medio ": {"son-2-cuadros", false},
		"menos de 5 palabras ":  {"tres", false},
	}
)

func Test_InputTextNum(t *testing.T) {
	for prueba, data := range dataTextNum {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelTextNum.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
