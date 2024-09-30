package main

import "math"
import "strconv"
import (
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

type Point interface {
	distance0() float64
}

type Point2d struct {
	x, y float64
}

func newPoint2d(x, y float64) *Point2d {
	p := new(Point2d)
	p.x, p.y = x, y
	return p
}

func (p *Point2d) distance0() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type Point3d struct {
	x, y, z float64
}

func newPoint3d(x, y, z float64) *Point3d {
	p := new(Point3d)
	p.x, p.y, p.z = x, y, z
	return p
}

func (p *Point3d) distance0() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func sumOfDistance0(ary []Point) float64 {
	sum := 0.0
	for _, p := range ary {
		sum += p.distance0()
	}
	return sum
}

func interface1_test() {
	pf()
	a := []Point{
		newPoint2d(0, 0), newPoint2d(10, 10),
		newPoint3d(0, 0, 0), newPoint3d(10, 10, 10),
	}
	fmt.Println(a[0].distance0())
	fmt.Println(a[1].distance0())
	fmt.Println(a[2].distance0())
	fmt.Println(a[3].distance0())
	fmt.Println(sumOfDistance0(a))
}

//==================================================================

type Figure interface {
	kindOf() string
	area() float64
	print()
}
type Triangle struct {
	altitude, base float64
}

func newTriangle(a, b float64) *Triangle {
	p := new(Triangle)
	p.altitude, p.base = a, b
	return p
}

func (_ *Triangle) kindOf() string { return "Triangle" }

func (p *Triangle) area() float64 {
	return p.altitude * p.base / 2
}

func (p *Triangle) print() {
	fmt.Println("Triangle : area =", p.area())
}

// 四角形
type Rectangle struct {
	width, height float64
}

func newRectangle(w, h float64) *Rectangle {
	p := new(Rectangle)
	p.width, p.height = w, h
	return p
}

func (_ *Rectangle) kindOf() string { return "Rectangle" }

func (p *Rectangle) area() float64 {
	return p.width * p.height
}

func (p *Rectangle) print() {
	fmt.Println("Rectangle: area =", p.area())
}

type Circle struct {
	radius float64
}

func newCircle(r float64) *Circle {
	p := new(Circle)
	p.radius = r
	return p
}

func (_ *Circle) kindOf() string { return "Circle" }

func (p *Circle) area() float64 {
	return p.radius * p.radius * math.Pi
}

func (p *Circle) print() {
	fmt.Println("Circle: area =", p.area())
}

func sumOfArea(a []Figure) float64 {
	sum := 0.0
	for _, fig := range a {
		sum += fig.area()
	}
	return sum
}

func interface2_test() {
	pf()
	var a Figure = newTriangle(10, 10)
	fmt.Println(a.kindOf())
	fmt.Println(a.area())
	a.print()
	a = newRectangle(10, 10)
	fmt.Println(a.kindOf())
	fmt.Println(a.area())
	a.print()
	a = newCircle(10)
	fmt.Println(a.kindOf())
	fmt.Println(a.area())
	a.print()

	var b []Figure = []Figure{
		newTriangle(100, 100), newRectangle(100, 100), newCircle(100),
	}
	fmt.Println(sumOfArea(b))
}

// ==================================================================
func sumOfInt(ary []interface{}) int {
	sum := 0
	for _, x := range ary {
		v, ok := x.(int)
		if ok {
			sum += v
		}
	}
	return sum
}

func sumOfFloat(ary []interface{}) float64 {
	sum := 0.0
	for _, x := range ary {
		v, ok := x.(float64)
		if ok {
			sum += v
		}
	}
	return sum
}

func interface3_test() {
	pf()
	a := []interface{}{1, 1.1, "abc", 2, 2.2, "def", 3, 3.3}
	fmt.Println(sumOfInt(a))
	fmt.Println(sumOfFloat(a))
}

// ==================================================================
type Num interface {
	number()
}

type Int int

func (n Int) number() {}

type Real float64

func (n Real) number() {}

func sumOfNum(ary []Num) (Int, Real) {
	var sumi Int = 0
	var sumr Real = 0.0
	for _, x := range ary {
		switch v := x.(type) {
		case Int:
			sumi += v
		case Real:
			sumr += v
		}
	}
	return sumi, sumr
}

func interface4_test() {
	pf()
	var ary []Num = []Num{
		Int(1), Real(1.1), Int(2), Real(2.2), Int(3), Real(3.3),
	}
	a, b := sumOfNum(ary)
	fmt.Println(a, b)

}

// ==================================================================
type myInt int

func (n myInt) String() string {
	return "myInt;" + strconv.Itoa(int(n))
}

func interface5_test(memo string) {
	pfm(memo)
	n := myInt(10)
	m := myInt(20)
	fmt.Println(n + m)
	fmt.Println(n * m)
}

// ==================================================================
// Foo
type Foo struct {
	a int
}

type FooI interface {
	getA() int
}

func (p *Foo) getA() int { return p.a }

// Bar
type Bar struct {
	b int
}

type BarI interface {
	getB() int
}

func (p *Bar) getB() int { return p.b }

// Baz
type Baz struct {
	Foo
	Bar
}

type BazI interface {
	FooI
	BarI
}

func interface6_test(memo string) {
	pfm(memo)
	a := []FooI{
		&Foo{1}, &Foo{2}, &Baz{},
	}
	b := []BarI{
		&Bar{10}, &Bar{20}, &Baz{},
	}
	c := []BazI{
		&Baz{}, &Baz{Foo{1}, Bar{2}}, &Baz{Foo{3}, Bar{4}},
	}
	for i := 0; i < 3; i++ {
		fmt.Println(a[i].getA())
		fmt.Println(b[i].getB())
		fmt.Println(c[i].getA())
		fmt.Println(c[i].getB())
	}

}

// ==================================================================
func quickSortInt(buff []int, low, high int) {
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
		quickSortInt(buff, low, i-1)
	}
	if high > j+1 {
		quickSortInt(buff, j+1, high)
	}
}

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

func quickSort(data SortI) {
	quickSortSub(data, 0, data.Len()-1)
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

type CmpI interface {
	Less(CmpI) bool
}

func quickSortCmpI(buff []CmpI, low, high int) {
	p := buff[low+(high-low)/2]
	i, j := low, high
	for {
		for buff[i].Less(p) {
			i++
		}
		for p.Less(buff[j]) {
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
		quickSortCmpI(buff, low, i-1)
	}
	if high > j+1 {
		quickSortCmpI(buff, j+1, high)
	}
}

type Cint int

func (n Cint) Less(m CmpI) bool {
	return n < m.(Cint)
}

func quick_sort_test() {
	pf()
	a := make([]int, 1000000)
	b := make(IntArray, 1000000)
	c := make([]CmpI, 1000000)
	for i := 0; i < 1000000; i++ {
		x := rand.Int()
		a[i] = x
		b[i] = x
		c[i] = Cint(x)
	}
	s := time.Now()
	quickSortInt(a, 0, len(a)-1)
	e := time.Now().Sub(s)
	fmt.Println(e)
	s = time.Now()
	quickSort(b)
	e = time.Now().Sub(s)
	fmt.Println(e)
	s = time.Now()
	quickSortCmpI(c, 0, len(a)-1)
	e = time.Now().Sub(s)
	fmt.Println(e)

}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	interface1_test()
	interface2_test()
	interface3_test()
	interface4_test()
	interface5_test("Stringer")
	interface6_test("embed Interface")
	quick_sort_test()
}
