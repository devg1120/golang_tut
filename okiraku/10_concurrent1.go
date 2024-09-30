package main

import "time"
import "sync"
import "strconv"

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

func concurent1_test() {
	pf()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	fmt.Println(runtime.NumGoroutine())
}

// ==================================================================
func test1(n int, name string) {
	for i := 1; i <= n; i++ {
		fmt.Println(i, name)
		time.Sleep(100 * time.Millisecond)
	}
}
func concurent2_test() {
	pf()
	runtime.GOMAXPROCS(1)
	test1(5, "foo")
	test1(5, "bar")
	fmt.Println("......................................")
	go test1(15, "foo")
	go test1(5, "bar")
}

// ==================================================================
func test2(n int, name string, c chan<- string) {
	for i := 1; i <= n; i++ {
		fmt.Println(i, name)
		time.Sleep(100 * time.Millisecond)
	}
	c <- name
}
func concurent3_test() {
	pf()
	runtime.GOMAXPROCS(1)
	c := make(chan string)
	go test2(6, "foo", c)
	go test2(4, "bar", c)
	go test2(8, "baz", c)
	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}
}

// ==================================================================
func test3(n int, name string, wg *sync.WaitGroup) {
	for i := 1; i <= n; i++ {
		fmt.Println(i, name)
		time.Sleep(100 * time.Millisecond)
	}
	wg.Done()
}
func concurent4_test() {
	pfm("WaitGroup")
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(3)
	go test3(6, "foo", &wg)
	go test3(4, "bar", &wg)
	go test3(8, "baz", &wg)
	wg.Wait()
}

// ==================================================================
func makeRoutine(code string, in <-chan int) chan int {
	out := make(chan int)
	go func() {
		for {
			<-in
			fmt.Print(code)
			time.Sleep(50 * time.Millisecond)
			out <- 0
		}
	}()
	return out
}

func concurent5_test() {
	pfm("goroutine sync")
	runtime.GOMAXPROCS(1)
	ch1 := make(chan int)
	ch2 := makeRoutine("h", ch1)
	ch3 := makeRoutine("e", ch2)
	ch4 := makeRoutine("y", ch3)
	ch5 := makeRoutine("!", ch4)
	ch6 := makeRoutine("\n", ch5)
	for i := 0; i < 10; i++ {
		ch1 <- 0
		<-ch6
	}
}

// ==================================================================
type Item interface {
	Eq(Item) bool
	Less(Item) bool
}

type Node struct {
	item        Item
	left, right *Node
}

func newNode(x Item) *Node {
	p := new(Node)
	p.item = x
	return p
}

func searchNode(node *Node, x Item) bool {
	for node != nil {
		switch {
		case x.Eq(node.item):
			return true
		case x.Less(node.item):
			node = node.left
		default:
			node = node.right
		}
	}
	return false
}

func insertNode(node *Node, x Item) *Node {
	switch {
	case node == nil:
		return newNode(x)
	case x.Eq(node.item):
		return node
	case x.Less(node.item):
		node.left = insertNode(node.left, x)
	default:
		node.right = insertNode(node.right, x)
	}
	return node
}
func searchMin(node *Node) Item {
	if node.left == nil {
		return node.item
	}
	return searchMin(node.left)
}

func deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	return node
}
func deleteNode(node *Node, x Item) *Node {
	if node != nil {
		if x.Eq(node.item) {
			if node.left == nil {
				return node.right
			} else if node.right == nil {
				return node.left
			} else {
				node.item = searchMin(node.right)
				node.right = deleteMin(node.right)
			}
		} else if x.Less(node.item) {
			node.left = deleteNode(node.left, x)
		} else {
			node.right = deleteNode(node.right, x)
		}
	}
	return node
}

func foreachNode(f func(Item), node *Node) {
	if node != nil {
		foreachNode(f, node.left)
		f(node.item)
		foreachNode(f, node.right)
	}
}

type Tree struct {
	root *Node
}

func newTree() *Tree {
	return new(Tree)
}

func (t *Tree) searchTree(x Item) bool {
	return searchNode(t.root, x)
}

func (t *Tree) insertTree(x Item) {
	t.root = insertNode(t.root, x)
}

func (t *Tree) deleteTree(x Item) {
	t.root = deleteNode(t.root, x)
}

func (t *Tree) foreachTree(f func(Item)) {
	foreachNode(f, t.root)
}

func (t *Tree) each() chan Item {
	ch := make(chan Item)
	go func() {
		t.foreachTree(func(x Item) { ch <- x })
		close(ch)
	}()
	return ch
}
func (t *Tree) printTree() {
	t.foreachTree(func(x Item) { fmt.Print(x, " ") })
	fmt.Println("")
}

type Int int

func (n Int) Eq(m Item) bool {
	return n == m.(Int)
}

func (n Int) Less(m Item) bool {
	return n < m.(Int)
}

func concurent6_test() {
	pfm("channel range")
	runtime.GOMAXPROCS(1)
	a := newTree()
	b := []int{5, 6, 4, 3, 7, 8, 2, 1, 9, 0}
	for _, x := range b {
		a.insertTree(Int(x))
	}
	for x := range a.each() {
		fmt.Println(x)
	}
}

// ==================================================================
func (t *Tree) makeGen() func() Item {
	ch := make(chan Item)
	go func() {
		t.foreachTree(func(x Item) { ch <- x })
		close(ch)
	}()
	return func() Item { return <-ch }
}

func concurent7_test() {
	pfm("generator")
	runtime.GOMAXPROCS(1)
	a := newTree()
	b := []int{5, 6, 4, 3, 7, 8, 2, 1, 9, 0}
	for _, x := range b {
		a.insertTree(Int(x))
	}
	resume := a.makeGen()
	for i := 0; i < 11; i++ {
		fmt.Println(resume())
	}
}

// ==================================================================
type Req struct {
	Color string
	Reply chan<- int
}

func newReq(color string, ch chan int) *Req {
	req := new(Req)
	req.Color = color
	req.Reply = ch
	return req
}

func sendColor(n int, color string, ch chan<- *Req) {
	in := make(chan int)
	v := newReq(color, in)
	for ; n > 0; n-- {
		ch <- v
		<-in
		time.Sleep(100 * time.Millisecond)
	}
	ch <- nil
}

func receiveColor(n int, ch <-chan *Req) {
	for n > 0 {
		req := <-ch
		if req == nil {
			n--
		} else {
			fmt.Println(req.Color)
			req.Reply <- 0
		}
	}
}
func concurent8_test() {
	pfm("data exchange")
	runtime.GOMAXPROCS(1)
	ch := make(chan *Req)
	go sendColor(8, "red", ch)
	go sendColor(7, "blue", ch)
	go sendColor(6, "green", ch)
	receiveColor(3, ch)
}

// ==================================================================
func test1_(n int, ch, quit chan<- int) {
	for ; n > 0; n-- {
		ch <- n
		time.Sleep(500 * time.Millisecond)
	}
	quit <- 0
}

func test2_(n int, ch chan<- float64, quit chan<- int) {
	for ; n > 0; n-- {
		ch <- float64(n) / 10.0
		time.Sleep(250 * time.Millisecond)
	}
	quit <- 0
}

func test3_(n int, ch chan<- string, quit chan<- int) {
	for ; n > 0; n-- {
		ch <- strconv.Itoa(n * 10)
		time.Sleep(750 * time.Millisecond)
	}
	quit <- 0
}

func concurent9_test() {
	pfm("select")
	runtime.GOMAXPROCS(1)
	ch1 := make(chan int)
	ch2 := make(chan float64)
	ch3 := make(chan string)
	quit := make(chan int)
	go test1_(6, ch1, quit)
	go test2_(8, ch2, quit)
	go test3_(4, ch3, quit)
	for n := 3; n > 0; {
		select {
		case c := <-ch1:
			fmt.Println(c)
		case c := <-ch2:
			fmt.Println(c)
		case c := <-ch3:
			fmt.Println(c)
		case <-quit:
			n--
		default:
			fmt.Println("None")
			time.Sleep(250 * time.Millisecond)
		}
	}

}

// ==================================================================
func fibo(n int) int {
	if n < 2 {
		return 1
	} else {
		return fibo(n-2) + fibo(n-1)
	}
}
func concurent10_test() {
	pfm("timeout")
	runtime.GOMAXPROCS(1)
	ch := make(chan int, 5)
	for _, n := range []int{41, 41, 39, 35, 36} {
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
func main() {
	fmt.Println("###", fn(), "###")
	concurent1_test()
	concurent2_test()
	concurent3_test()
	concurent4_test()
	concurent5_test()
	concurent6_test()
	concurent7_test()
	concurent8_test()
	concurent9_test()
	concurent10_test()
}
