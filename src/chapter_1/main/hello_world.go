// 需求 用一行语句打印表格
package main
import "fmt"

func main0() {
	fmt.Print("Hello, world.") // 这里默认字符串后面没有\n
	fmt.Println("Hello, world.") // 这里字符串末尾默认有个\n
	fmt.Printf("Hello, %s", "jason") // 可见Printf可以按照格式进行格式化
	fmt.Println("Hello, world!")
}

//主函数
func main1() {
	fmt.Println("姓名\t年龄\t籍贯\t住址")
	fmt.Println("john\t12\t河北\t北京")
}

// 主函数
func main2() {
	fmt.Println("我是jason, 我现在很好\r特殊")
}

// 常规函数1
func getCatName() (firstNme, secondName string){
	firstNme = "Hao"
	secondName = "Shuliang"
	return
}

// 常规函数2
func getDogName() (firstName string, secondName string){
	return "张三", "李四"
}


// 常规函数3
func getPigName() (firstName string, secondName string){
	m_1 := "网三"
	m_2 := "是四"
	return m_1, m_2
}

func main() {
	first, second := getCatName()
	fmt.Printf("The first name is %s\n", first)
	fmt.Printf("The second name is %s\n", second)
	fmt.Print("\n*********************************\n")
	// 前面是first和second的第一次赋值，所以必须采用的是 :=, 但是之后就直接使用= 就可以了
	first, second = getDogName()
	fmt.Printf("The first name is %s\n", first)
	fmt.Printf("The second name is %s\n", second)
	fmt.Print("\n*********************************\n")
	first, second = getPigName()
	fmt.Printf("The first name is %s\n", first)
	fmt.Printf("The second name is %s\n", second)
	fmt.Print("\n*********************************\n")
}