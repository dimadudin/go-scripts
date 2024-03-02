package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, v := range msg {
		for _, b := range badWords {
			if v == b {
				return i
			}
		}
	}
	return -1
}
