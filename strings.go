package hotels

func LongestString(ss []string) string {
	l := ""
	for _, s := range ss {
		if len(s) > len(l) {
			l = s
		}
	}
	return l
}
