package main

import (
	"encoding/json"
	"fmt"
)

// Person 结构体代表一个拥有名字和年龄的人。
type Person struct {
	FirstName string `json:"firstName,int"` // JSON 标签用于指定 JSON 中的字段名称
	LastName  string `json:"lastName"`      // 如果没有提供 json 标签，默认使用字段名
	Age       int    `json:"-"`             // 这个字段在序列化时将被忽略
	Email     string `json:"email"`         // 指定 JSON 中的字段名称
}

func main() {
	// 创建一个 Person 实例
	p := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Email:     "john.doe@example.com",
	}

	// 将 Person 实例序列化为 JSON 字符串
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// 打印序列化后的 JSON 字符串
	fmt.Println(string(jsonData))
}
