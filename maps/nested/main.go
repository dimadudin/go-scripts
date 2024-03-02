package main

func getNameCounts(names []string) map[rune]map[string]int {
	res := make(map[rune]map[string]int)
	for _, name := range names {
		firstChar := rune(name[0])
		_, charExists := res[firstChar]
		if !charExists {
			res[firstChar] = map[string]int{name: 1}
			continue
		}
		_, nameExists := res[firstChar][name]
		if !nameExists {
			res[firstChar][name] = 1
			continue
		}
		res[firstChar][name]++
	}
	return res
}
