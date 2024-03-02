package main

func getCounts(usernames []string) map[string]int {
	counts := make(map[string]int)
	for _, name := range usernames {
		_, ok := counts[name]
		if !ok {
			counts[name] = 1
		} else {
			counts[name]++
		}
	}
	return counts
}
