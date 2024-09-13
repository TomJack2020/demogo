package main

import (
	"demogo/tools"

	"gorm.io/gorm"
)

type Result struct {
	PublishId string `json:"publish_id"`
	ShortName string `json:"short_name"`
	ShopName  string `json:"shop_name"`
}

// localMyDb  localCkDbUrl   amazonCenter

func main() {

	db := tools.ConnectAmazonCenter()
	var rs []Result
	// 定义查询Select字段
	tA := "yibai_amazon_publish_sku_select_result a"
	//selectQuery := "a.publish_id,  if(b.short_name is null, '', b.short_name) short_name, if(b.shop_name is null, '', b.shop_name) shop_name"
	sql := tools.ReadSql("../sql/test.sql")
	//fmt.Println(sql)
	selectQuery := sql
	joinQuery1 := "left join yibai_sale_center_system.yibai_system_account b on b.id = a.account_id"
	whereQuery := "a.account_id =?"
	whereArgs := 7638

	// 执行查询
	//db.Limit(100).Table(tA).Select(selectQuery).Joins(joinQuery1).Where(whereQuery, whereArgs).Find(&rs)
	// 打印map类型
	var productMap []map[string]interface{}
	db.Limit(100).Model(&Result{}).Table(tA).Select(selectQuery).Joins(joinQuery1).Where(whereQuery, whereArgs).Find(&productMap)

	//fmt.Println(productMap)

	//tools.TestQuery()
	//tools.TestQuerys()

	// 同步数据库数据处理
	res := tools.ConfigDbConnUrl("")
	// 连接数据库 本地ck库数据连接
	ddLocalCk := tools.GetDbConByGorm(res["localCkDbUrl"], "clickhouse", "imdb.")

	// 连接数据库 亚马逊中心库数据连接
	ddLocalCk.AutoMigrate(&Result{})

	// 插入前清空表数据
	ddLocalCk.Where("1=1").Delete(&Result{})

	//[]map[string]interface{} 转化为[]Result  如果不是为了打印直接可以用[]Result 直接插入
	for _, item := range productMap {
		rs = append(rs, Result{
			PublishId: item["publish_id"].(string),
			ShortName: item["short_name"].(string),
			ShopName:  item["shop_name"].(string),
		})
	}
	// 同步数据
	InsertData(ddLocalCk, rs, len(rs))

}

func InsertData(db *gorm.DB, data []Result, num int) {

	for i := 0; i < num; i += 10000 {
		start := i
		end := i + 10000
		// 防止越界
		if end > num {
			end = num
		}
		ins := data[start:end]
		db.CreateInBatches(ins, len(ins))
	}
}
