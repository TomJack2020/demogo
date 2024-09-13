package tools

import (
	"gorm.io/gorm"
)

type YibaiAmazonPublishSkuSelectResult struct {
	AccountId string `json:"account_id"`
	PublishId string `json:"publish_id"`
	Sku       string `json:"sku"`
	Title     string `json:"title"`
}

type YibaiSystemAccount struct {
	AccountId string `json:"account_id"`
	ShortName string `json:"short_name"`
	ShopName  string `json:"shop_name"`
}

// ResultSelectAccount 选品结果表配账号数据表
type ResultSelectAccount struct {
	PublishId string `json:"publish_id"`
	ShortName string `json:"short_name"`
	ShopName  string `json:"shop_name"`
}

/*
 * ConnectAmazonCenter 连接亚马逊中心API
 */

//var AmazonCenterDb *gorm.DB

var amazonSelectResult YibaiAmazonPublishSkuSelectResult
var amazonSelectResults []YibaiAmazonPublishSkuSelectResult

func ConnectAmazonCenter() *gorm.DB {
	// TODO: connect to Amazon Center API
	res := ConfigDbConnUrl("")
	// 连接数据库 本地ck库数据连接
	db := GetDbConByGorm(res["amazonCenter"], "mysql", "yibai_sale_center_amazon.")
	return db
}

func TestQuery(publishId string) YibaiAmazonPublishSkuSelectResult {
	// TODO: test query
	db := ConnectAmazonCenter()
	db.Where("publish_id =?", publishId).Find(&amazonSelectResult)
	return amazonSelectResult

}

func SelectSelectResultWithAccountByPublishId(accountId int) []ResultSelectAccount {
	// TODO: select select result with account by publish id
	db := ConnectAmazonCenter()
	var rs []ResultSelectAccount
	// 定义查询Select字段
	tA := "yibai_amazon_publish_sku_select_result a"
	selectQuery := "a.publish_id,  if(b.short_name is null, '', b.short_name) short_name, if(b.shop_name is null, '', b.shop_name) shop_name"
	joinQuery1 := "left join yibai_sale_center_system.yibai_system_account b on b.id = a.account_id"
	whereQuery := "a.account_id =? "
	whereArgs := accountId
	// 执行查询 limit 10
	db.Limit(10).Table(tA).Select(selectQuery).Joins(joinQuery1).Where(whereQuery, whereArgs).Find(&rs)
	return rs

}
