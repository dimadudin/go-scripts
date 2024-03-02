package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	primCost := len(primary)
	secCost := primCost + len(secondary)
	tertCost := secCost + len(tertiary)
	return [3]string{primary, secondary, tertiary}, [3]int{primCost, secCost, tertCost}
}
