package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	//f, err := os.Create("profile.prof")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//
	//if err := pprof.StartCPUProfile(f); err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer pprof.StopCPUProfile()
	//
	// 启动 pprof HTTP 服务器
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil)) // 使用 nil 即可，因为 pprof 已经注册了其处理函数
	}()

	for {
		fmt.Println("Running...")
		time.Sleep(1 * time.Second) // 每秒打印一次，避免 CPU 占用过高
	}
}
