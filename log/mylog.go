package mylog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	// "timecron/lib" // Temporarily remove dependency on lib.PathFileExists
)

const logDir = "logs"

func LogInit(name string) (*log.Logger, *os.File) {
	fmt.Printf("DEBUG: LogInit called for log file name: '%s'\n", name)
	if name == "" { //没有名称时候,返回空日志
		fmt.Println("DEBUG: LogInit - name is empty, returning stdout logger")
		return log.New(os.Stdout, "", 0), nil
	}
	logPath := filepath.Join(logDir, name)
	fmt.Printf("DEBUG: LogInit - Resolved logPath: '%s'\n", logPath)

	err := os.MkdirAll(logDir, 0755) // Ensure the logs directory exists
	if err != nil {
		// Use Fatalf to print to stderr and exit if directory creation fails
		log.Fatalf("DEBUG: LogInit - 创建日志目录 '%s' 失败: %v", logDir, err)
	}
	fmt.Printf("DEBUG: LogInit - Directory '%s' ensured/created.\n", logDir)

	// Open the file, creating it if it doesn't exist, and appending to it if it does.
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// Use Fatalf to print to stderr and exit if file open/create fails
		log.Fatalf("DEBUG: LogInit - 打开/创建日志文件 '%s' 失败: %v", logPath, err)
	}
	// Only print creation success after OpenFile succeeds.
	// Standard logger (used by log.Printf below) is already set to os.Stderr by default if not changed.
	// If main.go sets log output to a file, this will go there too.
	log.Printf("%s 日志文件创建成功/已打开", name) 
	fmt.Printf("DEBUG: LogInit - Successfully opened/created log file: '%s'\n", logPath)

	logger := log.New(file, "", log.LstdFlags)
	return logger, file
}
