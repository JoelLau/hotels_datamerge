package hotels

import "strings"

func clean(s string, fns ...func(string) string) string {
	s = strings.TrimSpace(s)
	for _, fn := range fns {
		s = fn(s)
	}
	return s
}
