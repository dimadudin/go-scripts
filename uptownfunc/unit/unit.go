package main

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 10000
	}
	if tier == "premium" {
		return 15000
	}
	if tier == "enterprise" {
		return 50000
	}
	return 0
}
