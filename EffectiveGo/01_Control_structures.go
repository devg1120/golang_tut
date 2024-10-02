package main

import "fmt"
import "os"
import "runtime"
import "path/filepath"
import "time"
import "reflect"
import "regexp"

/*******************************************************************/

var farr = []string{
	"if",
	"for",
	"switch",
	"defer",
	"defer2",
	"select",
	"dynamic_func",
}
var fdic = map[string]FUNC_ENTRY{
	"if":           if_test,
	"for":          for_test,
	"switch":       switch_test,
	"defer":        defer_test,
	"defer2":       defer2_test,
	"select":       select_test,
	"dynamic_func": dynamic_func_test,
}

func main_org() {
	fmt.Println("###", fn(), "###")

	if_test()
	for_test()
	switch_test()
	defer_test()
	defer2_test()
	select_test()
	dynamic_func_test()
}

/*******************************************************************/

func if_test() {
	pf()
	i := 5
	if i > 0 {
		fmt.Println("xxx")
	}
}
func for_test() {
	pf()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	arr := [5]int{1, 2, 3, 4, 5}

	sum := 0

	spt("array")
	for _, value := range arr {
		sum += value
	}
	fmt.Println(sum)

	spt("map")
	grades := map[string]int{"Alice": 85, "Bob": 90, "Charlie": 95}
	for name, grade := range grades {
		fmt.Println(name, grade)
	}
	spt("string")
	message := "Hello, world!"
	for i, char := range message {
		fmt.Println(i, string(char))
	}
	spt("slice")
	fruits := []string{"apple", "banana", "cherry"}
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
	spt("bleak label")

OuterLoop:
	for i := 0; i < 3; i++ {
		fmt.Printf("i: %d\n", i)

		for j := 0; j < 3; j++ {
			fmt.Printf("j: %d\n", j)
			if j == 2 {
				fmt.Println("inner loop")
				break OuterLoop
			}
		}

		if i == 2 {
			fmt.Println("outer loop")
			break
		}
	}
	fmt.Println("done")
}

func switch_test() {
	pf()
	i := 1

	spt("int")
	switch i {
	case 0:
		fmt.Println("i is 0")
	case 1:
		fmt.Println("i is 1")
	default:
		fmt.Println("x is neither 1 nor 2")
	}

	spt("string")
	s := "bbb"
	switch s {
	case "aaa":
		fmt.Println("case aaa")
	case "bbb":
		fmt.Println("case bbb")
	default:
		fmt.Println("case default")
	}

	var z interface{} = "hello"

	spt("type")
	switch v := z.(type) {
	case int:
		fmt.Printf("%v is int[%T]\n", v, v)
	case string:
		fmt.Printf("%v is string[%T]\n", v, v)
	default:
		fmt.Printf("default: %v[%T]\n", v, v)
	}
}

func defer_test() {
	pf()
	defer fmt.Println("world1")
	defer fmt.Println("world2")

	fmt.Println("hello")
}

/********************************************************/
type deferTest int

func (t *deferTest) func1() *deferTest {
	fmt.Println("func1")
	return t
}

func (t *deferTest) func2() *deferTest {
	fmt.Println("func2")
	return t
}

func (t *deferTest) func3() *deferTest {
	fmt.Println("func3")
	return t
}

func defer2_test() {
	pf()
	var dt deferTest
	defer dt.func1().func2().func3()

	fmt.Println("main func")
}

/********************************************************/

func test1(ch chan<- string) {
	for {
		ch <- "test1"
		time.Sleep(2 * time.Second)
	}
}

func test2(ch chan<- string) {
	for {
		ch <- "test2"
		time.Sleep(4 * time.Second)
	}
}

func test3(quit chan<- int) {
	time.Sleep(10 * time.Second)
	quit <- 0
}

func select_test() {
	pf()
	c1 := make(chan string)
	c2 := make(chan string)
	quit := make(chan int)
	go test1(c1)
	go test2(c2)
	go test3(quit)

	cnt := 0
	for {
		select {
		case s1 := <-c1:
			fmt.Println(s1)
		case s2 := <-c2:
			fmt.Println(s2)
		case <-quit:
			fmt.Println("quit")
			return
		default:
			cnt = cnt + 1
			fmt.Printf("(cnt: %v)\n", cnt)
			time.Sleep(1 * time.Second)
		}
	}
}

/*******************************************************************/
func function1() {
	fmt.Println("print function1")
}

func function2() {
	fmt.Println("print function2")
}

func function3() {
	fmt.Println("print function3")
}

func dynamic_func_test() {
	pf()
	funcList := []func(){function1, function2, function3}

	for i := 0; i < 3; i++ {
		//p(reflect.TypeOf(funcList[i]))
		fv := reflect.ValueOf(funcList[i])
		p(runtime.FuncForPC(fv.Pointer()).Name())
		funcList[i]()
	}
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
