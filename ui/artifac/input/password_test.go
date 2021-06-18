package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelPassword = input.Password{}

	dataPassword = map[string]struct {
		inputData string
		result    bool
	}{
		"valida numero letras y caracter": {"c0ntra$3", true},
		"valida muchos caracteres":        {"M1 contraseÑ4", true},
		"valida 8 caracteres":             {"contrase", true},
		"invalida menos de 5 caracteres":  {"cont", false},
		"invalida solo numeros":           {"12345", false},
		"no valida menos de 5":            {"12e4", false},
		"sin data":                        {"", false},
	}
)

func Test_InputPassword(t *testing.T) {
	for prueba, data := range dataPassword {
		t.Run((prueba + ": " + data.inputData), func(t *testing.T) {
			if ok := modelPassword.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
