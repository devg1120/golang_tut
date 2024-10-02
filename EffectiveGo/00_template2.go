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
	"demo1",
	"demo2",
	"demo3",
}
var fdic = map[string]FUNC_ENTRY{
	"demo1":        demo1_test,
	"demo2":        demo2_test,
	"demo3":        demo3_test,
}

/*******************************************************************/
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
