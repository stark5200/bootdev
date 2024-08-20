package main

import (
	"fmt"
	"math"
)

func bulkSend(numMessages int) float64 {
	cost := 0.00
	for i := 0; i < numMessages; i++ {
		cost += 1 + float64(i) * 0.01
	} 
	return cost
}

func maxMessages(thresh int) int {
	cost := 0
	max := 0
	for i := 0; ; i++ {
		if cost + 100 + i > thresh {
			return max
		}
		cost += 100 + i
		max += 1
	} 
}

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1
	balance := float64(maxCostInPennies) - actualCostInPennies
	for balance > 0{
		actualCostInPennies *= costMultiplier
		balance -= actualCostInPennies
		maxMessagesToSend++
	}
	if balance < 0 {
		maxMessagesToSend--
	}
	return maxMessagesToSend
}

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		} 
		if i%3 == 0 {
			fmt.Println("fizz")
			continue
		} 
		if i%5 == 0 {
			fmt.Println("buzz")
			continue
		} 
		fmt.Println(i)
	}
}

// don't touch below this line

// fizz buzz test

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		isPrime := true 
	    if n == 2 {
		    fmt.Println(n)
			continue
		}
		if n == 3 {
		    fmt.Println(n)
			continue
		}
	    if n%2 == 0 {
		    continue
		}
		if n%3 == 0 {
		    continue
		}
	    for i:=3; i <= int(math.Floor(math.Sqrt(float64(n)))); i += 2 {
	        if (n % i == 0) {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Println(n)
			continue
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	fizzbuzz()

	test(10)
	test(20)
	test(30)
}

func countConnections(groupSize int) int {
	connections := 0
	for i := 0; i < groupSize; i++ {
		connections += i
	}
	return connections
}
