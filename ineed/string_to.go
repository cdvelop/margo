package ineed

import (
	"fmt"
	"log"
	"strconv"
)

//StringtoINT64 max "9223372036854775807"
func StringtoINT64(s string, bigint *int64) bool {
	var err error
	*bigint, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Printf("error en coversion de %v a int64 %v", s, err)
		return false
	}
	return true
}

// StringtoBYTE convierte un string a un
// uint8 8-bit integers (0 to 255)
func StringtoBYTE(s string) (byte, bool) {
	if len(s) > 0 && len(s) <= 3 {
		fmt.Printf("ok %v", len(s))

	}
	var i64 int64
	if !StringtoINT64(s, &i64) {
		return 0, false
	}
	return uint8(i64), true
}
