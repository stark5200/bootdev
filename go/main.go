package main

func getMonthlyPrice(tier string) int {
	if tier == "basic" {
		return 100
	}
	return 200
}