We will be writing a func which will accept a char and repeat it for 5 times

## writing Benchmarks in Golang
Benchmarks is similar to writing test in Golang
Below is the code for Golang test 
`
func Repeat(n string) string{
  a:=n  
  for i:=0;i<4;i++{
    n=n+a
  }

  return n
}
`

Below code is for Benchmarks
`
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
`
`b.N` times it runs and measures how long it takes 
`$ go test -bench=.` will run the Benchmarks

Note: By default Benchmarks are run sequentially 
