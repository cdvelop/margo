package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelObject = input.Radio{Values: []string{"Dama", "Varon"}}

	datatest = map[string]struct {
		inputData string
		result    bool
	}{
		"D Dato correcto":                {"D", true},
		"V Dato correcto":                {"V", true},
		"d Dato en minuscula incorrecto": {"d", false},
		"v Dato en minuscula incorrecto": {"v", false},
		"Dato no valido":                 {"1", false},
		"sin data":                       {"", false},
	}
)

func Test_RadioButton(t *testing.T) {
	for prueba, data := range datatest {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelObject.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
