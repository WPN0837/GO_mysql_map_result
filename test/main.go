package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"test/Database"
)

const (
	USERNAME = "mysqlsa"
	PASSWORD = "Lansi123"
	NETWORK  = "tcp"
	SERVER   = "rm-uf6sjzuqdryz62fs7lo.mysql.rds.aliyuncs.com"
	PORT     = 3306
	DATABASE = "lansi_sy_evaluate"
)

//func queryOne(DB *sql.DB) {
//	fmt.Println("query times:")
//	rows, err := DB.Query("select bank,id from sy_bank")
//	if err != nil {
//		fmt.Printf("error1 %v", err)
//		return
//	}
//	data := []map[string]string{}
//	for rows.Next() {
//		var bank string
//		var id string
//		err = rows.Scan(&bank, &id)
//		if err != nil {
//			fmt.Printf("Scan failed,err:%v", err)
//			return
//		}
//		b := map[string]string{"bank": bank, "id": id}
//		data = append(data, b)
//	}
//	r, _ := json.Marshal(data)
//	fmt.Println(string(r))
//}
//func con() {
//	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
//	DB, err := sql.Open("mysql", dsn)
//	if err != nil {
//		fmt.Printf("Open mysql failed,err:%v\n", err)
//		return
//	}
//	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
//	DB.SetMaxOpenConns(100)                  //设置最大连接数
//	DB.SetMaxIdleConns(16)                   //设置闲置连接数
//	queryOne(DB)
//	DB.Close()
//}

func main() {
	//pool := Database.DatabaseConnectPool{USERNAME: USERNAME, PASSWORD: PASSWORD, NETWORK: NETWORK, SERVER: SERVER, PORT: PORT, DATABASE: DATABASE}
	//db, err := pool.Get()
	//if err != nil {
	//	fmt.Println("Error:%s", err)
	//}
	//now := ""
	//err = Database.QueryOne(db, "select now()", &now)
	//if err != nil {
	//	fmt.Println("Error:%s", err)
	//}
	//fmt.Println(now)

	//con()

	//sql格式化参数
	args := []interface{}{}
	//查询结果字段名称  传入Database.Query时的顺序必须要和sql查询字段顺序一致
	result_names := []interface{}{"nickname", "inserttime"}

	// 批量查询
	var res, err = Database.Query("select nickname,inserttime from sy_account", args, result_names)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("批量查询")
		for _, r := range *res {
			fmt.Println(r)
		}
	}
	// 单条查询
	args = []interface{}{25, "nick_admin"}
	res, err = Database.Query("select nickname,inserttime from sy_account where id=? and nickname=? limit 1", args, result_names)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("单条查询")
		fmt.Println((*res)[0])
	}

}
