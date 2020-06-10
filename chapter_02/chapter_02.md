## 循环
> Go语言中使用for进行循环，和c语言相比没有`while...do` 和`do...while`

### Case1： 普通循环
```go
// 普通循环
package main

import "fmt"

func main() {
    fmt.Println("---- common ----")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
        for i2 := 11; i2 <= 13; i2++ {
            fmt.Printf("i2: %d\n", i2)
        }
    }
}
```

**运算结果**
```
---- common ----
i: 1
i2: 11
i2: 12
i2: 13
i: 2
i2: 11
i2: 12
i2: 13
i: 3
i2: 11
i2: 12
i2: 13
```

### Case2： 带有break的循环
```go
// 带有break的虚幻
package main

import "fmt"

func main() {
    fmt.Println("---- break ----")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
        for i2 := 11; i2 <= 13; i2++ {
            fmt.Printf("i2: %d\n", i2)
            break
        }
    }
}
```
**运算结果**
```
---- break ----
i: 1
i2: 11
i: 2
i2: 11
i: 3
i2: 11
```

### Case3： 带有break和标签的循环
```go
// 带有标签的循环
package main

import "fmt"

func main() {
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
```
**运算结果**
```
---- break label case 1 ----
i: 1
i2: 11
i: 2
i2: 11
i: 3
i2: 11
---- break label case 2 ----
i: 1
i2: 11
```