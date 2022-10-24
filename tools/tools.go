package tools

import "regexp"

// MatchUrl 根据指定的正则匹配
func MatchUrl(pattern, str string) string {
	reg := regexp.MustCompile(pattern)
	res := reg.FindStringSubmatch(str)
	if len(res) == 2 {
		return res[1]
	}
	return ""
}
