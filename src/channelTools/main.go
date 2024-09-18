package main

import (
	_case "demogo/src/channelTools/case"
	"fmt"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// 模拟工作
	// ...
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// wg := sync.WaitGroup{} //
	// for i := 1; i <= 5; i++ {
	// 	// 计数器加 1
	// 	wg.Add(1)
	// 	// 启动一个 goroutine
	// 	go func(id int) {
	// 		//
	// 		worker(id)
	// 		// 计数器减 1
	// 		wg.Done()
	// 	}(i)
	// }
	// wg.Wait() // 等待所有 goroutine 完成
	// fmt.Println("All goroutines finished")

	_case.GetCon()

}
