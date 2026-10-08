package function

import "slices"

func CheckInNumberSlice[T uint64 | int32](a T, s []T) bool {
	return slices.Contains(s, a)
}

// 只支持元素不重复的情景
func DelEleInSlice[T uint64 | int32](a T, old []T) (new []T) {
	for i, val := range old {
		if val == a {
			new = append(old[:i], old[i+1:]...)
			return
		}
	}
	return old
}
