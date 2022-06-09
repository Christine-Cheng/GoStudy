package main

import "fmt"

const (
	beijing  = 1
	shanghai = 2
	shenzhen = 3
)

const (
	beijing1 = 10 * iota
	shanghai1
	shenzhen1
)

//声明变量
func main() {
	//var关键字进行声明变量
	//方法一:声明一个变量默认值是0
	var a int
	fmt.Println("a=", a)
	fmt.Printf("a的类型是:%T\n", a)

	//方法二:声明变量并且初始化值
	var b int = 100
	fmt.Printf("b=%d,b的类型是:%T\n", b, b)

	//方法三:声明变量的时候去掉数据类型,go语言通过数值进行匹配类型
	var c = 105
	fmt.Printf("c=%d,c的类型是:%T\n", c, c)

	var cc = "Go语言是一种胶水语言"
	fmt.Printf("cc=%s,cc的类型是:%T\n", cc, cc)

	//方法4:"短声明" := 进行声明
	dd := 99
	fmt.Printf("dd=%d,dd的类似是:%T\n", dd, dd)

	e := "Go语言天生高并发"
	fmt.Printf("e=%s,e的类型是:%T\n", e, e)

	//同时声明多个变量
	var x, y = 100, "奥利奥好吃" //可以支持同时声明不同类型的变量
	fmt.Printf("x=%d,x的类型是:%T\n", x, x)
	fmt.Printf("y=%s,y的类型是:%T\n", y, y)

	var (
		zz = 100
		xy = "合欢宗"
	)
	fmt.Printf("zz=%d,zz的类型是:%T\n", zz, zz)
	fmt.Printf("xy=%s,xy的类型是:%T\n", xy, xy)

	//定义常量
	//go语言用const关键字定义常量
	//常量可以当作枚举来使用
	const length = 10
	fmt.Println("length=", length)
	fmt.Println("beijing=", beijing)
	fmt.Println("shanghai=", shanghai)
	fmt.Println("shenzhen=", shenzhen)

	//iota常量计数器 默认值是从0开始计数  可以进行K倍数递增 K*iota
	fmt.Println("beijing1=", beijing1)
	fmt.Println("shanghai1=", shanghai1)
	fmt.Println("shenzhen1=", shenzhen1)

}
