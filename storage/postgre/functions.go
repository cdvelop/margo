package postgre

import "fmt"

// SQLParameters string  //ej: postgres: "$" sqlite "?",
func (d *DNS) SQLsyntax() string {
	return `$1`
}

//SetListSyntax "%v=$%v"
func (d *DNS) SetListSyntax(key string, i byte, set *[]string) {
	*set = append(*set, fmt.Sprintf("%v=$%v", key, i))
}

func (d *DNS) TotalValuesSyntax(fields map[string]string) string {
	return fmt.Sprintf("$%v", len(fields))
}

func (d *DNS) MakeSqInsertSyntax(i *byte, setValue *[]string) {
	*setValue = append(*setValue, fmt.Sprintf("$%v", *i+1))
	*i++
}
