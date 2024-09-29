package main

import (
	"context"
	"fmt"
	"log"

	"github.com/risor-io/risor"
)

func main() {
	ctx := context.Background()
	script := "math.sqrt(input)"
	result, err := risor.Eval(ctx, script, risor.WithGlobal("input", 16))
        //result, err := risor.Eval(ctx, "math.min([5, 2, 7])")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The square root of 4 is:", result)
}
