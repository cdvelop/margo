package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelCheck = input.Check{Values: map[string]string{"1": "Admin", "2": "editor", "3": "visitante"}}

	dataCheck = map[string]struct {
		inputData string
		result    bool
	}{
		"una credencial":            {"1", true},
		"editor y admin":            {"1,2", true},
		"todas las credenciales ok": {"1,2,3", true},
		"0 no existe":               {"0", false},
		"-1 no valido":              {"-1", false},
		"una buena una inexistente": {"1,5", false},
		"sin data":                  {"", false},
		"espacios no":               {"luis , 3", false},
	}
)

func Test_Check(t *testing.T) {
	for prueba, data := range dataCheck {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			if ok := modelCheck.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
