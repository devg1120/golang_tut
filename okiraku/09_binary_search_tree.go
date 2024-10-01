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

func demo1_test() {
   pf()
}
func demo2_test() {
   pfm("ABC")
}

// ==================================================================
type Item interface {
    Eq(Item)   bool
    Less(Item) bool
}

type Node struct {
    item Item
    left, right *Node
}

func newNode(x Item) *Node {
    p := new(Node)
    p.item = x
    return p
}

type Tree struct {
    root *Node
}

func newTree() *Tree {
    return new(Tree)
}

func searchNode(node *Node, x Item) bool {
    for node != nil {
        switch {
        case x.Eq(node.item): return true
        case x.Less(node.item): node = node.left
        default: node = node.right
        }
    }
    return false
}

func (t *Tree) searchTree(x Item) bool {
    return searchNode(t.root, x)
}

func insertNode(node *Node, x Item) *Node {
    switch {
    case node == nil: return newNode(x)
    case x.Eq(node.item): return node
    case x.Less(node.item):
        node.left = insertNode(node.left, x)
    default:
        node.right = insertNode(node.right, x)
    }
    return node
}

func (t *Tree) insertTree(x Item) {
    t.root = insertNode(t.root, x)
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

func (t *Tree) deleteTree(x Item) {
    t.root = deleteNode(t.root, x)
}

func foreachNode(f func(Item), node *Node) {
    if node != nil {
        foreachNode(f, node.left)
        f(node.item)
        foreachNode(f, node.right)
    }
}

func (t *Tree) foreachTree(f func(Item)) {
    foreachNode(f, t.root)
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

func simple_test() {
   pf()
    a := newTree()
    for _, v := range []int{5,6,4,7,3,8,2,9,1,0} {
        a.insertTree(Int(v))
    }
    a.printTree()
    for i := 0; i < 10; i++ {
        fmt.Println(a.searchTree(Int(i)))
    }
    for i := 0; i < 10; i++ {
        a.deleteTree(Int(i))
        a.printTree()
    }
}
// ==================================================================
// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	//demo1_test()
	//demo2_test()
	simple_test()
}
