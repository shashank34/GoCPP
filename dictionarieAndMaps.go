package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "io"
    "sync"
)

/* 

type PhoneBook struct  {
    name string
    phoneNumber int64
}

*/
var Dictionary = struct {
    sync.RWMutex
    KeyMap map[string]string
} { KeyMap: make(map[string]string) }


func main() {
 // Enter your code here. Read input from STDIN. Print output to STDOUT
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024);
    totalCount, err := strconv.Atoi(readText(reader));
    checkError(err);
    // fmt.Fprintf(os.Stdout, "%d", totalCount);
    for i := 0; i < totalCount; i += 1 {
        phoneBookInput := readInputOfPhoneBook(reader);
        setDictionary(phoneBookInput);
    }
    queries(reader);   
}

func queries(reader *bufio.Reader) {
    for true {
        input := readText(reader)
        if (input == "") {
            break;
        }
        Dictionary.RLock();
        value,isPresent := Dictionary.KeyMap[input];
        Dictionary.RUnlock();
        if isPresent {
            fmt.Fprintf(os.Stdout,"%s=%s\n",input, value);
        } else {
            fmt.Fprintln(os.Stdout, "Not found");
        }
    }
}


func setDictionary(text []string)  {
   Dictionary.Lock();
   Dictionary.KeyMap[text[0]] = text[1];
   Dictionary.Unlock();
}


func readInputOfPhoneBook(reader *bufio.Reader) []string {
    input, _, err := reader.ReadLine();
    if err == io.EOF {
        var s []string;
        return s;
    }
    text := strings.TrimRight(string(input), "\r\n");
    return strings.Fields(text);
}

func readText(reader *bufio.Reader) string {
    input, _, err := reader.ReadLine();
    if err == io.EOF {
        return ""
    }
    return strings.TrimRight(string(input), "\r\n");
}

func checkError(err error) {
    if err != nil {
        panic(err);
    }
}
