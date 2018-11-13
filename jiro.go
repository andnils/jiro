package main

import "fmt"
import "os"
import "regexp"
import "strings"

func jirify(s string) string {
	isNumber, _ := regexp.MatchString("^\\d+$", s)
	if isNumber {
		return fmt.Sprintf("WH-%s", s)
	} else {
		return strings.ToUpper(s)
	}
}

func urlify(s string) string {
	return "https://your.jira.url/browse/" + jirify(s)
}

func main() {
	if len(os.Args) > 1 {
		fmt.Printf(urlify(os.Args[1]))
	} else {
		os.Exit(1)
	}
}
