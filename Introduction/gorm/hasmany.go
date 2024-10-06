package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// // User has many CreditCards, UserID is the foreign key
// type User struct {
// 	gorm.Model
// 	CreditCards []CreditCard
// }

// type CreditCard struct {
// 	gorm.Model
// 	Number string
// 	UserID uint
// }

// func main() {
// 	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
// 	dsn := "root:root123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

// 	//日志配置
// 	//设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
// 	newLogger := logger.New(
// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold:             time.Second, // 慢速 SQL 阈值
// 			LogLevel:                  logger.Info, // 日志级别
// 			IgnoreRecordNotFoundError: true,        // 忽略记录器的 ErrRecordNotFound 错误
// 			ParameterizedQueries:      true,        // 不要在 SQL 日志中包含参数
// 			Colorful:                  true,        // 禁用颜色
// 		},
// 	)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		Logger: newLogger,
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// 定义表结构
// 	// 迁移 schema
// 	// _ = db.AutoMigrate(&CreditCard{})
// 	// _ = db.AutoMigrate(&User{})

// 	//在大型系统和高并发系统一般不使用外键约束，自己在业务层面保证数据的一致性，外键约束影响性能
// 	//外间约束可以保证数据的完整性
// 	user := User{}
// 	// db.Create(&user)
// 	// db.Create(&CreditCard{
// 	// 	Number: "12",
// 	// 	UserID: user.ID,
// 	// })
// 	// db.Create(&CreditCard{
// 	// 	Number: "34",
// 	// 	UserID: user.ID,
// 	// })

// 	db.Preload("CreditCards").First(&user)
// 	for _, card := range user.CreditCards {
// 		fmt.Println(card)
// 	}
// }
