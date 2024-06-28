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

func ComputeCfgHashIdByName(name string) int64 {
	var id int64
	for _, c := range name {
		id = 31*id + int64(c)
	}
	return id
}
