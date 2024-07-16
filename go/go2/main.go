package main

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) (increase int) {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	increase = thisMonthBill - lastMonthBill
	return
}

func getBillForMonth(costPerSend, messagesSent int) (bill int) {
	bill = costPerSend * messagesSent
	return
}
