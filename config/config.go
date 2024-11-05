package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/tidwall/gjson"
)

var Version = "1.1.0"

// 所有结构数据
type Tconfig struct{}

// 读取json配置文件转结构体
func ReadConfigFile() (*Tconfig, error) {
	//配置文件地址在此固定写死
	file := "./config.json"
	runpath, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(runpath)
	file = path.Join(dir, file)
	fmt.Println(file, "config")
	jsonByte, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("配置文件读取失败")
		return nil, err
	}
	var configObj Tconfig
	if err := json.Unmarshal(jsonByte, &configObj); err != nil {
		fmt.Println("反序列化配置文件错误>>", err)
		return nil, err
	}
	return &configObj, nil
}

// 将config文件读取到json字符串
func ReadConfigFileToJson() (gjson.Result, error) {
	//配置文件地址在此固定写死
	file := "./config.json"
	runpath, err := os.Executable()
	if err != nil {
		log.Println(err)
	}
	dir := filepath.Dir(runpath)
	file = path.Join(dir, file)
	jsonByte, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("配置文件读取失败")
		/* 配置文件不存在,创建json文件 */
		str := fmt.Sprintf(`{
			"name": "timecron",
			"username":"admin",
			"email":"xnkyn@qq.com",
			"password":"21232f297a57a5a743894a0e4a801fc3",
			"task": [
				{
				"createtime":"%s",
				"name": "测试任务",
				"type": "1",
				"exec": "pwd",
				"isrun": "2",
				"time": "@every 20s",
				"desc": "30秒执行一次ls命令任务"
			    }
			]
		  }`, time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
		err := WriteConfigFile(file, []byte(str))
		if err != nil {
			log.Println("配置文件创建失败")
			return gjson.Parse(""), err
		}
		log.Println("配置文件创建成功")
		return gjson.Parse(str), nil
	}

	result := gjson.Parse(string(jsonByte))

	// 这里我们得到一个gjson实例
	// 后面可以在任意位置重用它获取值
	// lastName := result.Get("database").String()
	// fmt.Println(lastName)
	return result, nil
}

// 写入json到config文件
func WriteConfigFile(filePth string, data []byte) error {
	f, err := os.Create(filePth)
	if err != nil {
		fmt.Println("config文件创建失败")
		return err
	} else {
		_, err = f.Write(data) //写入文件要字节类型[]byte(data)
		if err != nil {
			// 写入失败处理
			fmt.Println("config文件写入失败")
			return err
		}
	}
	defer f.Close()
	return nil
}
