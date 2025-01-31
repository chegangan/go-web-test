package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	// Cfg 是一个指向 ini.File 结构的指针，用于存储和访问配置文件内容
	Cfg *ini.File

	// RunMode 表示应用程序的运行模式，可以是 "debug" 或 "release"
	RunMode string

	// HTTPPort 表示 HTTP 服务器监听的端口号
	HTTPPort int
	// ReadTimeout 表示读取请求的超时时间
	ReadTimeout time.Duration
	// WriteTimeout 表示写入响应的超时时间
	WriteTimeout time.Duration

	// PageSize 表示分页时每页显示的数据条数
	PageSize int
	// JwtSecret 是用于 JWT（JSON Web Token）签名和验证的密钥
	JwtSecret string
)

// init 函数在包被导入时自动执行，用于初始化配置
func init() {
	var err error
	// 使用 ini.Load 函数加载配置文件 "conf/app.ini"
	Cfg, err = ini.Load("conf/app.ini")
	// 如果加载配置文件出错，则记录错误日志并终止程序
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	// 调用 LoadBase 函数加载基础配置
	LoadBase()
	// 调用 LoadServer 函数加载服务器配置
	LoadServer()
	// 调用 LoadApp 函数加载应用程序配置
	LoadApp()
}

// LoadBase 函数用于加载基础配置，例如运行模式
func LoadBase() {
	// 从配置文件中读取 "RUN_MODE" 键的值，如果不存在则使用默认值 "debug"
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer 函数用于加载服务器配置，例如端口号和超时时间
func LoadServer() {
	// 获取名为 "server" 的节（section）
	sec, err := Cfg.GetSection("server")
	// 如果获取节出错，则记录错误日志并终止程序
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	// 从 "server" 节中读取 "HTTP_PORT" 键的值，如果不存在则使用默认值 8000
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	// 从 "server" 节中读取 "READ_TIMEOUT" 键的值，如果不存在则使用默认值 60，并将其转换为 time.Duration 类型
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	// 从 "server" 节中读取 "WRITE_TIMEOUT" 键的值，如果不存在则使用默认值 60，并将其转换为 time.Duration 类型
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp 函数用于加载应用程序配置，例如 JWT 密钥和分页大小
func LoadApp() {
	// 获取名为 "app" 的节
	sec, err := Cfg.GetSection("app")
	// 如果获取节出错，则记录错误日志并终止程序
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	// 从 "app" 节中读取 "JWT_SECRET" 键的值，如果不存在则使用默认值 "!@)*#)!@U#@*!@!)"
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	// 从 "app" 节中读取 "PAGE_SIZE" 键的值，如果不存在则使用默认值 10
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}
