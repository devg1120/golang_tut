package main

import "fmt"
import "os"
import "runtime"
import "path/filepath"
import "regexp"
//import "time"
//import "reflect"

/*******************************************************************/

var farr = []string{
	"func_arg",
}
var fdic = map[string]FUNC_ENTRY{
	"func_arg":       func_arg_test,
}

/*******************************************************************/

func func1(x int) string {
	return fmt.Sprint("hello world '", x, "'.")
}

func func2(x, y int) (int, int) {
	return x + y, x - y
}


func func3(x, y int) (add, sub int) {
	add = x + y
	sub = x - y
	return
}

func func4(i ...int) int {
	fmt.Println("------")
	fmt.Println(i)
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}
func func_arg_test() {
        pf()
	fmt.Println(func1(1234))
	fmt.Println(func1(5678))
	sp()
	x, y := func2(10, 5)
	fmt.Println(x)
	fmt.Println(y)
	spt("named return arg")
	x, y = func3(10, 5)
	fmt.Println(x)
	fmt.Println(y)
	spt("variable length arg")
	fmt.Println(func4())
	fmt.Println(func4(1))
	fmt.Println(func4(1, 2))
	fmt.Println(func4(1, 2, 3))
	s := []int{1, 2, 3, 4}
	fmt.Println(func4(s...))
}
//--------------------------------------------------------------

func demo1_test() {
        pf()
        sp()
}
func demo2_test() {
        pfm("ABC")
        spt("TITLE")
}

func demo3_test() {
        pf()
        sp()
}


/*******************************************************************/

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

func pfm(memo string) {
	pc, _, _, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}
	fmt.Printf("\n================================/ %s : %s\n", runtime.FuncForPC(pc).Name(), memo)
}
func sp() {
	fmt.Printf("-----------------\n")
}
func spt(title string) {
	fmt.Printf("----------------- %s \n", title)
}

func p(a ...interface{}) {
	fmt.Println(a...)
}

type FUNC_ENTRY func()

func main() {

	fmt.Println("###", fn(), "###")
	if len(os.Args) < 2 {
		for i := range farr {
			fdic[farr[i]]()
		}
	} else {

		for _, keyword := range os.Args[1:] {
			//spt(keyword)
			re := regexp.MustCompile(keyword)
			for i := range farr {
				name := farr[i]
				if re.MatchString(name) {
					fdic[name]()
				}
			}
		}

	}

}

/*******************************************************************/
