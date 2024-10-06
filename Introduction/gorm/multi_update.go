package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// type User struct {
// 	ID           uint           // Standard field for the primary key
// 	Name         string         // A regular string field
// 	Email        *string        // A pointer to a string, allowing for null values
// 	Age          uint8          // An unsigned 8-bit integer
// 	Birthday     *time.Time     // A pointer to time.Time, can be null
// 	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
// 	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
// 	CreatedAt    time.Time      // Automatically managed by GORM for creation time
// 	UpdatedAt    time.Time      // Automatically managed by GORM for update time
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

// 	//定义表结构
// 	//迁移 schema
// 	_ = db.AutoMigrate(&User{})

// 	var users = []User{
// 		{Name: "jinzhu1"},
// 		{Name: "jinzhu2"},
// 		{Name: "jinzhu3"},
// 	}
// 	//单一的sql语句
// 	// db.Create(&users)

// 	//为什么不一次性提交所有的？
// 	//sql语句有长度限制，需要分批提交
// 	db.CreateInBatches(users, 2)

// 	//for简单插入
// 	for _, user := range users {
// 		fmt.Println(user.ID)
// 	}

// 	// batch insert from `[]map[string]interface{}{}`
// 	db.Model(&User{}).Create([]map[string]interface{}{
// 		{"Name": "jinzhu_1", "Age": 18},
// 		{"Name": "jinzhu_2", "Age": 20},
// 	})
// }
