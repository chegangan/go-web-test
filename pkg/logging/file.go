package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
	// 尝试获取文件信息
	_, err := os.Stat(filePath)
	// 使用 switch 语句根据错误类型进行处理
	switch {
	// 如果文件不存在
	case os.IsNotExist(err):
		// 创建日志文件目录
		mkDir()
	// 如果权限不足
	case os.IsPermission(err):
		// 记录错误并终止程序
		log.Fatalf("Permission :%v", err)
	}

	// 以追加、创建、只写模式打开文件，权限为 0644
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// 如果打开文件失败
	if err != nil {
		// 记录错误并终止程序
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	// 返回文件句柄
	return handle
}

func mkDir() {
	// 获取当前工作目录
	dir, _ := os.Getwd()
	// 创建日志文件目录，权限为所有用户可读写执行
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	// 如果创建目录失败
	if err != nil {
		// 抛出panic异常，终止程序
		panic(err)
	}
}
