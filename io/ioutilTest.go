package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func main() {

	test6()
}

func test1() {
	content, err := ioutil.ReadFile("./1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file content: %s", content)
}

//WriteFile将数据写入由filename命名的文件。
//如果该文件不存在，则WriteFile使用权限perm创建它; 否则WriteFile会在写入之前截断它。
func test2() {
	content := "Hello world"
	//io.WriterTo
	if err := ioutil.WriteFile("./2.txt", []byte(content), 0644); err != nil {
		fmt.Println(err)
	}
	fmt.Println("write file success...")
}

func test3() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

/**
Len作用: 返回未读的字符串长度

Size的作用:返回字符串的长度

read的作用: 读取字符串信息
*/
func test4() {

	r := strings.NewReader("123456789")
	fmt.Println(r, reflect.TypeOf(r))
	fmt.Println(r.Len()) // 输出14  初始时，未读长度等于字符串长度
	var buf []byte
	buf = make([]byte, 9)
	readLen, err := r.Read(buf)
	fmt.Println("读取到的长度:", readLen, string(buf)) //读取到的长度5
	if err != nil {
		fmt.Println("错误:", err)
	}
	fmt.Println(buf, reflect.TypeOf(buf)) //adcde
	fmt.Println(r.Len())                  //9   读取到了5个 剩余未读是14-5
}

func test5() {
	file := strings.NewReader("testline \n 111111")
	/*file, err := os.Open("./1.txt")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}*/
	read := bufio.NewReader(file)
	for {
		c, pk, err := read.ReadLine()
		fmt.Printf("这一行的内容是：%s；pk是：%t\n", c, pk)
		if err == io.EOF {
			fmt.Printf("到头了\n")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Printf("读取错误")
			break
		}

	}
}

//这种适合没有换行符的文件读取
func test6() {
	FileHandle, err := os.Open("3.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()
	// 设置每次读取字节数
	buffer := make([]byte, 14)
	for {
		n, err := FileHandle.Read(buffer)
		//fmt.Println("n=", n)
		content := string(buffer[:n])
		fmt.Println("content=", content)
		// 控制条件,根据实际调整
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		if n == 0 {
			break
		}
		// 如下代码打印出每次读取的文件块(字节数)

	}
}

func test7() {
	// buffer reader
	var r io.Reader = bufio.NewReader(strings.NewReader(string("hello, world")))
	fmt.Println(r)

}

func test8() {
	s := bufio.NewScanner(strings.NewReader("/*block comments*///line comments\n{'type':4}"))
	fmt.Println(s)
	s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// read more.
		return 0, nil, nil
	})
	for s.Scan() {
		fmt.Println(s.Text())
	}
	fmt.Println("err is", s.Err())

}

func test9() {
	FileHandle, err := os.Open("2.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer FileHandle.Close()
	lineScanner := bufio.NewScanner(FileHandle)
	for lineScanner.Scan() {
		// 相同使用场景下可以使用如下方法
		// func (s *Scanner) Bytes() []byte
		// func (s *Scanner) Text() string
		// 实际逻辑 : 对读取的内容进行某些业务操作
		// 如下代码打印每次读取的文件行内容
		fmt.Println(lineScanner.Text())
	}
}

//这种适合有换行的文件读取
func test10() {
	//fmt.Println(os.O_CREATE, os.O_RDWR, os.O_APPEND)
	file, err := os.OpenFile("3.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error=", err)
	}
	defer file.Close()
	content := []byte("hello world1! \n hello world2! \n hello world3! \n")
	newWriter := bufio.NewWriterSize(file, 1024)
	if _, err = newWriter.Write(content); err != nil {
		fmt.Println(err)
	}
	if err = newWriter.Flush(); err != nil {
		fmt.Println("err2=", err)
	}

	file2, _ := os.OpenFile("3.txt", os.O_RDWR, 0666)
	defer file2.Close()
	read := bufio.NewReader(file2)
	for {
		line, _, err := read.ReadLine()
		fmt.Println("open error=", err, err == io.EOF)
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println("line=", string(line))
	}
}

func test11() {
	f, err := os.OpenFile("3.txt", os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理

	} else {
		content := "\n写入的文件内容\n"

		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)
		//fmt.Println("n=", strconv.FormatInt(n, 10), reflect.TypeOf(n))
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}

	defer f.Close()
}

func test12() {
	url := "https://www.baidu.com/img/PCfb_5bf082d29588c07f842ccde3f97243ea.png"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}
	defer resp.Body.Close()
	out, err := os.Create("test.png")
	wt := bufio.NewWriter(out)
	defer out.Close()
	n, err := io.Copy(wt, resp.Body)
	fmt.Println("write", n)
	if err != nil {
		panic(err)
	}
	wt.Flush()
}
