package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	numberOfTest, err := strconv.ParseInt(readString(reader), 10, 64)
	checkError(err)
	arrayOfStrings := make([]string, numberOfTest)
	for i := 0; int64(i) < numberOfTest; i += 1 {
		arrayOfStrings[i] = readString(reader)
	}
	//fmt.Println(numberOfTest);
	runTestOnStrings(arrayOfStrings)
}

func runTestOnStrings(testCases []string) {
	for i := 0; i < len(testCases); i += 1 {
		var even string
		var odd string
		for j, r := range testCases[i] {
			if j%2 == 0 {
				even += string(r)
			} else {
				odd += string(r)
			}
		}
		fmt.Println(even, odd)
	}

	//for i, r := range str {
	//   fmt.Println(i, r, string(r))
	// }

}

func readString(reader *bufio.Reader) string {
	readingInput, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(readingInput), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
