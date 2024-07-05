package utils

import "strings"

const (
	attrSep = "#"
)

func ParseAttrs(tags string) map[string]string {
	attrs := make(map[string]string)
	if len(tags) == 0 {
		return attrs
	}

	braceDepth, pairStart := 0, 0
	for i := 0; i < len(tags); i++ {
		c := tags[i]
		if c == '(' || c == '[' || c == '{' {
			braceDepth++
		} else if c == ')' || c == ']' || c == '}' {
			braceDepth--
		}

		if braceDepth == 0 && c == '#' {
			rawPair := tags[pairStart : i-pairStart]
			pairStart = i + 1
			AddAttr(attrs, rawPair)
		}
	}
	if braceDepth != 0 {
		panic("unbalanced brace pairs in tags: " + tags)
	}
	if pairStart < len(tags) {
		AddAttr(attrs, tags[pairStart:])
	}
	return attrs
}

const attrKeyValueSep = "=:"

func AddAttr(attrs map[string]string, rawPair string) {
	pair := TrimBracePairs(rawPair)
	sepIndex := strings.IndexAny(pair, attrKeyValueSep)
	var key, value string
	if sepIndex >= 0 {
		key = strings.TrimSpace(pair[:sepIndex])
		value = strings.TrimSpace(pair[sepIndex+1:])
	} else {
		trimmedPair := strings.TrimSpace(pair)
		key = trimmedPair
		value = trimmedPair
	}
	attrs[key] = value
}

// TrimBracePairs 去除字符串两端成对的圆括号
func TrimBracePairs(rawPair string) string {
	for len(rawPair) > 0 && rawPair[0] == '(' {
		braceDepth := 0
		level1Left := -1
		level1Right := -1
		for i, char := range rawPair {
			if char == '(' {
				braceDepth++
				if level1Left < 0 {
					level1Left = i
				}
			}
			if char == ')' {
				braceDepth--
				if level1Right < 0 && braceDepth == 0 {
					level1Right = i
					break
				}
			}
		}
		if level1Left >= 0 && level1Right == len(rawPair)-1 {
			rawPair = rawPair[level1Left+1 : level1Right]
		} else {
			break
		}
	}
	return rawPair
}

func IsNormalFieldName(name string) bool {
	return !strings.HasPrefix(name, "__") && !strings.HasPrefix(name, "#") && !strings.HasPrefix(name, "$")
}

func IndexOfBaseTypeEnd(s string) int {
	braceDepth := 0
	firstSharpIndex := -1 // '#'
	for i, c := range s {
		switch c {
		case '(', '[', '{':
			braceDepth++
		case ')', ']', '}':
			braceDepth--
		}
		if c == '#' && firstSharpIndex == -1 {
			firstSharpIndex = i
		}

		if braceDepth == 0 && (c == ',' || c == ';') {
			var strContainBaseType string
			if firstSharpIndex > 0 {
				strContainBaseType = s[:firstSharpIndex]
			} else {
				strContainBaseType = s[:i]
			}
			strContainBaseType = strings.ReplaceAll(strContainBaseType, "(", "")
			strContainBaseType = strings.ReplaceAll(strContainBaseType, ")", "")
			strContainBaseType = strings.ReplaceAll(strContainBaseType, "[", "")
			strContainBaseType = strings.ReplaceAll(strContainBaseType, "]", "")

			if strContainBaseType == "array" || strContainBaseType == "list" || strContainBaseType == "set" || strContainBaseType == "map" {
				return i
			} else {
				return -1
			}
		}
	}
	return -1
}

func ParseType(s string) (string, map[string]string) {
	sepIndex := strings.IndexAny(s, attrSep)
	if sepIndex < 0 {
		return s, map[string]string{}
	} else {
		braceDepth := 0
		for i, c := range s {
			switch c {
			case '(', '[', '{':
				braceDepth++
			case ')', ']', '}':
				braceDepth--
			}

			if braceDepth == 0 && c == '#' {
				return s[:i], ParseAttrs(s[i+1:])
			}
		}
		return s, make(map[string]string)
	}
}

func ParseTypeAndValidAttrs(s string) (string, map[string]string) {
	typeStr, attrs := ParseType(s)
	if len(attrs) > 0 {
		if _, ok := attrs["group"]; ok {
			panic("group为保留属性,只能用于table或var定义,是否用错? 如在excel中请使用&group=xxx")
		}
		if _, ok := attrs["seq"]; ok {
			panic("字段切割应该用'sep'，而不是'seq',请检查是否拼写错误")
		}
	}
	return typeStr, attrs
}
