package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"timecron/config"
	serve "timecron/gin"
	mylog "timecron/log"
	"timecron/timecron"
)

func init() {
	fmt.Println("DEBUG: main.go - init() started")
	if runtime.GOOS == "linux" { //windows上设置时区会报错,不设置也会正常显示,linux日志时间会差8小时
		TIME_LOCATION, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			log.Printf("time时区设置失败")
			panic(err)
		}
		time.Local = TIME_LOCATION // Re-enabled for now, assuming tzdata works
	}
	fmt.Println("DEBUG: main.go - init() finished")
}

func main() {
	fmt.Println("DEBUG: main.go - main() started")
	// cli.InitFlag()
	//初始化日志文件
	fmt.Println("DEBUG: main.go - Calling mylog.LogInit()")
	_, Writer := mylog.LogInit("main.log")
	fmt.Println("DEBUG: main.go - mylog.LogInit() returned")
	log.SetOutput(Writer) // 设置默认logger

	fmt.Println("DEBUG: main.go - Calling config.ReadConfigFileToJson()")
	cfg, err := config.ReadConfigFileToJson()
	fmt.Println("DEBUG: main.go - config.ReadConfigFileToJson() returned")
	if err != nil {
		log.Println("读取配置文件出错")
		fmt.Printf("DEBUG: main.go - Error reading config: %v\n", err)
		return
	}
	fmt.Println("DEBUG: main.go - Config loaded, proceeding to start services")
	fmt.Println(time.Now())
	log.Println("系统main启动")

	//初始化web服务 传递端口
	go serve.InitApi(cfg, nil)
	//初始化定时任务
	timecron.CronInit(cfg)
}
