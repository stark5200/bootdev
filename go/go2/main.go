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

func getInsuranceAmount(status insuranceStatus) int {
  amount := 0
  if !status.hasInsurance(){
    amount = 1
  } else {
    if status.isTotaled(){
      amount = 10000
    } else {
      if status.isDented(){
        amount = 160
        if status.isBigDent(){
          amount = 270
        }
      } else {
        amount = 0
      }
    }
  }
  return amount
}
