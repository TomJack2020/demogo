package walmart

import (
	"demogo/src/tools"

	"gorm.io/gorm"
)

type YibaiWalmartSkuPublishTitleLog struct {
	TitleId      int    `gorm:"column:title_id"`
	Sku          string `gorm:"column:sku"`
	UseCount     int    `gorm:"column:use_count"`
	SiteCode     string `gorm:"column:site_code"`
	LanguageCode string `gorm:"column:language_code"`
}

// localMyDb  localCkDbUrl   amazonCenter

var (
	res = tools.ConfigDbConnUrl("") // 连接数据库 本地ck库数据连接
)

/*
SynAmazonTitleLog 同步亚马逊标题使用情况到本地ck库
*/
func SynWalmartTitleLog(startNum int, endNum int) {

	// 51288  186986015
	dbWalmartCenter := tools.GetDbConByGorm(res["walmartCenter"], "mysql", "yibai_sale_center_walmart.") // 连接数据库 亚马逊中心库数据连接

	// 查询标题使用情况
	var rs []YibaiWalmartSkuPublishTitleLog // 定义结构体数组 名字为数据表名 驼峰原则
	dbWalmartCenter.Model(&YibaiWalmartSkuPublishTitleLog{}).Select("title_id,sku,use_count,site_code,language_code").Where(
		"language_code = 'en' and site_code = 'US' and id between ? and ?", startNum, endNum).Find(&rs)

	// 同步数据库数据处理
	res := tools.ConfigDbConnUrl("")                                              // 连接数据库 本地ck库数据连接
	ddLocalCk := tools.GetDbConByGorm(res["localCkDbUrl"], "clickhouse", "imdb.") // 连接本地ck库 存储到imdb
	ddLocalCk.AutoMigrate(&YibaiWalmartSkuPublishTitleLog{})                      // 连接数据库 亚马逊中心库数据连接 建表语句 如果存在自动跳过 字段少的会新增 数据不会删除

	// fmt.Println(len(rs))
	// 同步数据
	WalmartInsertData(ddLocalCk, rs, len(rs))

}

// InsertData 批量插入数据 分批次插入 大于10000的 分10000每次插入防止内存溢出
func WalmartInsertData(db *gorm.DB, data []YibaiWalmartSkuPublishTitleLog, num int) {

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
