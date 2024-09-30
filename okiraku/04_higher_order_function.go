package main

/*******************************************************************/
import "fmt"
import "os"
import "runtime"
import "path/filepath"

func fn() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Base(file)
}

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
	fmt.Printf("\n================================/ %s \n", runtime.FuncForPC(pc).Name())
}

func sp() {
	fmt.Printf("----\n")
}
func spt(title string) {
	fmt.Printf("---- %s \n", title)
}

/*******************************************************************/
//
func square(x int) int {
	return x * x
}

func cube(x int) int {
	return x * x * x
}

func sumOf(f func(int) int, n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += f(n)
	}
	return a
}

func sumOf_test() {
	pf()
	fmt.Println(sumOf(square, 1, 100))
	fmt.Println(sumOf(cube, 1, 100))

	a := func(x int) int { return x * x }(10)
	fmt.Println(a)
	b := func(x int) int { return x * x }
	fmt.Println(b(10))
	fmt.Println(sumOf(func(x int) int { return x * x }, 1, 100))
}

//==================================================================

func mapcar(f func(int) int, ary []int) []int {
	buff := make([]int, len(ary))
	for i, v := range ary {
		buff[i] = f(v)
	}
	return buff
}
func removeIf(f func(int) bool, ary []int) []int {
	buff := make([]int, 0)
	for _, v := range ary {
		if !f(v) {
			buff = append(buff, v)
		}
	}
	return buff
}
func filter(f func(int) bool, ary []int) []int {
	buff := make([]int, 0)
	for _, v := range ary {
		if f(v) {
			buff = append(buff, v)
		}
	}
	return buff
}

func foldl(f func(int, int) int, a int, ary []int) int {
	for _, x := range ary {
		a = f(a, x)
	}
	return a
}

func foldr(f func(int, int) int, a int, ary []int) int {
	for i := len(ary) - 1; i >= 0; i-- {
		a = f(ary[i], a)
	}
	return a
}
func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 != 0
}

func add(x, y int) int {
	return x + y
}

func hight_order_test() {
	pf()
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := mapcar(square, a)
	c := mapcar(cube, a)
	fmt.Println(b)
	fmt.Println(c)
	d := removeIf(isEven, a)
	e := removeIf(isOdd, a)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(foldl(add, 0, a))
	fmt.Println(foldr(add, 0, a))
}

// ==================================================================
func foo(n int) func(int) int {
	return func(x int) int { return x * n }
}

func clouser_test() {
	pf()
	foo10 := foo(10)
	foo20 := foo(20)
	fmt.Println(foo10(1))
	fmt.Println(foo20(10))

}

// ==================================================================
func mapcar2(f func(x int) int) func([]int) []int {
	return func(ary []int) []int {
		buff := make([]int, len(ary))
		for i, v := range ary {
			buff[i] = f(v)
		}
		return buff
	}
}

func square2(x int) int {
	return x * x
}

func cube2(x int) int {
	return x * x * x
}

func clouser2_test() {
	pf()
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	squareAry := mapcar2(square2)
	cubeAry := mapcar2(cube2)
	fmt.Println(squareAry(a))
	fmt.Println(cubeAry(a))
	fmt.Println(mapcar2(square2)(a))
	fmt.Println(mapcar2(cube2)(a))
}

// ==================================================================

func makeGen() func() int {
	prevNumber := -1
	return func() int {
		prevNumber += 2
		return prevNumber
	}
}

func generator_test() {
	pf()
	g1 := makeGen()
	for i := 0; i < 8; i++ {
		fmt.Println(g1())
	}
	g2 := makeGen()
	for i := 0; i < 8; i++ {
		fmt.Println(g2())
	}
}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	sumOf_test()
	hight_order_test()
	clouser_test()
	clouser2_test()
	generator_test()
}
