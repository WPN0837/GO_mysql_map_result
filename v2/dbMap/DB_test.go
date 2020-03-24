package dbMap_test

import (
	"dbMap"
	"fmt"
	"testing"
)

const (
	USERNAME = "mysqlsa"
	PASSWORD = "Lansi123"
	NETWORK  = "tcp"
	SERVER   = "rm-uf6sjzuqdryz62fs7lo.mysql.rds.aliyuncs.com"
	PORT     = "3306"
	DATABASE = "lansi_sy_evaluate"
)

func TestQuery(t *testing.T) {
	var pool = dbMap.ConnectPool{USERNAME: USERNAME, PASSWORD: PASSWORD, NETWORK: NETWORK, SERVER: SERVER, PORT: PORT, DATABASE: DATABASE}
	db, err := pool.Get()
	if err != nil {
		fmt.Println("数据库连接失败")
		fmt.Println(err)
		return
	}
	//sql格式化参数
	args := []interface{}{}
	//查询结果字段名称  传入Database.Query时的顺序必须要和sql查询字段顺序一致
	resultNames := []string{"nickname", "inserttime"}
	res, err := dbMap.Query(db, "select nickname,inserttime from sy_account", args, resultNames)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("批量查询")
		for _, r := range res {
			fmt.Println(r)
		}
	}

	args = []interface{}{25, "nick_admin"}
	res, err = dbMap.Query(db, "select nickname,inserttime from sy_account where id=? and nickname=? limit 1", args, resultNames)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("单条查询")
		fmt.Println((res)[0])
	}
}
func TestConnectPool_Get(t *testing.T) {
	var pool = dbMap.ConnectPool{USERNAME: USERNAME, PASSWORD: PASSWORD, NETWORK: NETWORK, SERVER: SERVER, PORT: PORT, DATABASE: DATABASE}
	_, err := pool.Get()
	if err != nil {
		fmt.Println("数据库连接失败")
		fmt.Println(err)
		return
	}
	pool.Close()
}