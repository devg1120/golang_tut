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

func pointer_test() {
	pf()
	var n int = 10
	var m int = 20
	var p *int = &n
	var q *int = &m
	fmt.Println(n)
	fmt.Println(*p)
	*p = 100
	fmt.Println(n)
	fmt.Println(*p)

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(p == q)
	fmt.Println(p != q)

}

//==================================================================

func array_pointer_test() {
	pf()
	var a [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	var b [8]int = [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	var p *[8]int = &a
	p1, p2 := &a[0], &a[1]
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(p1)
	fmt.Println(*p1)
	fmt.Println(p[0])
	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Println(p[1])
	*p1 = 10
	*p2 = 20
	fmt.Println(a)
	p = &b
	fmt.Println(*p)
}

// ==================================================================
func string_slice_pointer_test() {
	pf()
	var s string = "hello, world"
	var p *string = &s
	// p1, p2 := &s[0], &s[1]
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var q *[]int = &a
	q1, q2 := &a[0], &a[1]
	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(*p)
	*p = "oops!"
	fmt.Println(s)
	fmt.Println(*p)
	fmt.Println(q)
	fmt.Println(*q)
	fmt.Println(*q1)
	fmt.Println(*q2)
	*q = []int{10, 20, 30, 40}
	fmt.Println(q)
	fmt.Println(*q)
}

// ==================================================================
// MISSINBG
func swap(x *int, y *int) { // func swap(x int, y int) {
	tmp := *x //     tmp := x
	*x = *y   //     x = y
	*y = tmp  //     y = tmp
} // }

func timesArray(n int, ary *[8]int) {
	for i := 0; i < len(*ary); i++ {
		ary[i] *= n
	}
}

func func_pointer_test() {
	pf()
	var a int = 10
	var b int = 20
	var c [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	swap(&a, &b)
	timesArray(10, &c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

// ==================================================================
func multi_stage_ponter_test() {
	pf()
	var i int = 100
	var p *int
	var q **int
	p = &i
	q = &p
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(q)
	fmt.Println(*q)
	fmt.Println(**q)

}

func dynamic_memalloc_pointer_test() {
	pf()
	var p *int = new(int)
	var q *float64 = new(float64)
	var a *[8]int = new([8]int)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(q)
	fmt.Println(*q)
	fmt.Println(a)
	fmt.Println(*a)
	*p = 100
	*q = 1.2345
	a[0] = 10
	a[7] = 80
	fmt.Println(*p)
	fmt.Println(*q)
	fmt.Println(*a)
}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	pointer_test()
	array_pointer_test()
	string_slice_pointer_test()
	func_pointer_test()
	multi_stage_ponter_test()
	dynamic_memalloc_pointer_test()
}
