package main


import "fmt"

func main(){

  nums:= []int{1,2,3,4,5}
  fmt.Println(Sum(nums))
  
}

func Sum(arr[] int)int{

  sum:=0
  for i:=0;i<len(arr);i++{
    sum=sum+arr[i]
  }

  return sum
  
  

}
