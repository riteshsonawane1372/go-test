package main

import "testing"

func TestHello(t *testing.T) {
  
  t.Run("saying hello to people", func(t *testing.T){
  
    got:= Hello("Ritesh")
    want:= "Hello, Ritesh"
    
    assertCorrectMessage(t,got,want)

  })

  t.Run("No Arg passed in Hello", func (t *testing.T){
  
    got:= Hello("")
    want:= "Hello, "

    assertCorrectMessage(t,got,want)


  })

}

func assertCorrectMessage(t testing.TB, got, want string){
  
  t.Helper()
  if got!=want{
    t.Errorf("got %q but want is %q", got,want)
  }

}
