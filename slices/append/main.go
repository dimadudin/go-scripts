package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	dailyCosts := make([]float64, 0)
	for _, v := range costs {
		if v.day < len(dailyCosts) {
			dailyCosts[v.day] += v.value
		} else {
			for i := len(dailyCosts); i < v.day; i++ {
				dailyCosts = append(dailyCosts, 0.0)
			}
			dailyCosts = append(dailyCosts, v.value)
		}
	}
	return dailyCosts
}
