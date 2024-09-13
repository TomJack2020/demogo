package main

import (
	"fmt"
	"time"
)

// 求和测试
func sumNumTest(n int64) int64 {

	var s, i int64 = 0, 0
	for i <= n {
		s += i
		i++
	}
	//fmt.Println(s)
	return s
}

func main() {

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
		ch <- sumNumTest(s[i])

	}

	for i := 0; i < len(s); i++ {
		res = append(res, <-ch) // 从通道 c 中接收
	}
	close(ch)

	// 获取当前时间
	t2 := time.Now()
	fmt.Println("计算结果是：", res, "\n", "耗时：", t2.Sub(t))

}
