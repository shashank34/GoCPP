func workingWithMaps {
    // var m = map[string]int{}; identical to make(map[string]int);
    // fmt.Println(m);
    // delete(m, "raaste")
    // i return value of key and ok is bool if key exists or not
    // i, ok := m["route"]
    // to iterate Over the map
    // for key, value := range m {
    //  fmt.Println("Key:", key, "Value:", value)
    // }
    
    // hits := make(map[string]map[string]int)
    //type Person struct {
    //    Name  string
      //   Likes []string
    // }
    // var people []*Person

     // likes := make(map[string][]*Person)
    // for _, p := range people {
        // for _, l := range p.Likes {
       //      likes[l] = append(likes[l], p)
      //  }
    //}
    
    // for _, p := range likes["cheese"] {
    //    fmt.Println(p.Name, "likes cheese.")
    // }
    
   //  fmt.Println(len(likes["bacon"]), "people like bacon.")
/*  Concurrency
Maps are not safe for concurrent use: it's not defined what happens when you read and write to them simultaneously. If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by some kind of synchronization mechanism. One common way to protect maps is with sync.RWMutex.

This statement declares a counter variable that is an anonymous struct containing a map and an embedded sync.RWMutex.

var counter = struct{
    sync.RWMutex
    m map[string]int
}{m: make(map[string]int)}
To read from the counter, take the read lock:

counter.RLock()
n := counter.m["some_key"]
counter.RUnlock()
fmt.Println("some_key:", n)
To write to the counter, take the write lock:

counter.Lock()
counter.m["some_key"]++
counter.Unlock() 


Iteration order
When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next. Since Go 1 the runtime randomizes map iteration order, as programmers relied on the stable iteration order of the previous implementation. If you require a stable iteration order you must maintain a separate data structure that specifies that order. This example uses a separate sorted slice of keys to print a map[int]string in key order:

import "sort"

var m map[int]string
var keys []int
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)
for _, k := range keys {
    fmt.Println("Key:", k, "Value:", m[k])
}


   */
 
}
