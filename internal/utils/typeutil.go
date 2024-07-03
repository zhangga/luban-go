package utils

import "strings"

func MakeFullName(module, name string) string {
	if len(module) == 0 {
		return name
	}
	if len(name) == 0 {
		return module
	}
	return module + "." + name
}

func MakeNamespace(module, subModule string) string {
	if len(module) == 0 {
		return subModule
	}
	if len(subModule) == 0 {
		return module
	}
	return module + "." + subModule
}

func Contain[T comparable](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ContainKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func Any[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}
	return false
}

func All[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func FirstOrDefault[T any](arr []T, predicate func(T) bool, defaultValue T) T {
	for _, v := range arr {
		if predicate(v) {
			return v
		}
	}
	return defaultValue
}

func FirstOfList[T any](arr []T, predicate func(T) bool) T {
	for _, v := range arr {
		if predicate(v) {
			return v
		}
	}
	var t T
	return t
}

func ComputeCfgHashIdByName(name string) int64 {
	var id int64
	for _, c := range name {
		id = 31*id + int64(c)
	}
	return id
}

func ToCsStyleName(name string) string {
	var sb strings.Builder
	for _, s := range strings.Split(name, "_") {
		if len(s) == 0 {
			continue
		}
		sb.WriteString(strings.ToUpper(string(s[0])))
		sb.WriteString(s[1:])
	}
	return sb.String()
}

func GetValueOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}
