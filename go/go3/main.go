package main

func reformat(message string, formatter func(string) string) string {
	return message + formatter("")
}
