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
	fmt.Printf("\n================================/ %s \n", runtime.FuncForPC(pc).Name())
}

func sp() {
	fmt.Printf("----\n")
}
func spt(title string) {
	fmt.Printf("---- %s \n", title)
}

/*******************************************************************/

func square(x int) int {
	return x * x
}

func divMod(x, y int) (int, int) {
	return x / y, x % y
}

func hello() {
	pf()
	fmt.Printf("square: %d\n", square(5))

	p, q := divMod(10, 3)
	fmt.Println(p)
	fmt.Println(q)
}

func foo1(a int, args ...int) {
	fmt.Println(a, args)
}

func foo0(args ...int) {
	fmt.Println(args)
}

func zoo(x int, y ...int) {
	fmt.Println(x, y)
}

func variable_length_arguments() {
	pf()

	foo0()
	foo0(1)
	foo0(1, 2)
	foo0(1, 2, 3)

	sp()

	foo1(1)
	foo1(1, 2)
	foo1(1, 2, 3)
	foo1(1, 2, 3, 4)

	spt("slice aegv")

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	zoo(0, a...)
	fmt.Println(append(a, b...))
}

func find(n int, ary []int) bool {
	for _, v := range ary {
		if n == v {
			return true
		}
	}
	return false
}

func position(n int, ary []int) int {
	for i, v := range ary {
		if n == v {
			return i
		}
	}
	return -1
}

func count(n int, ary []int) int {
	c := 0
	for _, v := range ary {
		if n == v {
			c++
		}
	}
	return c
}

func data_sarch() {
	pf()
	a := []int{1, 2, 3, 1, 2, 3, 4, 5}
	fmt.Println(find(4, a))
	fmt.Println(find(6, a))
	fmt.Println(position(5, a))
	fmt.Println(position(7, a))
	fmt.Println(count(3, a))
	fmt.Println(count(8, a))
}

func binarySearch(n int, ary []int) bool {
	low := 0
	high := len(ary) - 1
	for low <= high {
		mid := low + (high-low)/2
		if ary[mid] == n {
			return true
		} else if ary[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func binarySearch_test() {
	pf()
	a := []int{10, 20, 30, 40, 50, 60, 70, 80}
	fmt.Println(binarySearch(10, a))
	fmt.Println(binarySearch(40, a))
	fmt.Println(binarySearch(80, a))
	fmt.Println(binarySearch(0, a))
	fmt.Println(binarySearch(45, a))
	fmt.Println(binarySearch(90, a))

}

func insertSort(ary []int) {
	for i := 1; i < len(ary); i++ {
		tmp := ary[i]
		j := i - 1
		for ; j >= 0 && tmp < ary[j]; j-- {
			ary[j+1] = ary[j]
		}
		ary[j+1] = tmp
	}
}

func insertSort_test() {
	pf()
	a := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	b := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	c := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	insertSort(a)
	insertSort(b)
	insertSort(c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {
	hello()
	variable_length_arguments()
	data_sarch()
	binarySearch_test()
	insertSort_test()
}
