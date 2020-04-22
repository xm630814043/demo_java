package main

import (
	"demo_crawler/requetsq"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//url := "http://www.baidu.com/"
	//queue := make(chan string)
	//go func() {
	//	queue <- url
	//}()
	//for uri := range queue {
	//	requetsq.Downloads(uri, queue)
	//}

	start := time.Now()
	ch := make(chan bool)
	for i := 0; i < 10; i++ {
		//strconv.Itoa()函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字
		//go
		go requetsq.ParseUrls("https://movie.douban.com/top250?start="+strconv.Itoa(25*i), ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
	elapsed := time.Since(start)
	fmt.Printf("Took %s", elapsed)
}