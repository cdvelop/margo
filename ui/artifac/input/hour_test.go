package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelHour = input.Hour{}

	dataHour = map[string]struct {
		inputData string
		result    bool
	}{
		"correcto":    {"23:59", true},
		"correcto 00": {"00:00", true},
		"correcto 12": {"12:00", true},

		"icorrecto 24":       {"24:00", false},
		"icorrecto sin data": {"", false},
		"icorrecto caracter": {"13-34", false},
	}
)

func Test_InputHour(t *testing.T) {
	for prueba, data := range dataHour {
		t.Run((prueba + " " + data.inputData), func(t *testing.T) {
			if ok := modelHour.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
