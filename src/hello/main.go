// hello project main.go
package main

//引入多个包
import (
	"arrayLearn"
	"conf"
	"control"
	"fmt"
	"function"
	"interfaces"
	"method"
	"myError"
	"point"
	"structObject"
	"time"
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
	//const j := 11  //常量不能使用 :=
	//j = "222" 常量不能重新赋值
	fmt.Println(j)

	//类型转换
	var k int = 12
	//var l float32 = k    //不能显示转换
	var l float32 = float32(k)
	fmt.Println(l)

	//类型自动推导
	m := 11.11
	fmt.Printf("m is of type %T\n", m)

	fmt.Println(time.Now())

	//控制语句
	control.SwitchControl()
	control.DeferControl()

	//指指针
	point.Point()

	//结构体
	structObject.EchoPerson()
	structObject.StructGrammar()

	//数组
	arrayLearn.Array()
	arrayLearn.Slice()
	arrayLearn.Map()

	//函数
	pos := function.Adder()
	fmt.Println(pos) //内存地址
	fmt.Println(pos(1))

	//方法
	method.Method()

	//接口
	interfaces.Interfaces()

	//错误
	myError.PrintError()

}
