package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelMac = input.Mac{}

	dataMac = map[string]struct {
		inputData string
		result    bool
	}{
		"mac correcta ":         {"57:3c:f0:5e:44:44", true},
		"incorrecta ":           {"57:3c:f0:5e:44:44.", false},
		"incorrecto con guion ": {"57-3c-f0-5e-44-44", false},
		"sin data ":             {"", false},
	}
)

func Test_InputMac(t *testing.T) {
	for prueba, data := range dataMac {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelMac.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
