package sqlite

import "fmt"

// SQLParameters string  //ej: postgres: "$" sqlite "?",
func (d *DNS) SQLsyntax() string {
	return `?`
}

//SetListSyntax "%v=?", k
func (d *DNS) SetListSyntax(key string, i byte, set *[]string) {
	*set = append(*set, fmt.Sprintf("%v=?", key))
}

func (d *DNS) TotalValuesSyntax(fields map[string]string) string {
	return "?"
}

func (d *DNS) MakeSqInsertSyntax(i *byte, setValue *[]string) {
	*setValue = append(*setValue, "?")
}
