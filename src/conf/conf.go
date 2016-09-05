// conf project conf.go
package conf

func Add(x, y int) int {
	return x + y
}

//多返回值
func Swap(x, y string) (string, string) {
	return x, y
}

//命名多返回值
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum + 2
	return
}
