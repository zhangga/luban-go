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

func FileExtWithoutDot(fullName string) string {
	if index := strings.LastIndex(fullName, "."); index >= 0 {
		return fullName[index+1:]
	}
	return ""
}

func IsExcelFile(fullName string) bool {
	return strings.HasSuffix(fullName, ".csv") ||
		strings.HasSuffix(fullName, ".xls") ||
		strings.HasSuffix(fullName, ".xlsx") ||
		strings.HasSuffix(fullName, ".xlsm")
}
