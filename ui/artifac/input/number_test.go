package input_test

import (
	"testing"

	"github.com/cdvelop/margo/ui/artifac/input"
)

var (
	modelNumber = input.Number{}

	dataNumber = map[string]struct {
		inputData string
		result    bool
	}{
		"numero correcto 100":              {"100", true},
		"un caracter 0":                    {"0", true},
		"un caracter 1":                    {"1", true},
		"grande 10000232326263727":         {"10000232326263727", true},
		"grande y letra 10000232e26263727": {"10000232E26263727", false},
		"numero negativo -100 ":            {"-100", false},
		"texto en vez de numero ":          {"h", false},
		"texto y numero":                   {"h1", false},
	}
)

func Test_InputNumber(t *testing.T) {
	for prueba, data := range dataNumber {
		t.Run((prueba), func(t *testing.T) {
			if ok := modelNumber.Validate(data.inputData); ok != data.result {
				t.Fatalf("resultado [%v]", ok)
			}
		})
	}
}
