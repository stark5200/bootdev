package main

func reformat(message string, formatter func(string) string) string {
	firstFormat := formatter(message)
	secondFormat := formatter(firstFormat)
	thirdFormat := formatter(secondFormat)
	return "TEXTIO: "+thirdFormat
}
