package main

import (
	"GolangStudy/Practice/shop/user-web/global"
	"GolangStudy/Practice/shop/user-web/initialize"
	"fmt"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	myvalidator "GolangStudy/Practice/shop/user-web/validator"
)

func main() {
	//1.初始化logger
	initialize.InitLogger()
	//2.初始化配置文件
	initialize.InitConfig()
	//3.初始化routers
	Router := initialize.Routers()
	//4.初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		panic(err)
	}
	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	zap.S().Infof("启动服务器,端口:%d", global.ServerConfig.Port)
	Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
}
