package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

// User 结构体定义用户数据模型
// 使用 `bson` 标签来指定 MongoDB 中的字段名
type User struct {
	ID        int       `bson:"id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Age       int       `bson:"age"`
	Interests []string  `bson:"interests"`
	CreatedAt time.Time `bson:"created_at"`
}
type GameUser struct {
	ID   int    `bson:"id"`
	Name string `bson:"name"`
	Sex  int    `bson:"sex"`
	NOO  string `bson:"noo"`
}

// randomString 生成指定长度的随机字符串
// 用于生成随机兴趣
func randomString(length int) string {
	// 定义可用字符集
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		// 随机选择字符集中的字符
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	// 设置MongoDB连接的上下文和超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // 确保在函数结束时取消上下文

	// 连接到MongoDB
	// 这里使用的是本地MongoDB，端口为默认的27017
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(
		options.Credential{
			Username: "root",
			Password: "root",
		}))
	if err != nil {
		log.Fatal(err)
	}

	// 确保在函数结束时断开连接
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// 选择数据库和集合
	database := client.Database("testdb")
	collection := database.Collection("users")
	//database := client.Database("game")
	//collection := database.Collection("users")

	//insert(collection)
	find(collection)
	//find_30_60(collection)
	//find_game(collection)
}

func insert(collection *mongo.Collection) {
	// 生成200条用户数据
	var users []interface{}
	for i := 1; i <= 200; i++ {
		user := User{
			ID:        i,
			Name:      fmt.Sprintf("User%d", i),
			Email:     fmt.Sprintf("user%d@example.com", i),
			Age:       rand.Intn(50) + 18,                                          // 随机年龄18-67
			Interests: []string{randomString(5), randomString(5), randomString(5)}, // 生成3个随机兴趣
			CreatedAt: time.Now(),
		}
		users = append(users, user)
	}

	// 批量插入数据
	insertResult, err := collection.InsertMany(context.Background(), users)
	if err != nil {
		log.Fatal(err)
	}

	// 打印插入的文档数量
	fmt.Printf("Inserted %v documents\n", len(insertResult.InsertedIDs))

	// 验证插入：计算集合中的文档数量
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Collection has %v documents\n", count)
}
func find(collection *mongo.Collection) {
	// 查询 id 为 30 的数据
	filter := bson.M{"id": 30}
	var result User

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found with id 30")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Found document with id 30: %+v\n", result)
	}
}
func find_30_60(collection *mongo.Collection) {
	// 查询 id 在 30 到 60 之间的数据
	filter := bson.M{
		"id": bson.M{
			"$gte": 30,
			"$lte": 60,
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// 遍历查询结果
	var results []User
	for cursor.Next(context.Background()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// 打印结果
	fmt.Printf("Found %d documents with id between 30 and 60:\n", len(results))
	for _, user := range results {
		fmt.Printf("User: %+v\n", user)
	}
}
func find_game(collection *mongo.Collection) {
	// 查询 id 为 30 的数据
	filter := bson.M{"name": "thj"}
	var result GameUser

	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document found with name thj")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Found document with name thj: %+v\n", result)
	}
}
