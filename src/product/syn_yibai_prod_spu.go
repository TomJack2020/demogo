package product

import (
	"demogo/src/tools"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type YibaiProdSpu struct {
	Id                int    `json:"id"`
	Sku               string `json:"sku"`
	Spu               string `json:"spu"`
	TitleCn           string `json:"title_cn"`
	ProductStatus     int    `json:"product_status"`
	ProductIsMulti    int    `json:"product_is_multi"`
	ProductCategoryId int    `json:"product_category_id"`
}

// 定义初始变量
var (
	res = tools.ConfigDbConnUrl("") // 定义map 存放数据库连接信息
)

/*
SynYibaiProdSpu 同步亚马逊中心库商品spu数据到本地ck库
*/
func SynYibaiProdSpu() {

	dbConProduct := tools.GetDbConByGorm(res["productSystem"], "mysql", "yibai_prod_base.") // 连接数据库 亚马逊中心库数据连接
	ddLocalCk := tools.GetDbConByGorm(res["localCkDbUrl"], "clickhouse", "imdb.")           // 连接本地ck库 存储到imdb

	// 查询标题使用情况
	var rs []YibaiProdSpu // 定义结构体数组 名字为数据表名 驼峰原则
	dbConProduct.Select("id,spu,title_cn,product_status,product_is_multi,product_category_id").Where("end_time <> ''").Find(&rs)

	// 同步数据库数据处理
	// 连接数据库 本地ck库数据连接

	ddLocalCk.AutoMigrate(&YibaiProdSpu{}) // 连接数据库 亚马逊中心库数据连接 建表语句 如果存在自动跳过 字段少的会新增 数据不会删除

	// 同步数据
	InsertData(ddLocalCk, rs, len(rs))
}

// InsertData 批量插入数据 分批次插入 大于10000的 分10000每次插入防止内存溢出
func InsertData(db *gorm.DB, data []YibaiProdSpu, num int) {

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

func ReadSpu(limirNum int, OffsetNum int) []string {

	ddLocalCk := tools.GetDbConByGorm(res["localCkDbUrl"], "clickhouse", "imdb.") // 连接本地ck库 存储到imdb
	ddLocalCk.Logger = logger.Default.LogMode(logger.Silent)
	var rs []string // 定义结构体数组 名字为数据表名 驼峰原则
	ddLocalCk.Limit(10).Model(&YibaiProdSpu{}).Select("spu").Limit(limirNum).Offset(OffsetNum).Find(&rs)

	return rs

}
