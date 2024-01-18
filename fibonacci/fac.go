package main
import "fmt"

func fac(n int) int {
  if n==1{
    return 1
  }

  return n*fac(n-1)
}

func main() {
  var res int = fac(10)
  fmt.Println(res)
}
