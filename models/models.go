package models

/**
数据库的配置在这里，包括数据库的连接，关闭，表名的设置等，通过读取配置文件的方式获取数据库的配置信息
*/
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-web-test/pkg/setting"
	"log"
)

import (
	_ "github.com/jinzhu/gorm/dialects/mysql" // 导入 MySQL 驱动，并执行其 init 函数，注册驱动
)

var db *gorm.DB // 定义一个全局变量 db，类型为 *gorm.DB，用于存储数据库连接

type Model struct {
	ID         int `gorm:"primary_key" json:"id"` // ID 字段，设置为主键，JSON 格式化时名称为 "id"
	CreatedOn  int `json:"created_on"`            // CreatedOn 字段，JSON 格式化时名称为 "created_on"
	ModifiedOn int `json:"modified_on"`           // ModifiedOn 字段，JSON 格式化时名称为 "modified_on"
}

// init 函数在包被导入时自动执行，用于初始化数据库连接
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)

	// 从配置文件中读取数据库配置信息
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err) // 如果获取配置信息失败，则记录日志并终止程序
	}

	dbType = sec.Key("TYPE").String()              // 从配置中获取数据库类型
	dbName = sec.Key("NAME").String()              // 从配置中获取数据库名称
	user = sec.Key("USER").String()                // 从配置中获取数据库用户名
	password = sec.Key("PASSWORD").String()        // 从配置中获取数据库密码
	host = sec.Key("HOST").String()                // 从配置中获取数据库主机地址和端口
	tablePrefix = sec.Key("TABLE_PREFIX").String() // 从配置中获取数据表前缀

	// 使用 gorm.Open 函数连接数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err) // 如果连接数据库失败，则记录错误日志
	}

	// 设置表名Handler
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName // 返回添加了表前缀的表名
	}

	db.SingularTable(true)       // 设置表名为单数形式，默认gorm会创建复数形式的表名
	db.LogMode(true)             // 开启GORM日志，会打印sql语句
	db.DB().SetMaxIdleConns(10)  // 设置最大空闲连接数
	db.DB().SetMaxOpenConns(100) // 设置最大打开连接数
}

// CloseDB 函数用于关闭数据库连接
func CloseDB() {
	defer db.Close() // 使用 defer 语句，确保在函数返回前关闭数据库连接
}
