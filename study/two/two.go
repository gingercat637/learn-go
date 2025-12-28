package two

import (
	"fmt"
	"math"
)

func Run() {
	var testZeroBool bool
	fmt.Println("bool-zero:", testZeroBool)

	var testInt8 int8
	fmt.Println("int8-zero:", testInt8)

	var testInt16 int16
	fmt.Println("int16-zero:", testInt16)

	var testInt32 int32
	fmt.Println("int32-zero:", testInt32)

	var testInt64 int64
	fmt.Println("int64-zero:", testInt64)

	testInt8 = 127
	fmt.Println("int8-max:", testInt8)
	// 取以2为底128的对数
	fmt.Println("int8-log2:", math.Log2(float64(testInt8)))

	testInt8 = -128
	fmt.Println("int8-min:", testInt8)

	var testUint8 uint8
	fmt.Println("uint8-zero:", testUint8)
	testUint8 = math.MaxUint8
	fmt.Println("uint8-log2:", math.Log2(float64(testUint8)))

	fmt.Println("int-max", math.MaxInt)
	fmt.Println("int-log2", math.Log2(float64(math.MaxInt)))
	// fmt.Println("uint-max", math.MaxUint)
	fmt.Println("uint-log2", math.Log2(float64(math.MaxUint)))

	var byteVal byte = 127
	fmt.Println("byte-max:", byteVal)
	testInt16 = 127
	fmt.Println("byte和int16是否相等(需要强转)：", byteVal == uint8(testInt16))

	// 浮点数判断
	f1 := 0.1
	f2 := 0.2
	fmt.Println("f1+f2是否等于0.3:", f1+f2 == 0.3)

	// 字符串
	s1 := "你好世界"
	fmt.Println("s1[0]:", s1[0])
	fmt.Println("s1:", s1)
	fmt.Println("string(s1[0]):", string(s1[0]))

	b1 := []byte(s1)
	fmt.Println("b1:", b1)
	fmt.Println("b1[0]:", b1[0])
	fmt.Println("string(b1[0]):", string(b1[0]))

	r1 := []rune(s1)
	fmt.Println("r1:", r1)
	fmt.Println("r1[0]:", r1[0])
	fmt.Println("string(r1[0]):", string(r1[0]))

	// 循环字符串
	for _, ch := range s1 {
		fmt.Println("ch:", string(ch))
	}

	// 定义数组
	var a1 = [3]int{1,2,3}
	fmt.Println("a1:", a1)

	a2 := [3]rune{'a', 'b', 'c'}
	fmt.Println("a2:", a2)

	// a3 := [3]string{'a', 'b', 'c'}
	// fmt.Println("a3:", a3)

}
