We can initialize arr in two ways 
`nums:= [5]int{1,2,3,4,5}`
`nums:=[...]int{1,2,3,4,5}` 

There are two cases 

1. `func Sum(arr[5]int)`: Here function needs an argument of arr of size `5`
2. `func Sum(arr[]int)` : Here function needs an argument of arr of size [...]

If we dont pass an correct argument then the testcase will fail

`go test -cover`: Will give the test coverage
