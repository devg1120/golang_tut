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
type Foo struct {
	a, b int
}

func (x Foo) printA() {
	fmt.Println("a =", x.a)
}

func (x Foo) printB() {
	fmt.Println("b =", x.b)
}

type Bar struct {
	foo Foo
	c   int
}

func (x Bar) printC() {
	fmt.Println("c =", x.c)
}

type Baz struct {
	Foo
	c int
}

func (x Baz) printB() {
	fmt.Print("Baz:")
	x.Foo.printB()
}

func (x Baz) printC() {
	fmt.Println("c =", x.c)
}

func struct_embedded1() {
	pf()
	x := Foo{1, 2}
	y := Bar{Foo{10, 20}, 30}
	z := Baz{Foo{100, 200}, 300}
	x.printA()
	x.printB()
	y.foo.printA()
	y.foo.printB()
	y.printC()
	z.printA()
	z.printB()
	z.printC()

}

//==================================================================

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

type FixedList struct {
	List
	size, limit int
}

func newFixedList(limit int) *FixedList {
	p := new(FixedList)
	p.top = new(Cell)
	p.limit = limit
	return p
}

func (p *FixedList) insertNth(n, x int) bool {
	if p.size >= p.limit {
		return false
	}
	result := p.List.insertNth(n, x)
	if result {
		p.size++
	}
	return result
}

func (p *FixedList) deleteNth(n int) bool {
	result := p.List.deleteNth(n)
	if result {
		p.size--
	}
	return result
}

func fixed_list_test() {
	pf()

	a := newList()
	for i := 0; i < 8; i++ {
		fmt.Println(a.insertNth(i, i))
	}
	a.printList()
	for i := 0; i < 8; i++ {
		n, ok := a.nth(i)
		fmt.Println(n, ok)
	}
	for !a.isEmpty() {
		a.deleteNth(0)
		a.printList()
	}

	b := newFixedList(6)
	for i := 0; i < 8; i++ {
		fmt.Println(b.insertNth(i, i))
	}
	b.printList()
	for i := 0; i < 8; i++ {
		fmt.Println(b.nth(i))
	}
	for !b.isEmpty() {
		b.deleteNth(0)
		b.printList()
	}
}

//==================================================================

type Stack struct {
	content *List
}

func newStack() *Stack {
	st := new(Stack)
	st.content = newList()
	return st
}

func (st *Stack) push(x int) {
	st.content.insertNth(0, x)
}

func (st *Stack) pop() (int, bool) {
	x, ok := st.content.nth(0)
	if ok {
		st.content.deleteNth(0)
	}
	return x, ok
}

func (st *Stack) top() (int, bool) {
	return st.content.nth(0)
}

func (st *Stack) isEmpty() bool {
	return st.content.isEmpty()
}

func stack_test() {
	pf()
	st := newStack()
	for i := 0; i < 8; i++ {
		st.push(i)
		fmt.Println(st.top())
	}
	for !st.isEmpty() {
		fmt.Println(st.pop())
	}

}

// ==================================================================
type Queue struct {
	size        int
	front, rear *Cell
}

func newQueue() *Queue {
	return new(Queue)
}

func (q *Queue) enqueue(x int) {
	cp := newCell(x, nil)
	if q.size == 0 {
		q.front = cp
		q.rear = cp
	} else {
		q.rear.next = cp
		q.rear = cp
	}
	q.size++
}

func (q *Queue) dequeue() (int, bool) {
	if q.size == 0 {
		return 0, false
	} else {
		x := q.front.item
		q.front = q.front.next
		q.size--
		if q.size == 0 {
			q.rear = nil
		}
		return x, true
	}
}

func (q *Queue) isEmpty() bool {
	return q.size == 0
}

func queue_test() {
	pf()
	que := newQueue()
	for i := 0; i < 8; i++ {
		que.enqueue(i)
		fmt.Println(i)
	}
	for !que.isEmpty() {
		fmt.Println(que.dequeue())
	}
}

// ==================================================================
type RQueue struct {
	front, rear, cnt int
	buff             []int
}

func makeQueue(size int) *RQueue {
	q := new(RQueue)
	q.buff = make([]int, size)
	return q
}

func (q *RQueue) isEmpty() bool {
	return q.cnt == 0
}

func (q *RQueue) isFull() bool {
	return q.cnt == len(q.buff)
}

func (q *RQueue) enqueue(x int) bool {
	if q.isFull() {
		return false
	}
	q.buff[q.rear] = x
	q.cnt++
	q.rear++
	if q.rear >= len(q.buff) {
		q.rear = 0
	}
	return true
}

func (q *RQueue) dequeue() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	x := q.buff[q.front]
	q.cnt--
	q.front++
	if q.front >= len(q.buff) {
		q.front = 0
	}
	return x, true
}

func (q *RQueue) top() (int, bool) {
	if q.isEmpty() {
		return 0, false
	}
	return q.buff[q.front], true
}

func (q *RQueue) clear() {
	q.front = 0
	q.rear = 0
	q.cnt = 0
}

func (q *RQueue) length() int {
	return q.cnt
}

func ring_buffer_queue_test() {

	pf()
	q := makeQueue(10)
	fmt.Println(q.isEmpty())
	fmt.Println(q.length())
	for i := 0; i < 10; i++ {
		q.enqueue(i)
	}
	fmt.Println(q.isFull())
	fmt.Println(q.length())
	for !q.isEmpty() {
		fmt.Println(q.dequeue())
	}
	fmt.Println(q.isEmpty())
	fmt.Println(q.length())

}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	struct_embedded1()
	fixed_list_test()
	stack_test()
	queue_test()
	ring_buffer_queue_test()
}
