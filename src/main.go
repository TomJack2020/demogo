package main

import (
	"demogo/src/walmart"
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	// sql := tools.ReadSql("sql/test.sql")
	// fmt.Println(sql)

	// amazon.AmazonTruncateTable("yibai_amazon_sku_publish_title_log")

	//只用同步一次 yibai_prod_spu表
	// product.SynYibaiProdSpu()
	// fmt.Println("同步yibai_prod_spu表完成 请前往数据库核实查看")
	// skuList := []string{"26121900001", "24111900001", "31131900001", "22111900001", "24111900002", "10111900001", "31121900042"}
	// amazon.SynAmazonTitleLog(skuList)

	// 同步spu数据到imdb
	// product.SynYibaiProdSpu()

	// 190000000   Amazon商品ID起始值  同步标题使用次数 百万每次间隔
	// for i := 0; i < 190; i++ {
	// 	startNum := i * 1000000
	// 	endNum := (i + 1) * 1000000
	// 	amazon.SynAmazonTitleLog(startNum, endNum)
	// 	fmt.Printf("同步第 % d批数据", (i + 1))
	// }

	// 64041250   Amazon商品ID起始值  同步标题使用次数 百万每次间隔
	for i := 0; i < 65; i++ {
		startNum := i * 1000000
		endNum := (i + 1) * 1000000
		walmart.SynWalmartTitleLog(startNum, endNum)
		fmt.Printf("同步第 % d批数据", (i + 1))
	}

}
