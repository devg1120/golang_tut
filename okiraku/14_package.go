package main

import (
    "foo"
    "bar"
)


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

func pfm(memo string) {
	pc, _, _, success := runtime.Caller(1)

	if !success {
		println("functionName: runtime.Caller: failed")
		os.Exit(1)
	}
	fmt.Printf("\n================================/ %s : %s\n", runtime.FuncForPC(pc).Name(), memo)
}
func sp() {
	fmt.Printf("----\n")
}
func spt(title string) {
	fmt.Printf("---- %s \n", title)
}

/*******************************************************************/
/*
 # tree sample

      sample/
      |-- bar
      |   |-- bar.go
      |   `-- go.mod
      `-- foo
          |-- foo.go
          `-- go.mod
*/
/*
 # cat ./go.mod
 replace foo => ./sample/foo
 replace bar => ./sample/bar


 # go get foo
 # go get bar

*/

func package1_test() {
    pf()
    foo.Test()
    fmt.Println("foo.A =", foo.A)
    bar.Test()
    fmt.Println("bar.A =", bar.A)
}
// ==================================================================
// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
        package1_test() 
}
