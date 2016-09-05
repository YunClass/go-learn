// point project point.go
package point

import (
	"fmt"
)

func Point() {

	i, j := 10, 20

	p := &i         //生成一个指向i对象的指针
	fmt.Println(p)  //打印的是一个引用地址，即内存地址
	fmt.Println(*p) //打印的是i的值

	*p = 11
	fmt.Println(i) //i的值变成了11

	p = &j //生成一个指向j对象的指针
	*p = *p / 2
	fmt.Println(j) //j的值变成10

}
