package dbMap

import (
	. "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//数据库类型
const driverName = "mysql"


type ConnectPool struct {
	USERNAME string
	PASSWORD string
	NETWORK  string
	SERVER   string
	PORT     string
	DATABASE string
	con      *DB
}

func NewConnectPool(username, password, host, port, database string) *ConnectPool {
	/**
	username
	返回一个连接池的地址
	 */
	return &ConnectPool{
		username, password, "tcp", host, port, database, nil,
	}
}
func (c *ConnectPool) connect() error {
	ServerInfo := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8", c.USERNAME, c.PASSWORD, c.NETWORK, c.SERVER, c.PORT, c.DATABASE)
	DB, err := Open(driverName, ServerInfo)
	if err != nil {
		return err
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
	c.con = DB
	return nil
}
func (c *ConnectPool) Get() (*DB, error) {
	/**
	返回一个数据库连接
	*/
	if c.con != nil {
		return c.con, nil
	}
	err := c.connect()
	return c.con, err
}
func (c *ConnectPool) Close() {
	/**
	关闭数据库连接
	*/
	if c.con != nil {
		_ = c.con.Close()
	}
}
func Query(db *DB, Sql string, args []interface{}, dest []string) ([]map[string]string, error) {
	/*
		数据库查询入口 只能返回map[string]string类型结果集
		@params
		Sql: 被执行的sql
		args: Sql的参数  没有参数传空interface{}数组 顺序和类型一定要和sql中一致
		dest: 返回结果集map中key的名称
		@result
		[]map[string]string: 查询结果map对象数组
		error: 错误信息,有错误返回错误信息   没有错误返回nil
	*/
	//定义返回的结果集
	var results []map[string]string
	//定义接收查询结果切片 interface{}类型
	destP := make([]interface{}, len(dest))
	//把每个结果初始化为*string
	for i := range dest {
		destP[i] = new(string)
	}
	//调用DB查询 得到游标rows
	rows, err := db.Query(Sql, args...)
	if err != nil {
		return results, err
	}
	//循环取出所有行
	for rows.Next() {
		//定义存储单行查询结果的map对象
		result := make(map[string]string)
		//取出当前行的查询结果到destP中
		err = rows.Scan(destP...)
		if err != nil {
			return results, err
		}
		//循环把destP中的结果按照dest的key存储到result中
		for i := range dest {
			//destP[i]中存储的是interface{}类型的string指针
			//destP[i].(*string)是string指针值
			//*destP[i].(*string)是string的值
			result[dest[i]] = *destP[i].(*string)
		}
		//把单行查询结果添加到results中
		results = append(results, result)
	}
	return results, err
}
