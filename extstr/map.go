package extstr

// Map 本地修改原ss
func Map(ss []string, mapping func(string) string) []string {
	for i := 0; i < len(ss); i++ {
		ss[i] = mapping(ss[i])
	}
	return ss
}

// Mapx 新建新的,并执行map
func Mapx(ss []string, mapping func(string) string) []string {
	ssx := make([]string, 0, len(ss))
	for _, s := range ss {
		ssx = append(ssx, mapping(s))
	}
	return ssx
}
