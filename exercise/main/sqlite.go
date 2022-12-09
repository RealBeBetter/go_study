/**
 * @author: Real
 * Date: 2022/11/17 21:50
 */
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"reflect"
	"time"
)

const userCount = 10000
const orderCount = 100000

func insert(db *sql.DB, id int, uid int, weight float64, createdAt string) int64 {
	stmt, err := db.Prepare("INSERT INTO tb_order (id, uid, weight, createdAt) values(?, ?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec(id, uid, weight, createdAt)
	checkErr(err)
	insertId, err := res.LastInsertId()
	checkErr(err)
	// 获取最后插入的 ID 值
	return insertId
}

func insertOrder(db *sql.DB, id int, uid int, weight float64, createdAt string) int64 {
	stmt, err := db.Prepare("INSERT INTO tb_order(id, uid, weight, createdAt) values(?, ?, ?, ?)")
	if err != nil {
		fmt.Println("prepare sql error , error is", err)
		return 1
	}

	res, err := stmt.Exec(id, uid, weight, createdAt)
	checkErr(err)
	insertId, err := res.LastInsertId()
	checkErr(err)
	// 获取最后插入的 ID 值
	return insertId
}

func updateWeight(db *sql.DB, weight float64, id int) int64 {
	stmt, err := db.Prepare("update tb_order set weight = ? where id = ?")
	checkErr(err)
	res, err := stmt.Exec(weight, id)
	checkErr(err)
	// 执行 sql 语句
	affect, err := res.RowsAffected()
	checkErr(err)
	return affect
}

func selectOrder(db *sql.DB, order *Order) {
	rows, err := db.Query("SELECT * FROM tb_order")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&order.id, &order.uid, &order.weight, &order.createdAt)
		checkErr(err)
	}

}

func selectOrdersByUid(db *sql.DB, uid int) []Order {
	sqlStatement := `SELECT * FROM tb_order WHERE uid = ?`
	_, err := db.Prepare(sqlStatement)
	checkErr(err)
	var orders []Order
	rows, err := db.Query(sqlStatement, uid)
	// rows, err := stmt.Query(sqlStatement, uid)
	if err != nil {
		fmt.Printf("query failed, err: %v\n", err)
	}
	count := 0
	var order Order
	for rows.Next() {
		err = rows.Scan(&order.id, &order.uid, &order.weight, &order.createdAt)
		fmt.Println("select order success, order is", order)
		checkErr(err)
		orders = append(orders, order)
		count++
	}
	return orders
}

func deleteOrder(db *sql.DB, id int) int64 {
	stmt, err := db.Prepare("delete from tb_order where id=?")
	checkErr(err)
	res, err1 := stmt.Exec(id)
	checkErr(err1)
	affect, err2 := res.RowsAffected()
	checkErr(err2)
	return affect

}

func execSql() {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	defer db.Close()
	checkErr(err)
	// 执行数据表，添加数据表
	exec, err := db.Exec(sqlTable)
	if err != nil {
		fmt.Println("create database failed, error is", err)
		return
	}
	fmt.Println(exec)

	// 2生成10000个用户id(1-10000)
	// 3生成100000条订单记录并插入数据库中
	// 	1)订单id顺序生成
	// 	2)用户id从上一步生成的用户id中随机选择
	// 	3)计费重量取整数最大重量100KG
	// 	4)在所有订单中，货物重量的分布权重大致为1/W，例如
	//		2KG货物订单的数量与8KG货物订单的数量比例约为(1/2):(1/8)=4:1
	// 5命令行输入任意1个用户id，则计算出此用户的总费用并显示

	countNum := 0
	for i := 1; i <= 100; i++ {
		countNum += i * i
	}

	count := 0
	var distribute []int = make([]int, countNum)
	for i := 1; i <= 100; i++ {
		for j := 1; j <= i; j++ {
			distribute[count] = i
			count++
		}
	}

	for i := 1; i <= orderCount; i++ {
		uid := rand.Intn(userCount) + 1
		// weight := rand.Intn(100) + 1
		weight := distribute[rand.Intn(count)]
		timeStr := time.Now().Format("2006-01-02 15:04:05")
		order := insertOrder(db, i, uid, float64(weight), timeStr)
		fmt.Println("order is", order)
	}

	db.Close()

}

func calUserDeliveryFee(uid int) {
	db, err := sql.Open("sqlite3", "./data.db")
	checkErr(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("close database connection failed, error is", err)
		}
	}(db)
	checkErr(err)
	fmt.Println("ready to select order, uid is", uid)
	orders := selectOrdersByUid(db, uid)
	fmt.Printf("select result is %+v\n", orders)
	var totalDeliveryFee float64
	var userId int
	if orders != nil && len(orders) != 0 {
		fmt.Println("the count of order is", len(orders))
		for index, order := range orders {
			// 计算费用
			fmt.Printf("the %dth th order is %+v\n", index+1, order)
			weight := order.weight
			user := order.uid
			fmt.Printf("weight is %f, typeof is %s\n", weight, reflect.TypeOf(weight))
			fee := CalDeliveryFee(weight)
			userId = user
			totalDeliveryFee += fee
		}
	}
	fmt.Println("the delivery fee of user", userId, "is", totalDeliveryFee)
	err = db.Close()
	if err != nil {
		fmt.Println("close database connection failed, error is", err)
		return
	}
}

var sqlTable = `
DROP TABLE if exists "tb_order";
CREATE TABLE if not exists "tb_order" (
    "id" INTEGER PRIMARY KEY AUTOINCREMENT,
    "uid" INTEGER NULL,
	"weight" DOUBLE,
	"createdAt" TIMESTAMP default (datetime('now', 'localtime'))
);
CREATE INDEX id_idx ON tb_order (id);
`

type Order struct {
	id        int       `json:"id"`
	uid       int       `json:"uid"`
	weight    float64   `json:"weight"`
	createdAt time.Time `json:"created_at"`
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("this operation has error, error is", err)
	}
}
