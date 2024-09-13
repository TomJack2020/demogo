package main

import (
	"demogo/tools"
	"fmt"
)

// localMyDb  localCkDbUrl   amazonCenter

func main() {

	fmt.Println("hello world")

	res := tools.ConfigDbConnUrl("")
	fmt.Println(res)

	//db := tools.ConnectProductSystem()
	//var product []tools.YibaiProdSku
	//db.Limit(10).Where("sku = ?", "H38676B").Find(&product)
	//fmt.Println(product)

	// tools.InsertProduct(db, tools.YibaiProdSku{Sku: "H38676B", Name: "测试商品", Price: 100, Stock: 100})\
	//var productMap []map[string]interface{}
	//db.Limit(10).Model(&tools.YibaiProdSku{}).Select("sku,spu,title_cn").Find(&productMap)
	//fmt.Println(productMap)

	//rss := t.SelectSelectResultWithAccountByPublishId(7638)
	//fmt.Println(rss)

}
