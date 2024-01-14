package main
import ("fmt" 
        "time")

func power(base int, power int) int {
  var x int = base
  for i:=1; i<= power; i++ {
    x *= base
  }
  return x
}

func main() {
  t0 := time.Now()
  var n int = 0
  var limit int = power(10,8)
  for i:=0; i <= limit; i++ {
    n +=i
  }
  t1 := time.Now()
  fmt.Println(n)
  fmt.Println("It took %v to run.\n", t1.Sub(t0))
}
