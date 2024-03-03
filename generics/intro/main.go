package main

func getLast[T any](s []T) T {
	if len(s) < 1 {
		var res T
		return res
	}
	return s[len(s)-1]
}
