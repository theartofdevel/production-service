package array

import (
	"sort"
)

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func ContainsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func RemoveByIndexInt(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func RemoveByValueInt(s []int, el int) []int {
	index := GetElementIndexInt(s, el)
	return RemoveByIndexInt(s, index)
}

func GetElementIndexInt(s []int, el int) int {
	a := sort.IntSlice(s[0:])
	sort.Sort(a)

	return sort.SearchInts(a, el)
}

func AreStringArraysIdentically(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	// create a map of string -> int
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		// 0 value for int is 0, so just increment a counter for the string
		diff[_x]++
	}

	for _, _y := range y {
		_ys := _y
		// If the string _y is not in diff bail out early
		if _, ok := diff[_ys]; !ok {
			return false
		}

		diff[_ys] -= 1
		if diff[_ys] == 0 {
			delete(diff, _ys)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}
