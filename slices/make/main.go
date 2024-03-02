package main

func getMessageCosts(messages []string) []float64 {
	meassageCosts := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		meassageCosts[i] = 0.01 * float64(len(messages[i]))
	}
	return meassageCosts
}
