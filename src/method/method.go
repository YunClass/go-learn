// method project method.go
package method

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
	Age  int
}

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func (person *Person) SayName() {

	fmt.Println(person.Name)

}

func Method() {

	person := &Person{"张三", 10}
	person.SayName()

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.abs())

}
