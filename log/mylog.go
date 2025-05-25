package mylog

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"timecron/lib"
)

const logDir = "logs"

func LogInit(name string) (*log.Logger, io.Writer) {
	if name == "" { //没有名称时候,返回空日志
		return log.New(os.Stdout, "", 0), os.Stdout
	}
	logPath := filepath.Join(logDir, name)

	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		log.Fatalf("创建日志目录失败: %v", err)
	}

	_, err = lib.PathFileExists(logPath)
	if err != nil {
		f, err := os.Create(logPath)
		if err != nil {
			log.Println(err.Error())
			return log.New(os.Stdout, "LOG_FILE_CREATE_ERROR: ", log.LstdFlags), os.Stdout
		}
		log.Println(name, "日志文件创建成功")
		f.Close()
	}

	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("打开日志文件失败: %v", err)
	}

	multiWriter := io.MultiWriter(file, os.Stdout)

	logger := log.New(multiWriter, "", log.LstdFlags)

	return logger, multiWriter

}
