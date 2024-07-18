package main

func reformat(message string, formatter func(string) string) string {
	firstFormat := formatter(message)
	secondFormat := formatter(firstFormat)
	thirdFormat := formatter(secondFormat)
	return "TEXTIO: "+thirdFormat
}

func conversions(converter func(int) int, x, y, z int) (int, int, int) {
	convertedX := converter(x)
	convertedY := converter(y)
	convertedZ := converter(z)
	return convertedX, convertedY, convertedZ
}

func main() {
	newX1, newY1, newZ1 := conversions(func(a int) int {
	    return a + a
	}, 1, 2, 3)
	// newX is 2, newY is 4, newZ is 6

	newX2, newY2, newZ2 := conversions(func(a int) int {
	    return a * a
	}, 1, 2, 3)
	// newX is 1, newY is 4, newZ is 9

	print(newX1, newY1, newZ1, newX2, newY2, newZ2)
}
