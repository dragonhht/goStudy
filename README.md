# Go学习

## 1、包定义

-   编写Go语言的源文件时，第一行必须使用关键字`package`定义包名

-   每个Go程序必须有一个名为`main`的包，只有在包`main`下，`main`方法才正常可运行

-   同一路径下只存在一个`package`,一个`package`可由多个源文件组成

## 2、导入

-   使用关键字`import`导入源文件所依赖的package包

-   `import`的两种语法格式

    -   单个包逐个导入

    ```go
    package main
    import "fmt"
    import "time"
    ```

    -   导入多个包

    ```go
    package main
    import (
        "fmt"
        "time"
    )
    ```

-   当导入一个未使用的包时，运行程序将会报错

-   如果一个包被导入了多次，则该包只会被导入一次

-   一个main包中导入其他的包，这些导入的包将会被顺序导入

-   如果导入的包中存在依赖关系，如导入的包依赖与包A,则会先导入A包，然后初始化A包中的常量和变量，最后如果A包中存在`init`,则会自动执行init()

-   当包都导入完成后，才会对main中的`常量`和`变量`进行初始化,然后执行main中的`init()`函数,最后执行`main()`函数

-   `import`时使用`别名`、`.`及`_`的意义

    -   使用别名&nbsp;&nbsp;&nbsp;(使用包中函数时，通过别名调用)

    ```go
    package main
    import (
        user "test1"
    )

    func main() {
        user.Test()
    }
    ```

    -   导入时使用`.`&nbsp;&nbsp;&nbsp;(使用包中函数时，省略包名，直接通过函数名使用，不推荐使用，容易造成冲突)

    ```go
    package main
    import (
        "fmt"
        . "time"
    )

    func main() {
        fmt.Println(Time{})
    }
    ```

    -   导入时使用`_`&nbsp;&nbsp;&nbsp;(导入该包，但只执行`init()`方法，不能调用其他函数)

## 3、数据类型

-   `bool`： 布尔型，值为`true`、`false`,如：`var ok bool = true`

-   整数类型: 

    -   `uint8` 或 `byte`: 无符号8位整型(0 ~ 2^8 -1)

    -   `uint16`: 无符号16位整型(0 ~ 2^16 -1)

    -   `uint32`: 无符号32位整型(0 ~ 2^32 -1)

    -   `uint64`: 无符号64位整型(0 ~ 2^64 -1)

    -   `uint`: `uint32`或`uint64`（根据系统而定32位系统则为`uint32`）

    -   `int8`: 有符号8位整形(-2^7 ~ 2^7 -1)

    -   `int16`: 有符号16位整形(-2^15 ~ 2^15 -1)

    -   `int32` 或 `rune`: 有符号32位整形(-2^31 ~ 2^31 -1)

    -   `int64`: 有符号64位整形(-2^63 ~ 2^63 -1)

    -   `int`: `int32`或`int64`（根据系统而定32位系统则为`int32`）

    -   `uintptr`： 无符号整型,用于存放一个指针

-   浮点型:

    -   `float32`: 32位浮点型

    -   `float64`: 64位浮点型

    -   `complex64`： 32位实数和虚数

    -   `complex128`： 64位实数和虚数

-   `string`: 字符串类型，b编码统一为`UTF-8`

-   派生类型

## 4、变量与常量

### 1、变量

-   定义： 使用关键字`var`定义变量(全局变量必须使用`var`关键字，局部变量则可省略),如: `var count int = 32`，定义局部变量时可省略`var`关键字,但需要使用`：=`进行赋值,如`n1, n2 := 2, 3`

-   同一行申明统一数据类型时各变量名用`,`隔开，如： `var a, b, c, d int = 1, 2, 3, 4`

-   分组申明格式:

```go
package main
var (
	s1 int
	s2 uint
	s3 string
	s4 bool
)
```

-   Go中不存在类型的隐式转换，所有的转换必须显示，且转换格式为`变量 [:]= 转换类型(需转换的值)`， 如： `var z = int64(a)`，将`int`类型的变量a转换为`int64`

-   以`大写字母`开头命名的变量为可导出的变量(公用变量)，其他包可以调用；以`小写字母`开头命令的变量为不可导出的，是私有变量

-   特殊变量`_`

### 2、常量

-   Go中的常量使用关键字`const`申明，如：`const testCount, d1, d3  = 33, "Hello", true`或`const d4 int = 32`,常量的申明也可使用分组申明,当使用分组申明时若为对常量赋值，则该常量将自动继承前一常量的非空表达式

-   特殊常量`iota`:
    
    -   iota在const关键字出现时将会被重置为0
    
    -   const中每新增一行常量申明将使iota计数一次(分组申明)
    
## 5、控制语句

### 1、条件语句

#### 1、`if-else`:

-   `if-else`的嵌套使用:

```cgo
func ifTest() {
	a := 2
	if a > 0 {
		fmt.Println("a > 0")
	} else {
		if a == 0 {
			fmt.Println("a == 0")
		} else {
			fmt.Println(" a <  0")
		}
	}
}
```

#### 2、`switch`:

-   `switch`语句的使用:

```cgo
func switchTest(num int) {
	switch num {
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	default:
		fmt.Println("this is not 1 or 2")
	}
}
```

#### 3、`select`：

### 2、循环语句

-   `for`的无限循环：当for语句为指明条件时，则默认为`for true {}`,如：

```cgo
/**
	无限循环.
 */
func forTest1() {
	count := 1
	for {
		count++
		fmt.Println(count)
		Sleep(1 * Second)
	}
}
```

-   设置条件的`for`循环:

```cgo
func forTest2() {
	for i := 0; i < 5; i++{
		fmt.Println(i)
	}
}
```

-   模拟`foreach`(若不需要`index`可使用特殊变量`_`)

```cgo
func forTest3() {
	arr := []string{"a", "b", "c", "d", "e"}
	for key, value := range arr {
		fmt.Print(key)
	}
}
```

-   类似于`while`语句的写法:

```cgo
func forTest4() {
	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}
}
```

### 3、`goto`:

```cgo
func gotoTest() {
	goto One
	fmt.Println("this after One")
	One:
		fmt.Println("after")
}
```

### 4、`break`与`continue`与其他语言一致

## 6、函数

-   函数的定义：

    -   定义一个传入两个参数，并返回一个值的函数
    
    ```cgo
    func fun3(n1, n2 int) int {
        return n1 + n2
    }
    ```
    
    -   定义一个传入两个参数，并返回两个值的函数
    
    ```cgo
    func fun4(str1, str2 string) (string, string) {
    	return str1, str2
    }
    ```

-   在默认情况下，Go语言使用的都是值传递(数组等类型除外)，若需要使用引用传递，需使用指针，如下:

```cgo
func fun2(n *int) {
	*n = 9
}
```

-   将函数作为值:

```cgo
fun6 := func(n1, n2 int) int {
    return n1 + n2
}

func main() {
    fmt.Println(fun6(3, 4))
}
```

-   闭包

```cgo
func fun7(n1, n2 int) func(n3, n4 int) int {
	i := n1
	return func(n3, n4 int) int {
		return i + n2 + n3 + n4
	}
}

func main() {
    fun7 := fun7(2, 3)
    fmt.Println(fun7(3, 4))
    fmt.Println(fun7(5, 6))
}
```

-   `方法`:一个包含了接受者的函数

```cgo
/*
	定义结构体.
 */
type type1 struct {
	name string
}

/*
	定义属于type1类型的方法.
 */
func (t type1) getName() string {
	return t.name
}

func main() {
    // 调用方法
    var t type1
    t.name = "dragonhht"
    fmt.Println(t.getName())
}
```
