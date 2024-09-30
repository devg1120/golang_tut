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

func fact(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func fact_test() {
	pf()
	for i := 0; i < 13; i++ {
		fmt.Println(i, ":", fact(i))
	}
}

func facti(n, a int) int {
	if n == 0 {
		return a
	} else {
		return facti(n-1, a*n)
	}
}

func facti_test() {
	pf()
	for i := 0; i < 13; i++ {
		fmt.Println(i, ":", facti(i, 1))
	}
}

func quicksort(low, high int, buff []int) {
	pivot := buff[low+(high-low)/2]
	i, j := low, high
	for {
		for pivot > buff[i] {
			i++
		}
		for pivot < buff[j] {
			j--
		}
		if i >= j {
			break
		}
		buff[i], buff[j] = buff[j], buff[i]
		i++
		j--
	}
	if low < i-1 {
		quicksort(low, i-1, buff)
	}
	if high > j+1 {
		quicksort(j+1, high, buff)
	}
}

func quicksort_test() {
	pf()
	a := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	fmt.Println(a)
	quicksort(0, 9, a)
	fmt.Println(a)
}

func main() {
	fmt.Println("###", fn(), "###")
	fact_test()
	facti_test()
	quicksort_test()

}
