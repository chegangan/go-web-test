package logging // 定义logging包

import (
	"fmt"           // 导入fmt包，用于格式化字符串
	"log"           // 导入log包，Go语言内置的日志库
	"os"            // 导入os包，用于文件操作
	"path/filepath" // 导入path/filepath包，用于处理文件路径
	"runtime"       // 导入runtime包，用于获取运行时信息
)

// 定义日志级别类型
type Level int

// 定义全局变量
var (
	F *os.File // 日志文件句柄

	DefaultPrefix      = "" // 默认日志前缀，默认为空字符串
	DefaultCallerDepth = 2  // 调用栈深度，用于获取调用者信息，2表示调用Debug/Info/Warn/Error/Fatal函数的上一层调用者

	logger     *log.Logger                                           // 日志记录器
	logPrefix  = ""                                                  // 当前日志前缀
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"} // 日志级别字符串切片
)

// 定义日志级别常量
const (
	DEBUG   Level = iota // DEBUG级别，iota从0开始
	INFO                 // INFO级别，iota递增
	WARNING              // WARNING级别
	ERROR                // ERROR级别
	FATAL                // FATAL级别
)

// 初始化函数，在包加载时自动执行
func init() {
	filePath := getLogFileFullPath() // 获取日志文件完整路径，该函数需要自行实现
	F = openLogFile(filePath)        // 打开/创建日志文件，该函数需要自行实现

	logger = log.New(F, DefaultPrefix, log.LstdFlags) // 创建日志记录器，输出到F文件，使用默认前缀和标准日志格式
}

// Debug日志函数
func Debug(v ...interface{}) { // v ...interface{} 表示接受任意数量的参数
	setPrefix(DEBUG)     // 设置日志前缀
	logger.Println(v...) // 写入日志，Println会自动添加换行符
}

// Info日志函数
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn日志函数
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error日志函数
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal日志函数，会终止程序运行
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...) // 写入日志并终止程序，Fatalln会自动添加换行符
}

// 设置日志前缀函数
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth) // 获取调用者信息，包括文件名和行号
	if ok {                                                 // 如果获取成功
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line) // 构建日志前缀，格式为[级别][文件名:行号]
	} else { // 如果获取失败
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level]) // 构建日志前缀，只包含级别
	}

	logger.SetPrefix(logPrefix) // 设置日志前缀
}
