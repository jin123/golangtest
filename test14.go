package main
 
import (
	"fmt"
	"time"
)
 
func main() {
	var ch chan int = make(chan int, 10)
	var x int = 0
	var nonempty bool
 
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
		x, nonempty = <-ch
		fmt.Println("x =", x, ", length of ch is", len(ch), ", nonempty", nonempty)
		x, nonempty = <-ch
		fmt.Println("x =", x, ", length of ch is", len(ch), ", nonempty", nonempty)
		x, nonempty = <-ch
		fmt.Println("x =", x, ", length of ch is", len(ch), ", nonempty", nonempty)
		x, nonempty = <-ch
		fmt.Println("x =", x, ", length of ch is", len(ch), ", nonempty", nonempty)
		
	}()
	select {
	case ch <- 3:
		fmt.Println("ch write 3")
	case ch <- 4:
		fmt.Println("ch write 4")
	case ch <- 5:
		fmt.Println("ch write 5")
	case ch <- x:
		fmt.Println("ch write", x)
	case x = <-ch:
		fmt.Println("ch read", x)
	default:
		fmt.Println("nothing to do")
	}
	time.Sleep(time.Second*2)
}