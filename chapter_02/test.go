package main

import (
	"fmt"
	"reflect"
)

func LoopTest() {
	// 不使用标记
	fmt.Println("---- common ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
		}
	}

	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	fmt.Println("---- break label case 1 ----")

	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
	re:
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}

	fmt.Println("---- break label case 2 ----")

re1:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)

		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re1
		}
	}
}

// 指定类型，但是不指定参数个数 采用语法糖
func FuncManyParamsCase1(args ...int) {
	for _, arg := range args {
		fmt.Println("Arg1 is", arg)
	}
}

// 指定类型，但是不指定参数个数 采用原始方法
func FuncManyParamsCase2(args []int) {
	for _, arg := range args {
		fmt.Println("Arg2 is", arg)
	}
}

// 不指定参数个数，也不指定参数的类型
func FuncManyParamsCase3(args ...interface{}) {
	for _, arg := range args {
		fmt.Printf("Arg3 is %v, type is %v\n", arg, reflect.TypeOf(arg))
	}
}

func main() {
	fmt.Println("Test")
	// LoopTest()
	fmt.Println("--------------FunManyParamsCase1Start-------------")
	FuncManyParamsCase1(1, 2, 3, 4, 5, 6)
	fmt.Println("--------------FunManyParamsCase1End-------------")

	fmt.Println("--------------FunManyParamsCase2Start-------------")
	FuncManyParamsCase2([]int{1, 2, 3, 4, 5, 6})
	fmt.Println("--------------FunManyParamsCase2End-------------")

	fmt.Println("--------------FunManyParamsCase3Start-------------")
	FuncManyParamsCase3(1, 2, 3, 4, '4', '8', "Test")
	fmt.Println("--------------FunManyParamsCase3End-------------")

}
