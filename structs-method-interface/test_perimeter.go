package main

import "testing"

func TestPerimeter(t *testing.T){
  
  got:= Perimeter(10.0,10.0)
  want:= 40.0

  if want!=got{
    t.Errorf("got %.2f but want is %.2f",got,want)
  }


}
