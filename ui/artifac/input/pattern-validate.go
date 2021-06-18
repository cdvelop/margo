package input

import "regexp"

//patternValidate ..
func patternValidate(pattern, valueIN string) bool {
	// log.Printf("pattern: %v value: %v", p, valueIN)
	pvalid := regexp.MustCompile(pattern)
	// log.Printf("value: [%v] match: [%v]", valueIN, pvalid.MatchString(valueIN))
	return pvalid.MatchString(valueIN)
}
