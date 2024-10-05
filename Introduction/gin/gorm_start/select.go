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

// 	//通过first查询单个数据(升序排序)
// 	// var user User
// 	// db.First(&user)
// 	// fmt.Println(user.ID)

// 	// 通过first查询单个数据(降序排序)
// 	// db.Last(&user)
// 	// fmt.Println(user)

// 	//通过主键查询
// 	// result := db.First(&user, 10)
// 	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 	// 	fmt.Println("未找到")
// 	// }
// 	// fmt.Println(user.ID)

// 	//通过主键查询
// 	// result := db.First(&user, "10")
// 	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 	// 	fmt.Println("未找到")
// 	// }
// 	// fmt.Println(user.ID)

// 	// //通过主键查询
// 	// result := db.First(&user, "abc")
// 	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 	// 	fmt.Println("未找到")
// 	// }
// 	// fmt.Println(user.ID)

// 	//通过主键查询
// 	// result := db.First(&user, []int{1, 3, 5, 13})
// 	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 	// 	fmt.Println("未找到")
// 	// }
// 	// fmt.Println(user.ID)

// 	//检索全部对象
// 	var users []User
// 	result := db.Find(&users)
// 	fmt.Println("总共影响了", result.RowsAffected)
// 	for _, user := range users {
// 		fmt.Println(user.ID)
// 	}
// }
