package hotels

import (
	"regexp"
	"strings"
)

var pascalCaseBoundary = regexp.MustCompile(`([a-z0-9])([A-Z])`)

func ToLowerCaseWithSpaces(s string) string {
	s = pascalCaseBoundary.ReplaceAllString(s, "$1 $2")
	return strings.ToLower(s)
}

func ToNilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

// func ToUniqueSlice(ss []string) []string {
// 	if len(ss) <= 0 {
// 		return ss
// 	}

// 	result := make([]string, 0, len(ss))
// 	seen := make(map[string]struct{})
// 	for _, s := range ss {
// 		if _, ok := seen[s]; !ok {
// 			seen[s] = struct{}{}
// 			result = append(result, s)
// 		}
// 	}

// 	return slices.Collect(maps.Keys(seen))
// }

func LongestString(ss []string) string {
	l := ""
	for _, s := range ss {
		if len(s) > len(l) {
			l = s
		}
	}
	return l
}
