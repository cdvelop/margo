package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelRut = input.Rut{}

	dataRut = map[string]struct {
		inputData string
		result    bool
	}{
		"ok 7863697-1":                {"7863697-1", true},
		"cambio digito a k 7863697-k": {"7863697-k", false},
		"cambio digito a 0 7863697-0": {"7863697-0", false},
		"ok 14080717-6":               {"14080717-6", true},
		"incorrecto 14080717-0":       {"14080717-0", false},
	}
)

func Test_InputRut(t *testing.T) {
	for prueba, data := range dataRut {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelRut.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
