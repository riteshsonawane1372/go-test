Below is an Golang Test Example

`
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
} 
`

follow the Below Key Points 
1. xxx_test.go
2. Test function must start with "Test" func and It should take "*testing.T" as a argument
3. Import testing

"godoc -http=localhost:8000" you can launch a local go docs. Go and check the port
'localhost:8000/pkg' will show all the pkg installed on our system

Below is how we passed argument to func Hello, See the Test

` 
package main

import "testing"

func TestHello(t *testing.T) {
  
  t.Run("saying hello to people", func(t *testing.T){
  
    got:= Hello("Ritesh")
    want:= "Hello, Ritesh"
    
    if got !=want{
      t.Errorf("got %q but want is %q", got,want)
    }

  })

  t.Run("No Arg passed in Hello", func (t *testing.T){
  
    got:= Hello("")
    want:= "Hello, "

    if got!=want{
      t.Errorf("got %q but want %q",got,want)
    }


  })

} 
`

Now lets Refactor this Test 

`
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
` 

Now try to fail the test and see the output, for now we can comment `t.Helper`

Note: It is imp to see the Test fail, what msg it shows 
