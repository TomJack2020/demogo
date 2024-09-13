package tools

import (
	"fmt"
	"gorm.io/gorm"
)

/*
 * ConnectAmazonCenter 连接亚马逊中心API
 */

type YibaiProdSku struct {
	Sku     string `json:"sku"`
	Spu     string `json:"spu"`
	TitleCn string `json:"title_cn"`
	Price   string `json:"title_en"`
}

func ConnectProductSystem() *gorm.DB {
	// TODO: connect to Amazon Center API
	res := ConfigDbConnUrl("")
	// 连接数据库 本地ck库数据连接
	db := GetDbConByGorm(res["productSystem"], "mysql", "yibai_prod_base.")
	return db
}

func TestQueryProductSystem() {
	db := ConnectProductSystem()
	var products []YibaiProdSku
	// 返回数据类型为[]YibaiProdSku
	db.Limit(10).Where("sku = ?", "H38676B").Find(&products)
	fmt.Println(products)

}
