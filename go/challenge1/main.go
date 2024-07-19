package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	return costPerMessage * float64(numMessages) * calculateDiscount(numMessages)
}

func calculateDiscount(messagesSent int) float64 {
	discount := 0
	if messagesSent > 1000 {
		discount = 10
	}
	if messagesSent > 5000 {
		discount = 20
	}
	return float64(float64(100 - discount)/float64(100))
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}