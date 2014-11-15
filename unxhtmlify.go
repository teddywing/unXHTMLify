package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading standard input", err)
	}
}

func unxhtmlify_line(s string) string {
	re := regexp.MustCompile(" ?/>")
	return re.ReplaceAllString(s, ">")
}
