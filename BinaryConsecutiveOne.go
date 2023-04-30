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
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	number := int32(nTemp)
	var maximum, current int32 = 0, 0
	maximum = maximumConsecutiveOne(number, maximum, current)
	fmt.Println(maximum)
}

func maximumConsecutiveOne(number int32, maximum int32, current int32) int32 {
	if number < 1 {
		return maximum
	}
	if number%2 == 1 {
		current += 1
	} else {
		current = 0
	}
	if current > maximum {
		maximum = current
	}
	number = number / 2
	return maximumConsecutiveOne(number, maximum, current)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
