package extstr

var goKeywords = map[string]struct{}{
	"var":         {},
	"const":       {},
	"package":     {},
	"import":      {},
	"func":        {},
	"return":      {},
	"defer":       {},
	"go":          {},
	"select":      {},
	"interface":   {},
	"struct":      {},
	"break":       {},
	"case":        {},
	"continue":    {},
	"for":         {},
	"fallthrough": {},
	"else":        {},
	"if":          {},
	"switch":      {},
	"goto":        {},
	"default":     {},
	"chan":        {},
	"type":        {},
	"map":         {},
	"range":       {},
}

var goInternalType = map[string]struct{}{
	"string":     {},
	"bool":       {},
	"int":        {},
	"uint":       {},
	"byte":       {},
	"rune":       {},
	"int8":       {},
	"int16":      {},
	"int32":      {},
	"int64":      {},
	"uint8":      {},
	"uint16":     {},
	"uint32":     {},
	"uint64":     {},
	"uintptr":    {},
	"float32":    {},
	"float64":    {},
	"map":        {},
	"complex64":  {},
	"complex128": {},
	"iota":       {},
	"nil":        {},
	"Time":       {},
}

// IsGoKeywords 是否是Golang关键字
func IsGoKeywords(t string) bool {
	_, exist := goKeywords[t]
	return exist
}

// IsGoInternalType 是否是Golang内部类型
func IsGoInternalType(t string) bool {
	_, exist := goInternalType[t]
	return exist
}
