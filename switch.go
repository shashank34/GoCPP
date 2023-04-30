package main
import("fmt")
func main() {
  var x int
  var y *int
  z := 3
  y = &z
fmt.Println(x, *y)
  //x = &y
 var xtemp int
  x1 := 0
  x2 := 1
  for x:=0; x<5; x++ {
    xtemp = x2
    x2 = x2 + x1
    x1 = xtemp
  }
  fmt.Println(x2)
}
