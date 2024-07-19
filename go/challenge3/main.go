package main

func adder() func(int) int {
	sum := 0
	return func(count int) int {
		sum += count
		return sum
	}
}
