package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"timecron/request"
)

func Test_writeFile(t *testing.T) {
	data, _ := request.Get("http://127.0.0.1:4523/m1/1960452-0-default/admin/index/index", nil)
	fmt.Println(123, "我执行了")
	// config.WriteConfigFile("./test2.json", data)
	fmt.Println(data)
	type res struct {
		Content int `json:"content"`
	}
	var configObj res
	if err := json.Unmarshal(data, &configObj); err != nil {
		fmt.Println("反序列化配置文件错误>>", err)
	}
	fmt.Println(configObj)
}
