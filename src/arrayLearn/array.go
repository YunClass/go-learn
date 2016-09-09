// array project array.go
package arrayLearn

import (
	"fmt"
)

func Array() {

	var a [2]string
	a[0] = "ttt"
	a[1] = "sss"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

func Slice() {

	slice := []int{2, 3, 4, 5}
	fmt.Println(slice)

	for k, v := range slice {
		fmt.Println(k, v)
	}

	fmt.Println(len(slice)) //slice的长度

	//对slice切片，创建一个新的slice指向相同的数组
	fmt.Println(slice[1:2]) //包含1，不包含2

	//构造slice 由函数 make 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组
	a := slice[1:2]
	fmt.Println(cap(a)) //a截取的是slice第1位之后的长度 修改slice的值会修改整个数组的值

	//slice的零值是nil
	var z []int
	if z == nil {
		fmt.Println("nil")
	}

	//像slice添加元素
	z = append(z, 2) //第一个元素类型决定之后的类型
	fmt.Println(z)

	z = append(z, 1, 3, 33)
	fmt.Println(z, len(z), cap(z))

}

type Person struct {
	name, sex string
}

func Map() {

	m := make(map[string]Person)
	m["a"] = Person{
		"张三", "男",
	}

	fmt.Println(m)

	//文法与结构体相似
	var person1 = map[string]Person{
		"a": Person{"张三", "男"},
		"b": Person{"李四", "女"},
	}
	fmt.Println(person1)

	//若顶级类型相同可以省略
	var person2 = map[string]Person{
		"a": {"掌声呢", "是你"},
		"b": {"da", "sss"},
	}
	fmt.Println(person2)

	//获得元素
	fmt.Println(person2["a"])

	//删除元素
	delete(person2, "a")
	fmt.Println(person2)

	//双赋值，检测某个值是否存在
	v, ok := person2["a"]
	fmt.Println(v, ok)

}
