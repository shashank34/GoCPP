package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main()  {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		word := input.Text()
		sanitize := strings.TrimSpace(strings.ToLower(word))
		if strings.HasPrefix(sanitize, "i") {
			if strings.Contains(sanitize, "a") {
				if strings.HasSuffix(sanitize, "n") {
					fmt.Println("Found!")
				}  else {
					fmt.Println("Not Found!")
				}
			}  else {
				fmt.Println("Not Found!")
			}
		} else {
			fmt.Println("Not Found!")
		}
	}
}
