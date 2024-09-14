package draft

import (
	"demogo/tools"
	"fmt"
)

type DbUserInfo struct {
	Id       int
	Dbname   string
	Dbhost   string
	Username string
	Password string
	Dbuse    string
}

// localMyDb  localCkDbUrl   amazonCenter
func testgorm() {

	res := tools.ConfigDbConnUrl("")
	fmt.Println(res)
	// 连接数据库
	db := tools.GetDbConByGorm(res["localMyDb"], "mysql", "book.")
	fmt.Println(db)

	db.AutoMigrate(&DbUserInfo{})

	db.Exec("Truncate table db_user_info")

	users := []DbUserInfo{
		{Id: 1, Dbname: "test1", Dbhost: "127.0.0.1", Username: "lisi", Password: "123456", Dbuse: "ck"},
		{Id: 2, Dbname: "test2", Dbhost: "127.0.0.1", Username: "wangwu", Password: "123456", Dbuse: "ck"},
		{Id: 3, Dbname: "test2", Dbhost: "127.0.0.1", Username: "wangwu", Password: "123456", Dbuse: "ck"},
	}
	// 插入
	db.Create(&users)

	//关闭数据库连接
	db.Commit()

}
