Lets say we want to calculate Perimiter of triangle

````
func TestPerimeter(t \*testing.T){

    got:= Perimeter(10.0,10.0)
    want:= 40.0

    if want!=got{
            t.Errorf("got %.2f but want is %.2f",got,want)

        }
}

````
- Struct:

```
type rectangle struct {

          name string
          length int
}
```
- METHOD: A METHOD is a func with a receiver. below is an example 
```
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```
