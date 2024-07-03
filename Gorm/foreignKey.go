package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// User 有一张 CreditCard，UserID 是外键
//
//	type User struct {
//		gorm.Model
//		CreditCard *CreditCard
//	}
//
//	type CreditCard struct {
//		gorm.Model
//		Number string
//		UserID uint
//	}
type User struct {
	gorm.Model
	username string
}
type TestInodb struct {
	ID   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
}

func main() {
	var char = flag.String("f", "A", "")
	var val = flag.Int("num", 100000, "")
	flag.Parse()

	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	username := "root"    //账号
	password := "root"    //密码
	host := "172.26.80.1" //数据库地址，可以是Ip或者域名
	port := 3308          //数据库端口
	Dbname := "Java"      //数据库名
	timeout := "10s"      //连接超时，10秒
	/*
		docker run --name test-mysql -v "F:\docker_vh\test-mysql:/var/lib/mysql" -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=Gorm -p 3008:3306 mysql
	*/
	// root:root@tcp(127.0.0.1:3306)/Gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	db.AutoMigrate(&TestInodb{})
	//fmt.Println("Yes")
	//fmt.Println("%v", db)
	fmt.Println("Start ", time.Now())
	tx := db.Begin()
	// 假设有一个 User 模型
	for i := 0; i < *val; i++ {
		result := tx.Create(&TestInodb{
			Name: *char,
		})
		if i == 1 {
			var a TestInodb
			a.Name = "FALSE"
			db.Find(&a)
			fmt.Println(a)
		}
		if result.Error != nil {
			// 如果有错误发生，回滚事务
			tx.Rollback()
			fmt.Println("False")
			return
		}
	}
	if *char == "b" {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	s := fmt.Sprintf("End ", time.Now())
	fmt.Println(s)
}

/*
go run .\foreignKey.go -f b -num 50000
go run .\foreignKey.go -f a -num 10000
*/
