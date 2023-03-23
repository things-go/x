package binding

import (
	"regexp"
)

const (
	numberGt0RegexString  = `^[1-9]\d*$`
	numberGte0RegexString = `^\d+$`
)

var (
	rxNumberGt0  = regexp.MustCompile(numberGt0RegexString)
	rxNumberGte0 = regexp.MustCompile(numberGte0RegexString)
)
