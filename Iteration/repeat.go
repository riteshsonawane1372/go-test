package main

import "fmt"

func Repeat(n string) string{
  a:=n  
  for i:=0;i<4;i++{
    n=n+a
  }

  return n
}

func main(){
  fmt.Println(Repeat("a"))
}
