Below is an Golang Test

package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

follow the Below Key Points 
1. xxx_test.go
2. Test function must start with "Test" func and It should take "*testing.T" as a argument
3. Import testing

"godoc -http=localhost:8000" you can launch a local go docs. Go and check the port
'localhost:8000/pkg' will show all the pkg installed on our system
