package draft

import (
	"demogo/src/tools"
	"fmt"
	"log"
)

func conCk() {
	fmt.Println("Hello, world!")

	// 连接数据库clickhouse
	dsn := "tcp://localhost:9000?username=default&password=test123&database=imdb&block_size=4096"
	driver := "clickhouse"

	connect := tools.ConnectDb(driver, dsn)
	// 连接数据库
	// db := tools.ConnectDb(driver,dsn)
	// connect,err := sqlx.Connect("clickhouse", dsn)

	// if err != nil {
	//     fmt.Printf("clickhouse open err %s", err.Error())
	// }
	// defer func() {
	//     _ = connect.Close()
	// }()

	// 定义结
	type product struct {
		Id  int
		spu string
		sku string
	}

	// 执行sql查询
	sql := "select id, spu, sku from imdb.syn_yibai_product_info where id > 1000 limit 10000"

	rows, err := connect.Queryx(sql)
	//checkErr(err)

	//数据预处理写入
	tx, err := connect.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// 预处理sql
	stmt, err := tx.Prepare("insert into imdb.yibai_amazon_publish_sku_select_result_tt_sku (id,sku,title) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	// 循环读取数据 写入clickhouse
	for rows.Next() {
		var p product
		err := rows.Scan(&p.Id, &p.spu, &p.sku) // 获取对象product
		if err != nil {
			fmt.Println(err)
		} else {
			if _, err := stmt.Exec(p.Id, p.sku, p.spu); err != nil {
				log.Fatal(err)
			}

		}
	}
	_ = tx.Commit()
	fmt.Println("数据写入成功")

	// 关闭数据库连接
	rows.Close()
	db.Close()

}
