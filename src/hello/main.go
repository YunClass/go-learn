// hello project main.go
package main

//引入多个包
import (
	"conf"
	"fmt"
)

//定义多个变量
var c, d bool

//初始化多个变量
var a, b int = 1, 3

func main() {
	fmt.Println("Hellssso World!")
	val := conf.Add(1, 3)
	fmt.Println(val)
	fmt.Println(conf.Swap("第一个", "第二个人"))
	fmt.Println(conf.Split(3))

	//i := 10 声明未被引用编译会报错
	// c d 都未被使用不会报错
	fmt.Println(a)

	//短声明
	e, f := "eee", "ffff"
	fmt.Println(e, f)

	//短申明多个不同类型的变量
	g, h := "ggg", true
	fmt.Println(g, h)

	const j = "1111%"
	fmt.Println(j)

	//j = "222" 常量不能重新赋值

	//类型转换
	var k int = 12
	//var l float32 = k    //不能显示转换
	var l float32 = float32(k)
	fmt.Println(l)

}
