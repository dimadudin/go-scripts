package main

func maxMessages(thresh int) int {
	if thresh < 100 {
		return 0
	}
	cost := 0
	i := 0
	for ; cost < thresh; i++ {
		cost += 100 + i
	}
	return i - 1
}
