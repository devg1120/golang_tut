package main

import "math"

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

//==================================================================

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	interface1_test()
	interface2_test()
}
