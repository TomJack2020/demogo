package main

import (
	"fmt"
	"time"
)

//以下实例通过两个 goroutine 来计算数字之和，在 goroutine 完成计算后，它会计算两个结果的和：

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

// 斐波那契函数测试阻塞通道
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// 求和测试
func sumNum(n int64) int64 {

	var s, i int64 = 0, 0
	for i <= n {
		s += i
		i++
	}
	//fmt.Println(s)
	return s
}

func main() {
	//s := []int{7, 2, 8, -9, 4, 0, 1, 3, 4, 5, 5}
	//c := make(chan int, 2) // 定义一个容量为 2 的通道
	//
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//
	//x, y := <-c, <-c       // 从通道 c 中接收
	//fmt.Println(x, y, x+y) // 输出结果

	//c1 := make(chan int, 10) // 定义一个容量为 10 的通道
	//go fibonacci(cap(c1), c1)
	//// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	//// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	//// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	//// 会结束，从而在接收第 11 个数据的时候就阻塞了。

	//sumNum(100000000) // 输出 55

	// 获取当前时间
	t := time.Now()

	// 测试并发计算
	var s []int64
	for j := 0; j < 1000; j++ {
		s = append(s, int64(j)*10000)
	}

	fmt.Println(s)

	var res []int64 // 定义结果数据集

	//// 采取for循环模式5000数组 耗时： 43.0813292s
	//for i := 0; i < len(s); i++ {
	//	res = append(res, sumNum(s[i]))
	//}

	// 采取并发计算 5000组数据 耗时 耗时： 33.2083s
	ch := make(chan int64, len(s))
	for i := 0; i < len(s); i++ {
		ch <- sumNum(s[i])

	}

	for i := 0; i < len(s); i++ {
		res = append(res, <-ch) // 从通道 c 中接收
	}
	close(ch)

	// 获取当前时间
	t2 := time.Now()
	fmt.Println("计算结果是：", res, "\n", "耗时：", t2.Sub(t))

}
