package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelText = input.Text{}

	dataText = map[string]struct {
		inputData string
		result    bool
	}{
		"no tilde ":                        {"peréz del rozal", false},
		"texto con ñ ":                     {"Ñuñez perez", true},
		"texto correcto + 3 caracteres ":   {"hola", true},
		"texto correcto 3 caracteres ":     {"los", true},
		"oracion ok ":                      {"hola que tal", true},
		"Dato numerico 100 no permitido ":  {"100", false},
		"menos de 3 caracteres ":           {"lo", false},
		"sin data ":                        {"", false},
		"un caracter numerico ":            {"8", false},
		"palabra mas numero no permitido ": {"son 4 bidones", false},
	}
)

func Test_InputText(t *testing.T) {
	for prueba, data := range dataText {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelText.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
