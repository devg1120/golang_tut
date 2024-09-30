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

type Point struct {
	x float64 // x, y float64
	y float64
}

func distance(p, q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func struct1_test() {
	pf()
	var p Point                         // p := Point{}
	var q Point = Point{10, 10}         // q := Point{10, 10}
	var r Point = Point{x: 100, y: 100} // r := Point{x:100, y:100}
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(p.x)
	fmt.Println(p.y)
	fmt.Println(q.x)
	fmt.Println(q.y)
	fmt.Println(r.x)
	fmt.Println(r.y)
	fmt.Println(distance(p, q))
	fmt.Println(distance(p, r))
	fmt.Println(distance(q, r))

}

// ==================================================================
type Point2 struct {
	x float64 // x, y float64
	y float64
}

func distance2(p, q *Point2) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}
func struct2_test() {
	pf()
	var p *Point2 = &Point2{}       // p := &Point{}
	var q *Point2 = &Point2{10, 10} // q := &Point
	var r *Point2 = new(Point2)     // r := new(Point)
	r.x, r.y = 100, 100
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(p.x)
	fmt.Println(p.y)
	fmt.Println(q.x)
	fmt.Println(q.y)
	fmt.Println(r.x)
	fmt.Println(r.y)
	fmt.Println(distance2(p, q))
	fmt.Println(distance2(p, r))
	fmt.Println(distance2(q, r))

}

// ==================================================================

type Point3 struct {
	x, y float64
}

func newPoint(x, y float64) *Point3 {
	p := new(Point3)
	p.x, p.y = x, y
	return p
}

func struct_array_test() {
	pf()
	var a []Point3 = []Point3{
		{x: 0, y: 0}, {10, 10}, {100, 100},
	}
	var b []*Point3 = make([]*Point3, 8)
	fmt.Println(a)
	fmt.Println(b)
	for i := 0; i < 8; i++ {
		b[i] = newPoint(float64(i), float64(i))
	}
	fmt.Println(b)
	for i := 0; i < 8; i++ {
		fmt.Println(*b[i])
	}
}

// ==================================================================
type PointX struct {
	x, y      float64
	distanceX func(*PointX) float64
}

type PointX3d struct {
	x, y, z   float64
	distanceX func(*PointX3d) float64
}

func distanceX(p, q *PointX) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func distanceX3d(p, q *PointX3d) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	dz := p.z - q.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func newPointX(x, y float64) *PointX {
	p := new(PointX)
	p.x, p.y = x, y
	p.distanceX = func(q *PointX) float64 { return distanceX(p, q) }
	return p
}

func newPointX3d(x, y, z float64) *PointX3d {
	p := new(PointX3d)
	p.x, p.y, p.z = x, y, z
	p.distanceX = func(q *PointX3d) float64 { return distanceX3d(p, q) }
	return p
}

func struct_func_test1() {
	pf()
	p1 := newPointX(0, 0)
	p2 := newPointX(10, 10)
	q1 := newPointX3d(0, 0, 0)
	q2 := newPointX3d(10, 10, 10)
	fmt.Println(p1.distanceX(p2))
	fmt.Println(q1.distanceX(q2))
}

// ==================================================================

type PointZ struct {
	x, y float64
}

type PointZ3d struct {
	x, y, z float64
}

func newPointZ(x, y float64) *PointZ {
	p := new(PointZ)
	p.x, p.y = x, y
	return p
}

func newPointZ3d(x, y, z float64) *PointZ3d {
	p := new(PointZ3d)
	p.x, p.y, p.z = x, y, z
	return p
}

func (p *PointZ) distance(q *PointZ) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func (p *PointZ3d) distance(q *PointZ3d) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	dz := p.z - q.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func struct_func_test2() {
	pf()
	p1 := newPointZ(0, 0)
	p2 := newPointZ(10, 10)
	q1 := newPointZ3d(0, 0, 0)
	q2 := newPointZ3d(10, 10, 10)
	fmt.Println(p1.distance(p2))
	fmt.Println(q1.distance(q2))
}

// ==================================================================

type Foo struct {
	a, b int
}

func (p Foo) swap() {
	p.a, p.b = p.b, p.a
	fmt.Println("Foo swap!", p)
}

type Bar struct {
	a, b int
}

func (p *Bar) swap() {
	p.a, p.b = p.b, p.a
	fmt.Println("Bar swap!", p)
}

func struct_func_test3() {
	pf()
	a := Foo{1, 2}
	b := &Foo{3, 4}
	c := Bar{5, 6}
	d := &Bar{5, 6}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	a.swap()
	b.swap()
	c.swap()
	d.swap()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// ==================================================================

type Cell struct {
	item int
	next *Cell
}

type List struct {
	top *Cell
}

func newCell(x int, cp *Cell) *Cell {
	newcp := new(Cell)
	newcp.item, newcp.next = x, cp
	return newcp
}

func newList() *List {
	lst := new(List)
	lst.top = new(Cell)
	return lst
}

func (cp *Cell) nthCell(n int) *Cell {
	i := -1
	for cp != nil {
		if i == n {
			return cp
		}
		i++
		cp = cp.next
	}
	return nil
}

func (lst *List) nth(n int) (int, bool) {
	cp := lst.top.nthCell(n)
	if cp == nil {
		return 0, false
	}
	return cp.item, true
}

func (lst *List) insertNth(n, x int) bool {
	cp := lst.top.nthCell(n - 1)
	if cp == nil {
		return false
	}
	cp.next = newCell(x, cp.next)
	return true
}

func (lst *List) deleteNth(n int) bool {
	cp := lst.top.nthCell(n - 1)
	if cp == nil || cp.next == nil {
		return false
	}
	cp.next = cp.next.next
	return true
}

func (lst *List) isEmpty() bool {
	return lst.top.next == nil
}

func (lst *List) printList() {
	cp := lst.top.next
	for ; cp != nil; cp = cp.next {
		fmt.Print(cp.item, " ")
	}
	fmt.Println("")
}

func linked_list_test() {
	pf()
	a := newList()
	for i := 0; i < 4; i++ {
		fmt.Println(a.insertNth(i, i))
	}
	a.printList()
	for i := 0; i < 5; i++ {
		n, ok := a.nth(i)
		fmt.Println(n, ok)
	}
	for !a.isEmpty() {
		a.deleteNth(0)
		a.printList()
	}

}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	struct1_test()
	struct2_test()
	struct_array_test()
	struct_func_test1()
	struct_func_test2()
	struct_func_test3()
	linked_list_test()
}
