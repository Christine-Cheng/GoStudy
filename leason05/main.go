package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	//一般用int表示整形的宽度,再32位系统中就是32位,再64位系统中就是64位
	integer()
	fmt.Println("-------------------------------------------------")
	unsignedInteger()
	fmt.Println("-------------------------------------------------")
	showFloat()
}

//
func integer() {
	var num8 int8 = 127
	var num16 int16 = 32767
	var num32 int32 = math.MaxInt32
	var num64 int64 = math.MaxInt64
	fmt.Println("有符号整形")
	fmt.Printf("num8的类型是:%T, num8的大小是:%d, num8是:%d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是:%T, num16的大小是:%d, num16是:%d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是:%T, num32的大小是:%d, num32是:%d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是:%T, num64的大小是:%d, num64是:%d\n", num64, unsafe.Sizeof(num64), num64)
}

//无符号整型
func unsignedInteger() {
	var num8 uint8 = 128
	var num16 uint16 = 32768
	var num32 uint32 = math.MaxUint32
	var num64 uint64 = math.MaxUint64
	var num uint = math.MaxUint
	fmt.Println("无符号整型")
	fmt.Printf("num8的类型是 %T, num8的大小 %d, num8是 %d\n", num8, unsafe.Sizeof(num8), num8)
	fmt.Printf("num16的类型是 %T, num16的大小 %d, num16是 %d\n", num16, unsafe.Sizeof(num16), num16)
	fmt.Printf("num32的类型是 %T, num32的大小 %d, num32是 %d\n", num32, unsafe.Sizeof(num32), num32)
	fmt.Printf("num64的类型是 %T, num64的大小 %d, num64是 %d\n", num64, unsafe.Sizeof(num64), num64)
	fmt.Printf("num的类型是 %T, num的大小 %d, num是 %d\n", num, unsafe.Sizeof(num), num)
}

//浮点型
func showFloat() {
	var num1 float32 = math.MaxFloat32
	var num2 float64 = math.MaxFloat64
	fmt.Println("浮点型")
	fmt.Printf("num1的类型是 %T, num1的大小 %d, num1是 %d\n", num1, unsafe.Sizeof(num1), num1)
	fmt.Printf("num2的类型是 %T, num2的大小 %d, num2是 %d\n", num2, unsafe.Sizeof(num2), num2)

}
