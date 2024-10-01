package main
import (
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
type Stream chan int

func makeInt(n, m int) Stream {
    s := make(Stream)
    go func() {
        for i := n; i <= m; i++ {
            s <- i
        }
        close(s)
    }()
    return s
}

func makeNum(n int) Stream {
    s := make(Stream)
    go func() {
        for { s <- n }
    }()
    return s
}

func makeFibo() Stream {
    s := make(Stream)
    go func() {
        a, b := 1, 1
        for {
            s <- a
            a, b = b, a + b
            if a < 0 { break }
        }
        close(s)
    }()
    return s
}

func stream1_test() {
   pf()
    runtime.GOMAXPROCS(1)
    s0 := makeInt(1, 20)
    for x := range s0 {
        fmt.Print(x, " ")
    }
    fmt.Println("")
    s1 := makeNum(1)
    for i := 0; i < 10; i++ {
        fmt.Print(<- s1, " ")
    }
    fmt.Println("")
    s2 := makeFibo()
    for x := range s2 {
        fmt.Print(x, " ")
    }
    fmt.Println("")


}

// ==================================================================
func streamMap(f func(int) int, in Stream) Stream {
    s := make(Stream)
    go func(){
        for {
            x, ok := <- in
            if !ok { break }
            s <- f(x)
        }
        close(s)
    }()
    return s
}

func streamFilter(f func(int) bool, in Stream) Stream {
    s := make(Stream)
    go func(){
        for {
            x, ok := <- in
            if !ok { break }
            if f(x) {
                s <- x
            }
        }
        close(s)
    }()
    return s
}

func stream2_test() {
   pf()
    square := func(x int) int { return x * x }
    s3 := streamMap(square, makeInt(1, 10))
    for x := range s3 {
        fmt.Print(x, " ")
    }
    fmt.Println("")
    isOdd := func(x int) bool { return x % 2 != 0 }
    s4 := streamFilter(isOdd, makeInt(1, 20))
    for x := range s4 {
        fmt.Print(x, " ")
    }
    fmt.Println("")
}

// ==================================================================

// 定数
const (
    GET   = 0
    RET   = 1
    LEFT  = 2
    RIGHT = 3
)

// フォークのリクエスト
type Req struct {
    req, fork, side int
    reply chan bool
}

// リクエストの生成
func newReq(req, fork, side int, reply chan bool) *Req {
    p := new(Req)
    p.req = req
    p.fork = fork
    p.side = side
    p.reply = reply
    return p
}

// フォークの管理
func forks(n int, ch chan *Req) {
    forkTbl := make([]bool, n)
    for i := 0; i < n; i++ {
        forkTbl[i] = true
    }
    for {
        r := <- ch
        switch r.req {
        case GET:
            if forkTbl[r.fork] {
                if n == 1 && r.side == RIGHT {
                    r.reply <- false
                } else {
                    forkTbl[r.fork] = false
                    n--
                    r.reply <- true
                }
            } else {
                r.reply <- false
            }
        case RET:
            forkTbl[r.fork] = true
            n++
            r.reply <- true
        }
    }
}

// フォークの取得
func getFork(fork, side int, out chan *Req, in chan bool) bool {
    r := newReq(GET, fork, side, in)
    for {
        out <- r
        if <- in {
            time.Sleep(100 * time.Millisecond)
            return true
        } else if side == LEFT {
            return false
        } else {
            time.Sleep(500 * time.Millisecond)
        }
    }
}

// フォークの返却
func retFork(fork, side int, out chan *Req, in chan bool) bool {
    time.Sleep(100 * time.Millisecond)
    out <- newReq(RET, fork, side, in)
    return <- in
}

// 哲学者の動作
func person(m, forkR, forkL int, out chan *Req, quit chan bool) {
    in := make(chan bool)
    for n := 2 ; n > 0; {
        fmt.Printf("Philosopher%d is thinking\n", m)
        time.Sleep(1000 * time.Millisecond)
        getFork(forkR, RIGHT, out, in)
        if getFork(forkL, LEFT, out, in) {
            fmt.Printf("Philosopher%d is eating\n", m)
            time.Sleep(500 * time.Millisecond)
            retFork(forkR, RIGHT, out, in)
            retFork(forkL, LEFT, out, in)
            n--
        } else {
            retFork(forkR, RIGHT, out, in)
        }
    }
    fmt.Printf("Philosopher%d is sleeping\n", m)
    quit <- true
}

func philosopher1_test() {
    runtime.GOMAXPROCS(1)
    ch := make(chan *Req)
    quit := make(chan bool)
    go forks(5, ch)
    go person(1, 0, 1, ch, quit)
    go person(2, 1, 2, ch, quit)
    go person(3, 2, 3, ch, quit)
    go person(4, 3, 4, ch, quit)
    go person(5, 4, 0, ch, quit)
    for n := 5; n > 0; n-- {
        <- quit
    }

}

// ==================================================================
func main() {
	fmt.Println("###", fn(), "###")
	stream1_test()
	stream2_test()
        philosopher1_test() 
}
