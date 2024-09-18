package main

import (
	_amazon "demogo/src/AmazonCenter/amazon"
	"fmt"
	"sync"
)

func main() {
	_amazon.Print("Hello World")
	// 调用函数 同步标题
	SyncAmazonTitleUseLog()
}

func SyncAmazonTitleUseLog() {
	wg := sync.WaitGroup{}
	// 190000000   Amazon商品ID起始值  同步标题使用次数 百万每次间隔
	for i := 0; i < 190; i++ {
		wg.Add(1)
		go func(i int) {
			startNum := i * 1000000
			endNum := (i + 1) * 1000000
			// 调用函数实现同步
			_amazon.SynAmazonTitleLog(startNum, endNum)
			fmt.Println("第", i+1, "次同步完成")
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("同步完成")
}
