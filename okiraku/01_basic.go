package main

import "fmt"
import "os"
import "runtime"

func gf() string {
	counter, _, _, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}
	return runtime.FuncForPC(counter).Name()
}

func pf() {
	pc, _, _, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}
	fmt.Printf("\n----------------------/ %s /---\n\n", runtime.FuncForPC(pc).Name())
}

func sp() {
        fmt.Printf("----\n")
}
//*****************************************************************************

func hello() {
	pf()
	fmt.Println("hello, world")
}

var a int = 10
var b int = 20
var c float64 = 1.234
var d float64 = 5.678

func sample01() {
	pf()
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(c * d)
	fmt.Println(c / d)
}

func sample01a() {
	pf()
	var (
		a, b = 30, 40
		c, d = 8.234, 9.678
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(c * d)
	fmt.Println(c / d)
}

func sample02() {
	pf()
	a, b := "abcd", "efgh" // var a, b string = "abcd", "efgh"
	c := a + b             // var c string = a + b
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func sample03() {
	pf()
	if i := 0; i == 0 {
		fmt.Println("zero")
	} else {
		fmt.Println("not zero")
	}
	i := 0
	for i < 10 {
		fmt.Println("hello, world")
		i += 1
	}
	i, sum := 1, 0
	for i <= 10000 {
		sum += i
		i++
	}
	fmt.Println(sum)
	for i := 0; i < 10; i++ {
		fmt.Println("hello, world")
	}
	for i := 1; i <= 100; i++ {
		switch i % 15 {
		case 0:
			fmt.Print("FizzBuzz")
		case 3, 6, 9, 12:
			fmt.Print("Fizz")
		case 5, 10:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println("")
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz")
		case i%3 == 0:
			fmt.Print("Fizz")
		case i%5 == 0:
			fmt.Print("Buzz")
		default:
			fmt.Print(i)
		}
		fmt.Print(" ")
	}
	fmt.Println("")
}

func sample04() {
	pf()
	var a [4]int
	b := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	// var b [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(a[0])
	fmt.Println(b[0])
	a[0] = 10
	b[0] = 20
	fmt.Println(a)
	fmt.Println(b)
}

func array_2d() {
	pf()
	// 2d array
	var a [3][3]int
	b := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	// var b [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(a))
	fmt.Println(len(b))
	fmt.Println(a[0][0])
	fmt.Println(b[2][2])
	a[0][0] = 10
	b[2][2] = 20
	fmt.Println(a)
	fmt.Println(b)
}

func slice() {
	pf()
	a := [4]int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8}
	c := a[:]
	d := b[2:6]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	c[0] = 10
	d[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	a2 := make([]int, 5, 10)
	fmt.Println(a2)
	fmt.Println(len(a2))
	fmt.Println(cap(a2))
	for i := 5; i <= 10; i++ {
		a2 = append(a2, i)
		fmt.Println(a2)
		fmt.Println(len(a2))
		fmt.Println(cap(a2))
	}

	var a3 = "1234567890"
	for i := 0; i <= len(a3); i++ {
		var s string = a3[i:]
		fmt.Println(s)
		fmt.Println(len(s))
	}
}

func map_test() {
	pf()
	var a map[string]int = map[string]int{"foo": 10, "bar": 20}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(a["foo"])
	a["foo"] = 100
	fmt.Println(a["foo"])
	a["baz"] = 30
	fmt.Println(a["baz"])
	fmt.Println(a)
	fmt.Println(len(a))

	b := make(map[string]int)
	b["foo"] = 10
	b["bar"] = 20
	b["baz"] = 30
	fmt.Println(b)
	delete(b, "baz")
	v, ok := b["baz"]
	fmt.Println(v)
	fmt.Println(ok)
	fmt.Println(b)

	c := []int{1, 2, 3, 4, 5}
	for i, v := range c {
		fmt.Println(i, v)
	}
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)

	d := map[string]int{"foo": 10, "bar": 20, "baz": 30}
	for k, v := range d {
		fmt.Println(k, v)
	}
	sum = 0
	for _, v := range d {
		sum += v
	}
	fmt.Println(sum)
}

func print_test() {
	pf()
    fmt.Printf("%d, %x, %o\n", 100, 100, 100)
    fmt.Printf("[%d]\n", 10)
    fmt.Printf("[%4d]\n", 10)
    fmt.Printf("[%4d]\n", 100000)
    fmt.Printf("[%4d]\n", 123456)
    fmt.Printf("[%-8d]\n", 123456)
    fmt.Printf("[%08d]\n", 123456)


    sp()
    type Foo struct {
        bar, baz int
    }
    a := 123456789
    b := Foo{10, 20}
    fmt.Printf("%v\n", a)
    fmt.Printf("%T\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%+v\n", b)
    fmt.Printf("%#v\n", b)
    fmt.Printf("%T\n", b)

    sp()
    z := "hello, world"
    fmt.Printf("[%s]\n", z)
    fmt.Printf("[%20s]\n", z)
    fmt.Printf("[%-20s]\n", z)
    fmt.Printf("[%q]\n", z)
    fmt.Printf("[%20q]\n", z)
    fmt.Printf("[%-20q]\n", z)

}
func main() {
	hello()
	sample01()
	sample01a()
	sample02() // string
	sample03() // if for switch
	sample04() // array
	array_2d() // 2d array
	slice()
	map_test()
	print_test()
}
