package main

import "errors"

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

func fact(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("fact : domain error")
	}
	a := 1
	for ; n > 1; n-- {
		a *= n
	}
	return a, nil
}

func error1_test() {
	pf()
	for x := 10; x >= -1; x-- {
		v, err := fact(x)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}

// ==================================================================

type MyError struct {
	msg string
}

func newMyError(s string) *MyError {
	err := new(MyError) //
	err.msg = s         //
	return err          //
}

func (err *MyError) Error() string {
	return err.msg
}

func fact2(n int) (int, error) {
	if n < 0 {
		return 0, newMyError("fact : domain error")
	}
	a := 1
	for ; n > 1; n-- {
		a *= n
	}
	return a, nil
}
func error2_test() {
	pf()
	for x := 10; x >= -1; x-- {
		v, err := fact2(x)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}

// ==================================================================
func fact3(n int) (int, error) {
	if n < 0 {
		panic("fact : domain error")
	}
	a := 1
	for ; n > 1; n-- {
		a *= n
	}
	return a, nil
}
func error3_test() {
	pf()
	for x := 10; x >= -1; x-- {
		v, err := fact3(x)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(v)
		}
	}
}

// ==================================================================
func baz() {
	panic("oops!")
}

func bar() {
	defer fmt.Println("bar end")
	fmt.Println("bar start!")
	baz()
}

func foo() {
	defer fmt.Println("foo end")
	fmt.Println("foo start!")
	bar()
}

func error4_test() {
	pf()
	foo()
}

// ==================================================================
func baz2() {
	panic("oops!")
}

func bar2() {
	defer fmt.Println("bar end")
	fmt.Println("bar start!")
	baz2()
}

func foo2() {
	defer func() {
		fmt.Println("foo end")
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("foo start!")
	bar2()
}

func error5_test() {
	pfm("recover")
	fmt.Println("main start!")
	foo2()
	fmt.Println("main end")
}

// ==================================================================
func foo3(n int) int {
	defer func() {
		fmt.Println("foo end")
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("foo start!")
	bar()
	return n * n
}

func error6_test() {
	pfm("recover")
	fmt.Println("main start!")
	fmt.Println(foo3(10))
	fmt.Println("main end")
}

// ==================================================================
type MyError2 struct {
	msg string
}

func newMyError2(s string) *MyError2 {
	err := new(MyError2)
	err.msg = s
	return err
}

func (err *MyError2) Error() string {
	return err.msg
}

func baz10() {
	panic(newMyError2("oops!"))
}

func baz20() {
	panic("oops!")
}

func barz(f func()) {
	defer fmt.Println("bar end")
	fmt.Println("bar start!")
	f()
}

func fooz(f func()) {
	defer func() {
		fmt.Println("foo end")
		err := recover()
		if err != nil {
			v := err.(*MyError)
			fmt.Println(v)
		}
	}()
	fmt.Println("foo start!")
	barz(f)
}

func error7_test() {
	pfm("error catch")
	fmt.Println("main start!")
	fooz(baz10)
	fooz(baz20)
	fmt.Println("main end")
}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	error1_test()
	error2_test()
	//error3_test()
	//error4_test()
	error5_test()
	error6_test()
	//error7_test()
}
