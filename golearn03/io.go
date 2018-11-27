package main

import (
	"strings"
	"fmt"
	"os"
	"time"
	"bufio"
)

func main() {
	//一个字节2的8次方,既可以是0-127的的任意数据
	//go 中文占用三个字节长度
	//1英文字符1个字节长度
	learnIOReadAt()
	n := string(127)
	m := string(257)
	s := "Golang中";
	cs := "中";
	csutf8 := "——";
	fmt.Println(len(n))
	fmt.Println(len(m))
	fmt.Println(len(csutf8))
	fmt.Println(len(s))
	fmt.Println(len(cs))
	learnIOWriteAt()
	learnIOReadfrom();
	for i:=0;i<100 ;i++  {
		go func(i int) {
			file, err := os.OpenFile("F:\\gocode\\golearn03\\testt.log", os.O_WRONLY, 0644)
			if err != nil {
				panic(err)
			}
			n, _ := file.Seek(0, os.SEEK_END)
			_, error := file.WriteAt([]byte("\r\nGo语言中文网"+string(i)), n)
			if error != nil {
				panic(error)
			}
			defer file.Close()
			defer fmt.Printf("%d--开始时间是:%v\r\n", i, time.Now())
		}(i)
	}
}

func learnIOWriteAt() {
	defer fmt.Printf("开始时间是:%v\r\n",time.Now())

	//WriteAt 从 p 中将 len(p) 个字节写入到偏移量 off 处的基本数据流中。它返回从 p 中被写入的字节数 n（0 <= n <= len(p)）以及任何遇到的引起写入提前停止的错误
	// 。若 WriteAt 返回的 n < len(p)，它就必须返回一个 非nil 的错误。

	//若 WriteAt 携带一个偏移量写入到目标中，WriteAt 应当既不影响偏移量也不被它所影响。

	//若被写区域没有重叠，可对相同的目标并行执行 WriteAt 调用。
	file, error := os.Create("F:\\gocode\\golearn03\\testt.log")
	if error != nil {
		panic(error)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余")
	n, err := file.WriteAt([]byte("Go语言中文网"),24)
	if err != nil {
		panic(err)
	}
	defer fmt.Printf("结束时间是:%v\r\n",time.Now())


	fmt.Println(n)

}
func learnIOReadAt() {
	reader := strings.NewReader("Ilovf-you")
	p := make([]byte, 7)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}

func learnIOReadfrom() {
	//ReadFrom 从 r 中读取数据，直到 EOF 或发生错误。其返回值 n 为读取的字节数。除 io.EOF 之外，在读取过程中遇到的任何错误也将被返回。
	//如果 ReaderFrom 可用，Copy 函数就会使用它
	file, err := os.Open("F:\\gocode\\golearn03\\testt.log")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}

