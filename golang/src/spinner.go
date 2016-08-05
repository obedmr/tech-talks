package main
import (
	"fmt"
	"time"
)

//start main OMIT
func main() {
	go spinner(100 * time.Millisecond)
        const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
//end main OMIT

//start spinner OMIT
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
	                fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
//end spinner OMIT

//start fib OMIT
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
//end fib OMIT
