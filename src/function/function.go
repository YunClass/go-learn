// function project function.go
package function

//函数也是值。他们可以像其他值一样传递，比如，函数值可以作为函数的参数或者返回值

//闭包
func Adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
