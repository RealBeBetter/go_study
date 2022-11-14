/**
 * @author: Real
 * Date: 2022/11/15 0:45
 */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 导入驱动，不用使用
	"time"
)

type User struct {
	Id       int
	Username string
	Age      int
	Address  string
	Phone    string
}

func main() {
	db, _ := ConnectMysql()
	// 插入数据
	if _, err := db.Exec("INSERT INTO tb_user(username, age, address, phone) VALUES(?, ?, ?, ?)",
		"Real", 22, "ZH, Hunan", "12345678910"); err != nil {
		fmt.Println(err)
		return
	}
	// 使用Query查询数据
	rows, err := db.Query("SELECT * FROM tb_user")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("使用Query:\n", "ID", "|", "Username", "|", "Age", "|", "Address", "|", "Phone")
	for rows.Next() {
		var (
			Id       int
			Username string
			Age      int
			Address  string
			Phone    string
		)
		if err := rows.Scan(&Id, &Username, &Age, &Address, &Phone); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(" ", Id, Username, Age, Address, Phone)
	}
	// 使用
	var (
		Id       int
		Username string
		Age      int
		Address  string
		Phone    string
	)
	_ = db.QueryRow("SELECT * FROM tb_user WHERE id = 1").Scan(&Id, &Username, &Age, &Address, &Phone)
	fmt.Println("使用QueryRow:\n", Id, Username, Age, Address, Phone)
}

type MySqlConfig struct {
	UserName string
	Password string
	IP       string
	Port     int
	Database string
}

func parseMysqlConfig(m MySqlConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		m.UserName, m.Password, m.IP, m.Port, m.Database)
}

func ConnectMysql() (*sql.DB, error) {
	// Mysql 数据库配置
	mysql := MySqlConfig{
		UserName: "root",
		Password: "123456",
		IP:       "127.0.0.1",
		Port:     3306,
		Database: "test",
	}
	db, err := sql.Open("mysql", parseMysqlConfig(mysql))
	if err != nil {
		return db, err
	}
	// 设置连接的生命周期，超过这个时间他会被Close，一般不会设置这个
	db.SetConnMaxLifetime(time.Hour * 24)
	// 设置最大空闲连接池的大小，默认为2
	db.SetMaxIdleConns(100)
	// 设置最大连接数量
	db.SetMaxOpenConns(20)
	// 通过 ping 测试数据库是否连接成功
	if err := db.Ping(); err != nil {
		return db, err
	}
	fmt.Println("Connect Mysql Success....")
	return db, nil
}
