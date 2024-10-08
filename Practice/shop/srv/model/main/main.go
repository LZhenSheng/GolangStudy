package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/anaskhan96/go-password-encoder"
)

func genMD5(code string) string {
	MD5 := md5.New()
	_, _ = io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}
func main() {
	// // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "root:root123456@tcp(127.0.0.1:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"

	// //日志配置
	// //设置全局的logger，这个logger在我们执行每个sql语句的时候会打印每一行sql
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold:             time.Second, // 慢速 SQL 阈值
	// 		LogLevel:                  logger.Info, // 日志级别
	// 		IgnoreRecordNotFoundError: true,        // 忽略记录器的 ErrRecordNotFound 错误
	// 		ParameterizedQueries:      true,        // 不要在 SQL 日志中包含参数
	// 		Colorful:                  true,        // 禁用颜色
	// 	},
	// )

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// 	Logger: newLogger,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// //定义表结构
	// //迁移 schema
	// _ = db.AutoMigrate(&model.User{})
	// fmt.Println(genMD5("sdlfaksjlfjsdj23432984729eifksadjfjsa;fkl"))
	// fmt.Println(genMD5("12345"))

	// Using the default options
	// salt, encodedPwd := password.Encode("generic password", nil)
	// fmt.Println(salt)
	// fmt.Println(encodedPwd)
	// check := password.Verify("generic password", salt, encodedPwd, nil)
	// fmt.Println(check) // true

	// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	//$算法$
	pwd := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(len(pwd))
	fmt.Println(salt)
	fmt.Println(encodedPwd)
	pwdInfo := strings.Split(pwd, "$")
	check := password.Verify("generic password", pwdInfo[2], pwdInfo[3], options)
	fmt.Println(check) // true
}
