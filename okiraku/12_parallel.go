package main

import (
	"math"
	"math/rand"
	"time"
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
func fibo(n int) int {
	if n < 2 {
		return 1
	} else {
		return fibo(n-2) + fibo(n-1)
	}
}

func parallel1_test() {
	pf()
	runtime.GOMAXPROCS(4) //
	ch := make(chan int, 5)
	for _, n := range []int{41, 41, 39, 35, 36} {
		//for _, n := range []int{45, 44, 40, 35, 36} {
		go func(x int) {
			ch <- fibo(x)
		}(n)
	}
	for i := 5; i > 0; {
		select {
		case n := <-ch:
			fmt.Println(n)
			i--
		case <-time.After(time.Second):
			fmt.Println("Timeout")
			i = 0
		}
	}
}

// ==================================================================
func parallel2_test() {
	pf()
	ch := make(chan int, 6)
	for n := 1; n <= 6; n++ {
		fmt.Println("-----", n, "-----")
		runtime.GOMAXPROCS(n)
		s := time.Now()
		for i := 0; i < 4; i++ {
			go func() {
				ch <- fibo(40)
			}()
		}
		for i := 4; i > 0; i-- {
			fmt.Print(<-ch, " ")
		}
		e := time.Now().Sub(s)
		fmt.Println(e)
	}

}

// ==================================================================

func leftPoint(n int) float64 {
	w := 1.0 / float64(n)
	s := 0.0
	for i := 0; i < n; i++ {
		x := float64(i) * w
		s += 4.0 / (1.0 + x*x)
	}
	return s * w
}

func midPoint(n int) float64 {
	w := 1.0 / float64(n)
	s := 0.0
	for i := 1; i <= n; i++ {
		x := (float64(i) - 0.5) * w
		s += 4.0 / (1.0 + x*x)
	}
	return s * w
}

func parallel3_test() {
	pf()
	n := 10
	for i := 1; i <= 9; i++ {
		fmt.Println("-----", n, "-----")
		pi := midPoint(n)
		fmt.Println(pi, math.Pi-pi)
		pi = leftPoint(n)
		fmt.Println(pi, math.Pi-pi)
		n *= 10
	}
}

// ==================================================================
const (
	N = 1000000000
	W = 1.0 / float64(N)
)

func leftPoint2(n, m int) float64 {
	s := 0.0
	for i := n; i < m; i++ {
		x := float64(i) * W
		s += 4.0 / (1.0 + x*x)
	}
	return s * W
}

func parallel4_test() {
	pf()
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 1; i <= 8; i *= 2 {
		fmt.Println("-----", i, "-----")
		ch := make(chan float64, i)
		s := time.Now()
		k := N / i
		for j := 0; j < i; j++ {
			go func(n, m int) {
				ch <- leftPoint2(n, m)
			}(j*k, (j+1)*k)
		}
		sum := 0.0
		for j := i; j > 0; j-- {
			sum += <-ch
		}
		e := time.Now().Sub(s)
		fmt.Println(sum)
		fmt.Println(e)
	}
}

// ==================================================================
func montePi(n, s int) float64 {
	c := 0
	r := rand.New(rand.NewSource(int64(s)))
	for i := n; i > 0; i-- {
		x := r.Float64()
		y := r.Float64()
		if x*x+y*y < 1.0 {
			c++
		}
	}
	return (4.0 * float64(c)) / float64(n)
}

func parallel5_test() {
	pf()
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 1; i <= 8; i *= 2 {
		fmt.Println("-----", i, "-----")
		n := 100000000 / i
		ch := make(chan float64, i)
		s := time.Now()
		for j := 0; j < i; j++ {
			go func(x int) {
				ch <- montePi(n, x)
			}(j + 1)
		}
		pi := 0.0
		for j := i; j > 0; j-- {
			pi += <-ch
		}
		fmt.Println(pi / float64(i))
		e := time.Now().Sub(s)
		fmt.Println(e)
	}
}

// ==================================================================

type SortI interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

func quickSortSub(data SortI, low, high int) {
	p := low + (high-low)/2
	i, j := low, high
	for {
		for data.Less(i, p) {
			i++
		}
		for data.Less(p, j) {
			j--
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
		switch {
		case p == i:
			p = j
		case p == j:
			p = i
		}
		i++
		j--
	}
	if low < i-1 {
		quickSortSub(data, low, i-1)
	}
	if high > j+1 {
		quickSortSub(data, j+1, high)
	}
}

func quickSortParaSub(data SortI, low, high int) {
	if high-low < 1024 {
		quickSortSub(data, low, high)
		return
	}
	p := low + (high-low)/2
	i, j := low, high
	for {
		for data.Less(i, p) {
			i++
		}
		for data.Less(p, j) {
			j--
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
		switch {
		case p == i:
			p = j
		case p == j:
			p = i
		}
		i++
		j--
	}
	ch := make(chan int, 2)
	go func() {
		quickSortParaSub(data, low, i-1)
		ch <- 0
	}()
	go func() {
		quickSortParaSub(data, j+1, high)
		ch <- 0
	}()
	<-ch
	<-ch
}

func quickSortPara(data SortI) {
	quickSortParaSub(data, 0, data.Len()-1)
}

type IntArray []int

func (ary IntArray) Len() int {
	return len(ary)
}

func (ary IntArray) Less(i, j int) bool {
	return ary[i] < ary[j]
}

func (ary IntArray) Swap(i, j int) {
	ary[i], ary[j] = ary[j], ary[i]
}

func parallel6_test() {
	pf()
	//max := 10
	max := 1000000
	//runtime.GOMAXPROCS(runtime.NumCPU())
	b := make(IntArray, max)
	for i := 0; i < max; i++ {
		x := rand.Int()
		b[i] = x
	}
	c := 0
	for i, v := range b {
		fmt.Println(i, v)
		c++
		if c > 10 { break}
	}

	quickSortPara(b)

	fmt.Println("")
	c = 0
	for i, v := range b {
		fmt.Println(i, v)
		c++
		if c > 10 { break}
	}
}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	parallel1_test()
	parallel2_test()
	parallel3_test()
	parallel4_test()
	parallel5_test()
	parallel6_test()
}
