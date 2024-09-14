package draft

import "fmt"

//type DbUserInfo struct {
//	Id       int    `gorm:"column:id;primary_key"`
//	Dbname   string `gorm:"column:dbname"`
//	Dbhost   string `gorm:"column:dbhost"`
//	Username string `gorm:"column:username"`
//	Password string `gorm:"column:password"`
//	Dbuse    string `gorm:"column:dbuse"`
//}

// localMyDb  localCkDbUrl   amazonCenter
func readLocalCk() {
	fmt.Println("readLocalCk")
	//res := tools.ConfigDbConnUrl("")
	//// 连接数据库 本地ck库数据连接
	//db := tools.GetDbConByGorm(res["localCkDbUrl"], "clickhouse", "imdb.")
	//db.AutoMigrate(&DbUserInfo{})

}
