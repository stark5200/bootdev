package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	// ?
}

func calculateDiscount(messagesSent int) float64 {
	// ?
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
