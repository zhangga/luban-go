package utils

import (
	"strings"
)

func StandardizePath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

func SplitFileAndSheetName(url string) (string, string) {
	sheetSepIndex := strings.Index(url, "@")
	if sheetSepIndex == -1 {
		return url, ""
	}

	lastPathSep := strings.LastIndex(url[:sheetSepIndex], "/")
	if lastPathSep >= 0 {
		return url[lastPathSep+1 : sheetSepIndex], url[sheetSepIndex+1:]
	}

	return url[:sheetSepIndex], url[sheetSepIndex+1:]
}
