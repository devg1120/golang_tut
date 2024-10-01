package main

/*******************************************************************/
import "fmt"
import "os"
import "runtime"
import "path/filepath"

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	demo1_test()
	demo2_test()
}
// ==================================================================
func demo1_test() {
	pf()
	sp()
}
func demo2_test() {
	pfm("ABC")
	spt("TITLE")
}

// ==================================================================
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
	fmt.Printf("---------------------\n")
}
func spt(title string) {
	fmt.Printf("--------------------- %s \n", title)
}

/*******************************************************************/

