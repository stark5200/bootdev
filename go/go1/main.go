package main


import (
	"fmt"
	"unicode/utf8"
)

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


func main() {
	const name = "boots"
	fmt.Printf("'name' byte length: %d\n", len(name))
	fmt.Printf("'name' rune length: %d\n", utf8.RuneCountInString(name))
	fmt.Println("=====================================")
	fmt.Printf("Hi %s, so good to have you back in the arcanum\n", name)
}

