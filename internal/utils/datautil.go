package utils

import "strings"

func IsIgnoreTag(tag string) bool {
	return tag == "##"
}

func ParseTags(rawTagStr string) []string {
	if len(rawTagStr) == 0 {
		return nil
	}

	var tags []string
	for _, t := range strings.Split(rawTagStr, ",") {
		t = strings.TrimSpace(t)
		if len(t) == 0 {
			continue
		}
		tags = append(tags, strings.ToLower(t))
	}
	return tags
}
