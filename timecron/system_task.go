package timecron

import (
	"log"
	"timecron/config"
	"timecron/request"
)

/*
系统任务
*/
var SystemTask = []TaskInfo{
	{
		Name:   "每周获取一次公告",
		Time:   "@every 20s",
		Type:   "4",
		Exec:   "http://baidu.com",
		System: true,
		Isrun:  "1",
	},
	{
		Name:   "每周定时检测更新版本",
		Time:   "",
		Type:   "",
		Exec:   "",
		System: true,
		Isrun:  "1",
	},
	{
		Name:   "定时清理日志或者文件",
		Time:   "",
		Type:   "",
		Exec:   "",
		System: true,
		Isrun:  "1",
	},
	{
		Name:   "定时检测系统状态",
		Time:   "",
		Type:   "",
		Exec:   "",
		System: true,
		Isrun:  "2",
	},
}

/*
获取版本号
*/
func GetVersion() {
	data, err := request.Get("https://cron.navjs.cn/cron/version?v="+config.Version, nil)
	if err != nil {
		return
	}

	log.Println(string(data))
}
