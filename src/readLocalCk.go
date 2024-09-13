package main

//type DbUserInfo struct {
//	Id       int    `gorm:"column:id;primary_key"`
//	Dbname   string `gorm:"column:dbname"`
//	Dbhost   string `gorm:"column:dbhost"`
//	Username string `gorm:"column:username"`
//	Password string `gorm:"column:password"`
//	Dbuse    string `gorm:"column:dbuse"`
//}

// localMyDb  localCkDbUrl   amazonCenter
func main() {

	//res := tools.ConfigDbConnUrl("")
	//// 连接数据库 本地ck库数据连接
	//db := tools.GetDbConByGorm(res["localCkDbUrl"], "clickhouse", "imdb.")
	//db.AutoMigrate(&DbUserInfo{})

	//users := []DbUserInfo{
	//	{Id: 1, Dbname: "test1", Dbhost: "127.0.0.1", Username: "lisi", Password: "123456", Dbuse: "ck"},
	//	{Id: 2, Dbname: "test2", Dbhost: "127.0.0.1", Username: "wangwu", Password: "123456", Dbuse: "ck"},
	//}
	//// 插入
	//db.Create(&users)

	//db.Exec("truncate imdb.db_user_info")

	// 批量插入构建

	//for i := 0; i < 20; i++ {
	//	startNum := i * num
	//	endNum := (i + 1) * num
	//	InsertData(db, startNum, endNum)
	//}
	//num := 125000
	//var insertUsers []DbUserInfo
	//// 循环插入100000条数据
	//for i := 0; i < num; i++ {
	//	insertUsers = append(insertUsers, DbUserInfo{
	//		Id:       i + 1,
	//		Dbname:   "test1" + strconv.Itoa(i),
	//		Dbhost:   "127.0.0.1" + strconv.Itoa(i),
	//		Username: "lisi" + strconv.Itoa(i),
	//		Password: "123456" + strconv.Itoa(i),
	//		Dbuse:    "ck" + strconv.Itoa(i),
	//	})
	//}

	// 批量插入10000条数据
	//InsertData(db, insertUsers, num)

}

//func InsertData(db *gorm.DB, data []DbUserInfo, num int) {
//	for i := 0; i < num; i += 10000 {
//		start := i
//		end := i + 10000
//		// 防止越界
//		if end > num {
//			end = num
//		}
//		ins := data[start:end]
//		db.CreateInBatches(ins, len(ins))
//	}
//
//}
