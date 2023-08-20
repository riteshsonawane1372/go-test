package main

import "testing"

func TestSum(t *testing.T){
  nums:= []int{1,2,3,4,5}
  t.Run("Collection of any Size", func(t *testing.T){
  
    got:= Sum(nums)
    want:= 15
    if got!=want{
      t.Errorf("got %d but want is %d",got,want)
    }

  })
  

}

