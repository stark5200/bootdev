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

/*
signatures

x int
p *int
a [3]int

f func(func(int,int) int, int) int
*/

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	getBillForMonth(lastMonthBill, costPerSend, numLastMonth)
	getBillForMonth(thisMonthBill, costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(bill, costPerSend, messagesSent int) {
	bill = costPerSend * messagesSent
}
