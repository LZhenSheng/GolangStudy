package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	//日志配置
	//设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // 慢速 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略记录器的 ErrRecordNotFound 错误
			ParameterizedQueries:      true,          // 不要在 SQL 日志中包含参数
			Colorful:                  true,          // 禁用颜色
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err)
	}

	//定义表结构
	//迁移 schema
	_ = db.AutoMigrate(&Product{})

	//create
	db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100})

	//read
	var product Product
	db.First(&product, 1) //根据整形主键查找
	// fmt.Println(product)
	db.First(&product, "code=?", "D42") //查找code=D42的记录

	//UPDATE
	db.Model(&product).Update("price", 200)
	//update多个字短
	db.Model(&product).Updates(Product{Price: 2000, Code: sql.NullString{"", true}}) //仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"price": 2000, "code": "DD42"})
	fmt.Println(product)

	//Delete
	db.Delete(&product, 1)
}
