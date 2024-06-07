package main

import (
	"GoLearn/common/models"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type A struct {
	AW string `json:"cc"`
}

// Scan 取出来的时候的数据
func (c *A) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), c)
}

// Value 入库的数据
func (c A) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

type B struct {
	models.Model
	FA    *A     `json:"fa"`
	FI    uint64 `json:"fid"`
	UserI uint   `json:"user_id"`
}

func main() {
	username := "root"    //账号
	password := "root"    //密码
	host := "172.20.64.1" //数据库地址，可以是Ip或者域名
	port := 3008          //数据库端口
	Dbname := "Gorm"      //数据库名
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
	// 连接成功
	fmt.Println("Yes")
	db.AutoMigrate(&B{})
	img := B{
		FA: &A{
			AW: "worksing",
		},
		FI:    21,
		UserI: 1,
	}
	db.Create(&img) // 注意这里传入的是地址
}
