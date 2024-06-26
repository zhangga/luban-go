package utils

func MakeFullName(module, name string) string {
	if len(module) == 0 {
		return name
	}
	if len(name) == 0 {
		return module
	}
	return module + "." + name
}

func Contain[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
