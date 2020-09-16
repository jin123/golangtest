package main

import "fmt"

func main() {

	map1 := make(map[string]string, 0)
	map1["one"] = "one"
	map1["two"] = "two"
	map1["three"] = "three"
	fmt.Println(len(map1))
	toWork(0, 10)
}

func toWork(start, end int) {
	//子Go程与主Go程完成同步，意思是子Go程没有全部执行完毕，主Go程不许退出。
	page := make(chan int)

	for i := start; i <= end; i++ {
		//开启子GO程
		go SpiderPage(i, page)
	}
	//主GO程开始消费管道中的数据
	for i := start; i <= end; i++ {
		fmt.Println("爬取完成", <-page)
	}
}
func SpiderPage(index int, page chan int) {
	page <- index
}
