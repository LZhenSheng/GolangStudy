package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Language struct {
	// gorm.Model
	Name    string
	AddTime time.Time //创建时自动加入时间
}

// // 在gorm中可以通过给某一个结构体添加TableName实现自定义表名
//
//	func (Language) TableName() string {
//		return "my_language"
//	}

func (l *Language) BeforeCreate(tx *gorm.DB) (err error) {
	l.AddTime = time.Now()
	return nil
}
func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	//日志配置
	//设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢速 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略记录器的 ErrRecordNotFound 错误
			ParameterizedQueries:      true,        // 不要在 SQL 日志中包含参数
			Colorful:                  true,        // 禁用颜色
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//添加前缀
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "maxshop_",
		},
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
	}

	// 定义表结构
	// 迁移 schema
	_ = db.AutoMigrate(&Language{})
	db.Create(&Language{
		Name: "ptyhon",
	})

}
