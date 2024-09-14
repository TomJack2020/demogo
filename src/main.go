package main

import (
<<<<<<< HEAD
	"demogo/src/walmart"
=======
	"demogo/tools"
>>>>>>> b8e0d48569c3ea2d9b65f42cfc63c5ab8a0334f9
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	// sql := tools.ReadSql("sql/test.sql")
	// fmt.Println(sql)

<<<<<<< HEAD
	// amazon.AmazonTruncateTable("yibai_amazon_sku_publish_title_log")
=======
	fmt.Println("hello world")

	res := tools.ConfigDbConnUrl("")
	fmt.Println(res)

	//db := tools.ConnectProductSystem()
	//var product []tools.YibaiProdSku
	//db.Limit(10).Where("sku = ?", "H38676B").Find(&product)
	//fmt.Println(product)
>>>>>>> b8e0d48569c3ea2d9b65f42cfc63c5ab8a0334f9

	//只用同步一次 yibai_prod_spu表
	// product.SynYibaiProdSpu()
	// fmt.Println("同步yibai_prod_spu表完成 请前往数据库核实查看")
	// skuList := []string{"26121900001", "24111900001", "31131900001", "22111900001", "24111900002", "10111900001", "31121900042"}
	// amazon.SynAmazonTitleLog(skuList)

<<<<<<< HEAD
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
=======
	//rss := t.SelectSelectResultWithAccountByPublishId(7638)
	//fmt.Println(rss)
>>>>>>> b8e0d48569c3ea2d9b65f42cfc63c5ab8a0334f9

}
