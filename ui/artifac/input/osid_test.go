package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelOsid = input.Osid{}

	dataOsid = map[string]struct {
		inputData string
		result    bool
	}{
		"correcto ": {"S-1-5-21-620454579-1741226705-1933409994-1001", true},
		"invalido ": {"S-1-5-21-6204O4579-1741226705-1933409994-100l", false},
		"sin data ": {"", false},
	}
)

func Test_InputOsid(t *testing.T) {
	for prueba, data := range dataOsid {
		t.Run((prueba + data.inputData), func(t *testing.T) {
			if ok := modelOsid.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
