package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
)

// Complete the solve function below.
func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
    var tip float64 = float64(tip_percent) * meal_cost/100;
    var tax float64 = float64(tax_percent) * meal_cost/100;
    var totalCost float64;
    totalCost = meal_cost + tip + tax;
    totalCost = math.Round(totalCost);
    fmt.Println("The total meal cost is",totalCost,"dollars."); 
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    meal_cost, err := strconv.ParseFloat(readLine(reader), 64)
    checkError(err)

    tip_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tip_percent := int32(tip_percentTemp)

    tax_percentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    tax_percent := int32(tax_percentTemp)

    solve(meal_cost, tip_percent, tax_percent)
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
