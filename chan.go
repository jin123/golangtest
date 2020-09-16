package main

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex
var m *sync.RWMutex

var done chan bool

func HelloWorld() {
	fmt.Println("Hello world goroutine")
	time.Sleep(1 * time.Second)
	//done <- true
}
func main() {
	//c := make(chan string, 1)
	testSelect2()
}

//读写通道
func wrChan() {
	// 创建 chan
	c := make(chan string)
	//fmt.Println(c)

	//go channelF1Reader(c)
	go func() {
		c <- "Good11111111222222!"
	}()
	fmt.Println("start testing......")
	cc := <-c

	//go channelF2Writer(c)

	// go channelF3ReaderWriter(c)

	fmt.Println(cc)

}

//定义 只读通道 函数
func channelF1Reader(mes <-chan string) {
	msg := <-mes

	fmt.Println(msg)
}

//定义 只写通道 函数
func channelF2Writer(mes chan<- string) {
	mes <- "www.ydook.com"
}

//定义 读写通道 函数
func channelF3ReaderWriter(mes chan string) {
	msg := <-mes
	fmt.Println(msg)
	mes <- "YDOOK"
}

func simpleChan() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine")
		// 通过通道通知main的goroutine
		ch <- 0
		fmt.Println("exit goroutine")
	}()
	fmt.Println("wait goroutine")
	// 等待匿名goroutine
	<-ch
	fmt.Println("all done")
}
func testSelect() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("我是生产者")
		ch <- 1
		ch <- 2
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("我是消费者")
		fmt.Println(<-ch)
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
	/*go func() {

		ch <- 1
		fmt.Println("go 1 ")
	}()
	go func() {

		ch <- 2
		fmt.Println("go 2 ")
	}()*/
	<-ch
	//select {
	//case ch <- 1:
	//	fmt.Println("坐等补充空间 ")
	//case ch <- 2:
	//	fmt.Println("坐等补充空间2 ")
	//case <-ch:
	//	fmt.Println("坐等消费")
	//case <-time.After(time.Second * 1):
	//	fmt.Println("timeout")
	//default:
	//	fmt.Println("default")

	//}
	fmt.Println("go exit ")
}
func testSelect2() {
	c := make(chan interface{})
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	x := 0
	y := 0
	ch1 <- 1
	ch2 <- 2
	time.Sleep(time.Second * 2)
	select { //随机进入
	case <-c:
	///	fmt.Println("Unblocked %v later.\n")

	case x = <-ch1:

		fmt.Println("ch1 case...", x)
	case y = <-ch2:

		fmt.Println("ch2 case...", y)
	default:

		fmt.Println("default go... \n ")
	}
}
func testRwLocked() {
	m = new(sync.RWMutex)

	// 多个同时读
	go read(1)
	go read(2)

	time.Sleep(5 * time.Second)
}

func write(i int) {
	println(i, "write start")
	m.Lock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	m.Unlock()
	println(i, "write over")
}

func read(i int) {
	println(i, "read start")

	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over")
	i++
}

func testLock() {
	//var a map[int]int
	a := make(map[int]int, 5)

	mySlice := make([]int, 2, 4)
	fmt.Println(mySlice)
	//mySlice[0] = 7
	//fmt.Println(mySlice)
	//mySlice[0]++

	mySlice[3] = 1

	fmt.Println(mySlice)

	return

	//fmt.Println(b)
	a[4] = 4
	a[3] = 3
	a[2] = 2
	a[1] = 1
	a[5] = 5
	a[6] = 1
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			lock.Lock()
			b[8] = 100
			lock.Unlock()
		}(a)
	}

	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second)
	fmt.Println(a)
}
func test_defer() (i int) {
	defer func() {

		fmt.Println(i)
		i++

	}()
	return 10
}
func manyDefer() (i int) {
	//var i int

	defer func() {
		i++
		fmt.Println("a defer2:", i) // 打印结果为 a defer2: 2
	}()

	defer func() {
		i++
		fmt.Println("a defer1:", i) // 打印结果为 a defer1: 1
	}()
	//j = i
	return i
}

func defer_func() {
	fmt.Println(f11())
	//fmt.Println(f4());
}

func f11() (result int) {
	//	result = 0   //先给返回值赋值
	func() { //再执行defer 函数
		result++
	}()
	return //最后返回
}

func f22() (r int) {
	t := 5
	r = t    //赋值指令
	func() { //defer 函数被插入到赋值与返回之间执行，这个例子中返回值r没有被修改
		t = t + 5
	}()
	return //返回
}

func f33() (t int) {
	t = 5 //赋值指令
	func() {
		t = t + 5 //然后执行defer函数,t值被修改
	}()
	return 11
}

func f44() (r int) {
	r = 1         //给返回值赋值
	func(r int) { //这里的r传值进去的，是原来r的copy，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return
}

func f1() (result int) {
	//fmt.Println(result)
	result += 100
	fmt.Println("1")
	fmt.Println(result)
	defer func() {
		fmt.Println(result)
		//result +=100
		fmt.Println("2")
		fmt.Println(result)
	}()
	fmt.Println("3")
	fmt.Println(result)
	return 0
}
func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (t int) {
	t = 5
	defer func() {
		t += 5
	}()
	fmt.Println(t)
	fmt.Println(t)
	return t
}
func f4() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 100
}

func quitChan() {
	i := 0
	ch := make(chan string, 0)
	defer func() {
		close(ch)
	}()

	go func() {
	DONE:
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {

			case m := <-ch:
				println(m)
				break DONE // 跳出 select 和 for 循环
			default:
				//fmt.Println("go die")
			}
		}
	}()

	time.Sleep(time.Second * 4)
	ch <- "stop2020"
}

func bufferChan() {

	ch := make(chan string, 3) // 创建了缓冲区为3的通道
	a := len(ch)               // 长度计算
	b := cap(ch)               // 容量计算
	fmt.Println(a)
	fmt.Println(b)
}

func singleChan() {
	inc := make(chan string)
	outc := make(chan string)

	go Echo(inc)
	fmt.Println(inc)
	go Receive(outc, inc)

	getStr := <-outc // 接收goroutine 2的返回

	fmt.Println(getStr)
}

// 定义goroutine 1
func Echo(out chan<- string) { // 定义发送通道类型
	time.Sleep(1 * time.Second)
	out <- "咖啡色的羊驼111111"
	close(out)
}

// 定义goroutine 2
func Receive(out chan<- string, in <-chan string) { // 定义发送通道类型和接受类型
	temp := <-in // 阻塞等待echo的通道的返回
	fmt.Println(temp)
	out <- temp
	close(out)
}

func errDisplay() {
	defer func() {
		fmt.Println("defer start")
		if err := recover(); err != nil {
			fmt.Println("system error")
		}
	}()

}

/*func f() {
defer func() { 　　　//必须要先声明defer，否则不能捕获到panic异常
　　　　fmt.Println("defer start")
　　　　if err := recover(); err != nil {
　　　　　　fmt.Println(err) //这里的err其实就是panic传入的内容，"bug"
　　　　}
　　　　fmt.Println("defer end")
　　}()
　　//for {
　　　　fmt.Println("func begin")
　　　　a := []string{"a", "b"}
　　　　fmt.Println(a[3]) 　　   // 越界访问，肯定出现异常
　　　　panic("bug")  　　　     // 上面已经出现异常了,所以肯定走不到这里了。
　　　　fmt.Println("func end") // 不会运行的.
　　　　time.Sleep(1 * time.Second)
　　//}
}*/
func receiveMsg() {
	msg := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		msg <- "goroutine 1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		msg <- "goroutine 2"
	}()
	go func() {
		time.Sleep(time.Second * 3)
		msg <- "goroutine 3"
	}()
	t := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	//
	//go func() {
	for i := 0; i <= len(t)-1; i++ {
		fmt.Println(t[i])
	}
	// }()
	//time.Sleep(time.Second * 4)
}

func receiveMsg2() {
	msg := make(chan string, 100)
	done := make(chan string, 100)
	//fmt.Println("receiveMsg2")
	go func() {
		time.Sleep(time.Second * 3)
		msg <- "goroutine 1"
		done <- "goroutine 1 done"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		msg <- "goroutine 2"
		done <- "goroutine 2 done"
	}()
	go func() {
		time.Sleep(time.Second * 1)
		msg <- "goroutine 3"
		done <- "goroutine 3 done"
	}()

	time.Sleep(time.Second * 3)
	go func() {
		//This statement is to make sure that all message is received.

		for i := range msg {
			fmt.Println("message :", i)
		}
	}()
	go func() {
		for j := 0; j < 3; j++ {
			fmt.Println("done status :", <-done)
		}
	}()

}
func receiveMsg3() {
	msg := make(chan string)
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 3)
		msg <- "goroutine 1"
		fmt.Println("Done 1")
	}()
	go func() {

		defer wg.Done()
		time.Sleep(time.Second * 2)
		msg <- "goroutine 2"
		fmt.Println("Done 2")
	}()
	go func() {

		defer wg.Done()
		time.Sleep(time.Second * 1)
		msg <- "goroutine 3"
		fmt.Println("Done 3")
	}()
	go func() {
		time.Sleep(time.Second * 3)
		for i := range msg {
			fmt.Println("message :", i)
		}
	}()
	wg.Wait()
}
func gotest2() {

	go func() {
		time.Sleep(time.Second * 10)
		fmt.Println("11111111")
	}()
	go func() {
		time.Sleep(time.Second * 20)
		fmt.Println("22222222222")
	}()
	time.Sleep(time.Second * 40)
	fmt.Println("33333333")
	/*
		t3 := make(chan string, 2)

			go func() {
				fmt.Println("11111111")
				t3 <- "value"
				fmt.Println("33333333333")

			}()
			fmt.Println("22222222222")
			fmt.Println(<-t3)

			fmt.Println("4444444")

		go func() {
			time.Sleep(time.Second * 3)
			//fmt.Println(<-t3)
		}()
		t3 <- "value1"
		t3 <- "value2"
		t3 <- "value3"

		fmt.Println("end ........")
	*/

}
func gotest() {
	channel := make(chan string, 2)

	go func() {
		fmt.Println("sleep 1")
		time.Sleep(100 * time.Second)
		fmt.Println("sleep 2")
	}()

	fmt.Println("1")
	channel <- "h1"
	fmt.Println("2")
	channel <- "w2"

	fmt.Println("3")
	channel <- "c3"

	fmt.Println("...")
	msg1 := <-channel
	fmt.Println(msg1)

	/*var c = make(chan bool)

	var nums int = 0
	for i := 0; i < 3; i++ {
		//fmt.Println("************")
		//fmt.Printf(" main %d  begin------------", i)
		go func() {
			time.Sleep(10 * time.Second)
			nums++
			fmt.Println("")
			fmt.Printf(" go %d begin------------", nums)
			if nums == 2 {
				c <- true
			}
			//	fmt.Println(" *** ** ")
			//	fmt.Printf("go  %d end ------------", i)
		}()
		//fmt.Println("************")
		//fmt.Printf(" main  %d  end ------------", i)

	}
	//fmt.Println("exit ?")
	<-c
	fmt.Println("exit 2 ?")*/

}
