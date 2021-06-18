package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelIp = input.Ip{}

	dataIp = map[string]struct {
		inputData string
		result    bool
	}{
		"ip correcta ":   {"192.168.1.1", true},
		"ip incorrecta ": {"192.168.1.1.8", false},
		"correcto ":      {"0.0.0.0", true},
		"sin data ":      {"", false},
	}
)

func Test_InputIp(t *testing.T) {
	for prueba, data := range dataIp {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelIp.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
