// control project control.go
package control

import (
	"fmt"
	"math"
	"time"
)

func ForControl() {

	sum := 1

	for i := 0; sum < 100; i++ {
		fmt.Print(i)
	}

	//相当于while循环
	for sum < 100 {
		sum += sum
	}

	fmt.Println(sum)

}

func Sqrt(x float64) float64 {
	if x < 0 {
		return Sqrt(-x)
	}

	return math.Sqrt(x)
}

func Pow(x, n, lim float64) float64 {
	//if  可以在条件执行之前执行一段语句
	if v := math.Pow(x, n); lim < 0 {
		return v
	}

	return lim
}

func SwitchControl() {

	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func DeferControl() {

	//延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用
	defer fmt.Println("world")

	fmt.Println("hello")

	//延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("finish")
}
