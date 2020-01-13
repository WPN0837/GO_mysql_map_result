package Database

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

const (
	USERNAME = "mysqlsa"
	PASSWORD = "Lansi123"
	NETWORK  = "tcp"
	SERVER   = "rm-uf6sjzuqdryz62fs7lo.mysql.rds.aliyuncs.com"
	PORT     = 3306
	DATABASE = "lansi_sy_evaluate"
)

var pool = DatabaseConnectPool{USERNAME: USERNAME, PASSWORD: PASSWORD, NETWORK: NETWORK, SERVER: SERVER, PORT: PORT, DATABASE: DATABASE}

type DatabaseConnectPool struct {
	USERNAME string
	PASSWORD string
	NETWORK  string
	SERVER   string
	PORT     int
	DATABASE string
	con      *DB
}

func init() {
	/*
		注册mysql驱动
	*/
	Register("mysql", mysql.MySQLDriver{})
}
func (c *DatabaseConnectPool) connect() error {
	ServerInfo := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", c.USERNAME, c.PASSWORD, c.NETWORK, c.SERVER, c.PORT, c.DATABASE)
	DB, err := Open("mysql", ServerInfo)
	if err != nil {
		fmt.Println(111)
		return err
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
	c.con = DB
	return nil
}
func (c DatabaseConnectPool) Get() (*DB, error) {
	if c.con != nil {
		return c.con, nil
	}
	err := c.connect()
	return c.con, err
}
func (c DatabaseConnectPool) Close() {
	if c.con != nil {
		c.con.Close()
	}
}
func Query(Sql string, args []interface{}, dest []interface{}) (*[]map[string]interface{}, error) {
	/*
		@params
		Sql: 被执行的sql
		args: Sql的参数  没有参数传空interface{}数组 顺序和类型一定要和sql中一致
		dest: 返回map对象中key值地址
		@result
		*[]map[string]interface{}: 查询结果map对象数组地址
		error: 错误信息,有错误返回错误信息   没有错误返回nil
	*/
	results := []map[string]interface{}{}
	var result *map[string]interface{}
	dest_p := []interface{}{}
	for i := range dest {
		dest_p = append(dest_p, &dest[i])
	}

	db, PoolErr := pool.Get()
	if PoolErr != nil {
		return &results, PoolErr
	}
	rows, err := db.Query(Sql, args...)
	if err != nil {
		return &results, err
	}
	for rows.Next() {
		result, err = rows.Scan(dest_p...)
		if err != nil {
			return &results, err
		}
		results = append(results, *result)
	}
	return &results, err
}
