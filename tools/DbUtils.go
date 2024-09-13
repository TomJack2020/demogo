package tools

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
	"os"
)

/*
Query 查询数据
参数：
driver 数据库驱动
dsn 数据库连接字符串
parms 查询参数
返回：
conn 数据库连接对象
*/

func ConnectDb(driver string, dsn string) *sqlx.DB {

	// dsn := "tcp://localhost:9000?username=default&password=test123&database=imdb&block_size=4096"
	connect, err := sqlx.Connect(driver, dsn)
	// 如果错误则打印错误信息
	if err != nil {
		fmt.Printf("clickhouse open err %s", err.Error())
		return nil
	}
	//返回连接对象
	return connect
}

/*
GetRows 获取查询结果集
参数：
connect 数据库连接对象
columns 要查询的字段名数组
返回：
*sqlx.Rows 查询结果集
*/

func GetRows(connect *sqlx.DB, sql string) *sqlx.Rows {

	// sql := "select %s from test_java.lookup limit 1, 10"
	// 打印sql语句
	// fmt.Println(fmt.Sprintf(sql,strings.Join(columns, ", ")))
	// rows, err := connect.Queryx(fmt.Sprintf(fmt.Sprintf(sql,strings.Join(columns, ", ")))) //注意:使用 sqlx 的 Queryx 方法
	rows, err := connect.Queryx(sql) //注意:使用 sqlx 的 Queryx 方法
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rows

}

/*
GetOneRow使用gorm连接数据库
*/

func GetDbConByGorm(dsn string, dbType string, PrefixName string) *gorm.DB {
	// "root:test123@tcp(localhost:3306)/test_java?charset=utf8&parseTime=True&loc=Local"

	switch dbType {
	case "mysql":
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   PrefixName, // 表名前缀，`Article` 的表名应该是 `it_articles`
				SingularTable: true,       // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`,
			},
		})
		if err != nil {
			fmt.Println(err.Error())
			panic("failed to connect database")
		}

		return db
	case "clickhouse":
		// clickhouse://default:test123@localhost:9000?database=imdb
		db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   PrefixName, // 表名前缀，`Article` 的表名应该是 `it_articles`
				SingularTable: true,       // 使用单数表名，启用该选项，此时，`Article` 的表名应该是 `it_article`,
			},
		})
		if err != nil {
			fmt.Println("failed to connect database", err.Error())
			panic("failed to connect database")
		}
		return db

	default:
		fmt.Println("not support driver")
		return nil

	}

}

/*
ReadSql 读取sql文件内容
参数：
path sql文件路径
返回：
string sql文件内容
*/
func ReadSql(path string) string {

	// 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return "error:" + err.Error()
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	// 读取文件内容
	content, _ := ioutil.ReadAll(file)
	return string(content)
}
