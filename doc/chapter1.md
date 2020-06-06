# 第一章节
## Test1： 用一行语句打印一个表格
```go
// 需求 用一行语句打印表格
package main
import "fmt"

// 主函数
func main0() {
	fmt.Println("姓名\t年龄\t籍贯\t住址\r\njohn\t12\t河北\t北京")
}
```
**执行结果如下:**
```
姓名    年龄    籍贯    住址
john    12      河北    北京
```
## Tip1：注释的使用\<comment\>
> 行注释和块注释
* 行注释采用`// 这里是注释`, 可以参见下面的示例
* 块注释采用`/* 这里面是爪式 */`
**注释样例**
```go
// 这里是行注释
package main
import "fmt"

func main() {
	/* fmt.Print("Hello, world.") // 这里默认字符串后面没有\n
	fmt.Println("Hello, world.") // 这里字符串末尾默认有个\n
	fmt.Printf("Hello, %s", "jason") // 可见Printf可以按照格式进行格式化
	 */
	fmt.Println("Hello, world!")
}
```


## 问题
### Q: 常见的print有什么区别
> go语言中有多个Print，分别是Print，Printf以Println, 我们在调试代码的时候常用的是Println
### A: 区别如下
| 类型 | 一句话区别 | 备注 |
| ------ | ----- | ------| 
|  PrintIn | 每句都会换行 | 调试代码常用 |
| Printf | 可以进行格式化输出 | 无 |
| Print  | 每行代码默认不会换行| 无|

**代码样例**
```go
// 需求 用一行语句打印表格
package main
import "fmt"

func main() {
	fmt.Print("Hello, world.") // 这里默认字符串后面没有\n
	fmt.Println("Hello, world.") // 这里字符串末尾默认有个\n
	fmt.Printf("Hello, %s", "jason") // 可见Printf可以按照格式进行格式化
} 
```