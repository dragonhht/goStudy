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

## 7、数组(支持多维数组，使用同其他语言类似)

-   数组的申明:

    -   申明数组，并不进行初始化,`var 数组名 [长度] 类型`, 如:
    
    ```cgo
    var arr [5] int
    ```
    
    -   申明数组并进行初始化操作:
    
    ```cgo
    // 方式一, 指定数组大小
    var arr1 = [5]int{1, 2, 3, 4, 5}
    fmt.Println(arr1)
    // 方式二, 不手动设置数组大小，通过元素个数自动设置数组大小
    var arr2 = []int{1, 2, 4, 5}
    fmt.Println(arr2)
    ```
    
-   数组元素的获取同其他语言相同，通过索引获取

```cgo
var arr2 = []int{1, 2, 4, 5}
fmt.Println(arr2[2])
```

-   函数中将数组作为参数

```cgo
func arrTest(arr []int) {
	arr[0] = 3
}

func main() {
    arr := []int{1, 2, ,3}
    arrTest(arr)
    fmt.Println(arr)
}
```

## 8、指针

> 指针用于指向一个值的内存地址

-   Go语言获取一个值得内存地址同c类似,使用取地址符`&`获取变量的地址，如`count := 3;fmt.Println(&count)`

-   Go语言中指针的申明同c语言类似,使用`*`来表示某个变量为指针类型, 如:

```cgo
var count = 3
var p *int = &count
fmt.Println(p)
```

-   将指针类型作为函数形参

```cgo
func pointTest(n *int) {
	*n = 9
}

func main() {
	var count = 3
	pointTest(&count)
	fmt.Println(count)
}
```

-   空指针`nil`:

> Go中的空指针`nil`同Java中的`null`类似

```cgo
var p1 *int
fmt.Println(p1)
```

-   空指针的判断

```cgo
var p1 *int
if p1 == nil {
    fmt.Println("count is nil")
}
```

## 9、结构体

-   结构体的定义

    -   结构体定义格式

    ```cgo
    type 结构体名称 struct {
       member1 type1;
       member1 type2;
       ...
       membern typen;
    }
    ```
    
    -   例如:
    
    ```cgo
    type Student struct {
    	name string
    	age int8
    }
    ```
    
-   结构体的使用(结构体内成员的访问通过`.`号进行)

```cgo
func main() {
	stu := Student{"dragonhht", 18}
	stu.age = 20
	fmt.Println(stu)
}
```

-   将结构体作为函数参数时，在Go中并不会想Java中一样使用引用传递，而是值传递，需要使用类似于Java中的引用传递需使用指针实现

```cgo
/*
	值传递，不改变原有的值
 */
func setStudentName(student Student) {
	student.name = "hello"
}

/*
	通过指针，改变了原有的值.
 */
func setStudentName1(student *Student) {
	student.name = "hello"
}

func main() {
 	stu := Student{"dragonhht", 18}
 	stu.age = 20
 	fmt.Println(stu)
 	setStudentName(stu)
 	fmt.Println(stu)
 	setStudentName1(&stu)
 	fmt.Println(stu)
 }
```

## 10、切片

-   创建切片(使用`make`函数): `var slice1 []int = make([]int, 5)`

-   切片在为初始化前默认为`nil`，长度为`0`

-   切片的初始化

    -   通过初始化数组的形式初始化切片:`slice2 := [] int {1, 2, 3}`
    
    -   通过数组的引用来初始化切片: `arr := []int{1, 2, 3, 4, 5, 6, 7};slice3 := arr[:]; slice4:= arr[1:3]`
    
-   切片的使用

    -   `len()`函数,获取切片长度:`fmt.Println(len(slice3))`
    
    -   `cap()`函数，计算容量:`fmt.Println(cap(slice3))`
    
    -   `append()`函数，向切片中添加新元素:`slice3 = append(slice3, 55);fmt.Println(slice3)`
    
    -   `copy()`函数，拷贝切片:如下代码为将`slice3`的内容拷贝到`slice4`
    
    ```cgo
    func main() {
    	arr := []int{1, 2, 3, 4, 5, 6, 7}
    	slice3 := arr[:]
    	var slice4 = make([]int, 5, 6)
    	copy(slice4, slice3)
    	fmt.Println(slice4)
    }
    ```

## 11、`Map`集合

-   定义Map:

    -   通过`map`关键字定义:格式为`var 变量名 map[键的类型]值得类型`, 如`map2 := map[string]string{"name": "dragonhht", "age": "18"}`
    
    -   使用`make `函数创建Map: 格式为`make(map[键类型]值类型)`, 如`map2 := make(map[string]string)`,该方式会进行初始化
    
    -   定义的Map不进行初始化默认值为`nil`,不能用来存放键值
    
-   使用Map

    -   放入键值: `map1["hello"] = "world"`
    
    -   使用Map中某个键的值:`fmt.Println(map1["name"])`
    
    -   `delete`函数，用于删除Map中的元素，`delete(map1, "name")`
    
## 12、`interface`接口

-   接口的定义:

    -   接口的定义格式:
    
    ```cgo
    type 接口名 interface {
       method_name1 [return_type]
       method_name2 [return_type]
       method_name3 [return_type]
       ...
       method_namen [return_type]
    }
    ```

    -   接口定义示例
    
    ```cgo
    type interface1 interface {
    	getName() string
    	setName(name string)
    }
    ```
    
-   接口体实现接口示例:

```cgo
package main

import "fmt"

/*
	定义接口.
 */
type Interface1 interface {
	getName() string
	setName(name string)
}

/*
	定义结构体.
 */
type Struct1 struct {
	name string
}

/*
	实现接口.
 */
func (struct1 Struct1) getName() string  {
	return struct1.name
}

/*
	实现接口.
 */
func (struct1 *Struct1) setName(name string) {
	struct1.name = name
}


func main() {
	stu := Struct1{"dragonhht"}
	fmt.Println(stu)
	fmt.Println(stu.getName())
	stu.setName("huang")
	fmt.Println(stu)
	fmt.Println(stu.getName())
}
```

## 13、错误处理

> Go语言通过内置的错误接口提供了非常简单的错误处理机制, 在编码中我们可以通过实现`error`接口类型来生成错误信息  

-   内置的错误接口如下:

```cgo 
type error interface {
	Error() string
}
```

## 14、反射

```cgo
func info(inter interface{}) {
	// 通过t可获取定义中的所有元素
	t := reflect.TypeOf(inter)
	fmt.Println("Type is ", t.Name())
	// 通过field可获取存储的值，也可去改变他
	filed := reflect.ValueOf(inter)
	// 获取所有的属性
	for i := 0; i < t.NumField(); i++ {
		// 获取属性和相应的值
		f := t.Field(i)
		val := filed.Field(i).Interface()
		fmt.Printf("%s : %v = %v\n", f.Name, f.Type, val)
	}

	// 获取所有的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s : %v\n", m.Name, m.Type)
	}
}
```

## 15、并发

### 1、Goroutine

> 当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的goroutine会用`go`语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。go语句会使其语句中的函数在一个新创建的goroutine中运行

-   通过使用关键字`go`启动一个线程(此处因main迅速结束，所以无法输出任何内容，该情况需使用`Channel`)

```cgo
func run() {
	for i := 0; i < 10; i++ {
		fmt.Println("Go run...", i)

	}
}

func main() {
	go run()
}
```

-   Goroutine通过通信类共享内存，而不是通过共享内存来通信

####   `Channels`

-   一个channels是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息

-   创建`channel`(channel通过`make`函数创建)

> 一个基于无缓存Channels的发送操作将导致发送者goroutine阻塞，直到另一个goroutine在相同的Channels上执行接收操作，当发送的值通过Channels成功传输之后，两个goroutine可以继续执行后面的语句。反之，如果接收操作先发生，那么接收者goroutine也将阻塞，直到有另一个goroutine在相同的Channels上执行发送操作。

```cgo
// 创建无缓冲的channel
c := make(chan bool)
// 创建缓冲区为5的channel
c2 := make(chan int, 5)
```

-   channel有`发送`和`接收`两个主要操作

    -   发送：发送语句将一个值从一个goroutine通过channel发送到另一个执行接收操作的goroutine，发送操作使用运算法`<-`
    
    ```cgo
    // 发送, <- 运算符分割channel和要发送的值
    c := make(chan bool)
    c <- true
    ```
    
    -   接收：接收通过channel发送的值，同发送一样，使用运算符`<-`
    
    ```cgo
    // <- 运算符写在channel对象之前
    <- c
    // 也可用变量接收发送的值
    // x := <- c
    ```
    
-   `close`操作(用于关闭channel)

> 关闭channel后，对该channel进行任何发送操作，则都会导致`panic`异常，但依旧可以进行接收操作，接收以前成功发送的数据；如果channel中已经没有数据的话讲产生一个零值的数据。

    -   close操作
    
    ```cgo
    close(c)
    ```

-   基于无缓存Channels的发送和接收操作将导致两个goroutine做一次`同步操作`。

-   异步操作

    -   通过带缓存的channel：
    
    ```cgo
    func main() {
    	// 调用CPU核数
    	runtime.GOMAXPROCS(runtime.NumCPU())
    	c := make(chan bool, 10)
    	for i := 0; i < 10; i++ {
    		go func() {
    			runCount++
    			fmt.Println(runCount)
    			c <- true
    		}()
    	}
    
    	for i := 0; i < 10; i++ {
    		<- c
    	}
    }
    ```
    
    -   通过`sync.WaitGroup`
    
    ```cgo
    func run(w *sync.WaitGroup) {
    	runCount++
    	fmt.Println(runCount)
    	// 标记任务完成，减少WaitGroup
    	w.Done()
    }
    
    func main() {
    	// 调用CPU核数
    	runtime.GOMAXPROCS(runtime.NumCPU())
    	//c := make(chan bool, 10)
    	// 创建WaitGroup
    	wg := sync.WaitGroup{}
    	// 添加10个WaitGroup
    	wg.Add(10)
    	for i := 0; i < 10; i++ {
    		go run(&wg)
    	}
    	//等待完成
    	wg.Wait()
    }
    ```
    
-   `Select`

> 可处理一个或多个channel的发送与接收  
> 同时有多个可用的channel时按随机顺序处理  
> 可用空的select来阻塞main函数  
> 可设置超时

