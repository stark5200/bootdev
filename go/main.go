package main

func getMonthlyPrice(tier string) int {
	price := 0
	if tier == "basic" {
		price = 100
	}
	if tier == "premium" {
		price = 150
	}
	if tier == "enterprise" {
		price = 500
	}
	return price*100
}
