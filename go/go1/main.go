package main

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
