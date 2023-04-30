package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "io"
)


func main() {
 //Enter your code here. Read input from STDIN. Print output
    reader := bufio.NewReaderSize(os.Stdin, 1024*1024);   
    numberOfTest, err := strconv.Atoi(readString(reader));
    checkError(err);
   // arrayOfStrings := make([]string, numberOfTest)
    
    for i := 0;  i < numberOfTest; i += 1 {
       runTestOnStrings(readString(reader));
    }
    //fmt.Println(numberOfTest);
    // 
}

func runTestOnStrings(testCases string) {
    stringLength := len(testCases);
    
    for j := 0; j < stringLength; j += 2 {
        fmt.Fprintf(os.Stdout, "%s", string(testCases[j]));
    }   
     
    fmt.Fprintf(os.Stdout, "%s"," ");
    
    for j := 1; j < stringLength; j += 2 {
         fmt.Fprintf(os.Stdout, "%s",string(testCases[j]));
    }
    
    fmt.Fprintf(os.Stdout, "%s","\n");
    
    
    //for i, r := range str {
     //   fmt.Println(i, r, string(r))
    // }

}


// func runTestOnStrings(testCases []string) {
//    for i := 0; i < len(testCases); i += 1 {
      
 //       for j :=0; j < len(testCases[i]); j += 2 {
   //         fmt.Fprintf(os.Stdout, "%s",string(testCases[i][j]));
       // }   
     //   fmt.Fprintf(os.Stdout, "%s"," ");
       // for j :=1; j < len(testCases[i]); j += 2 {
       //     fmt.Fprintf(os.Stdout, "%s",string(testCases[i][j]));
       // }
       // fmt.Fprintf(os.Stdout, "%s","\n");
   // }
    
    //for i, r := range str {
     //   fmt.Println(i, r, string(r))
    // }

// }


func readString(reader *bufio.Reader) string {
    readingInput, _, err := reader.ReadLine();
    if err == io.EOF {
        return "";
    }
    return strings.TrimRight(string(readingInput),"\r\n");
}

func checkError(err error) {
    if err != nil {
        panic(err);
    }
}
