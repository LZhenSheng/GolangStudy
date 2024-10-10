package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"GolangStudy/Practice/shop/srv/proto"
	"GolangStudy/Practice/shop/user-web/forms"
	"GolangStudy/Practice/shop/user-web/global"
	"GolangStudy/Practice/shop/user-web/global/response"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

// func InitTrans(locale string) (err error) {
// 	//修改gin的validator引擎属性，实现定制
// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		//注册一个获取json的tag的自定义方法
// 		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
// 			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
// 			if name == "-" {
// 				return ""
// 			}
// 			return name
// 		})
// 		zhT := zh.New() //中文翻译器
// 		enT := en.New() //英文翻译器
// 		//第一个是备用的，后面的是支持的语言环境
// 		uni := ut.New(enT, zhT, enT)
// 		trans, ok = uni.GetTranslator(locale)
// 		if !ok {
// 			return fmt.Errorf("uni.GetTranslator(%s)", locale)
// 		}
// 		switch locale {
// 		case "en":
// 			en_translations.RegisterDefaultTranslations(v, trans)
// 		case "zh":
// 			zh_translations.RegisterDefaultTranslations(v, trans)
// 		default:
// 			en_translations.RegisterDefaultTranslations(v, trans)
// 		}
// 		return
// 	}
// 	return
// }

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc的code转换成http的状态吗
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}
func GetUserList(ctx *gin.Context) {
	//拨号连接用户grpc服务器
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接[用户服务失败]",
			"msg", err.Error(),
		)
	}
	//生成grpc的client并调用接口
	userSrvClient := proto.NewUserClient(userConn)
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询[用户列表]失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	zap.S().Debug("获取用户列表页")
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			// BirthDay: time.Time(time.Unix(int64(value.BirthDay), 0)),
			BirthDay: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}
func PassWordLogin(ctx *gin.Context) {
	//表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": removeTopStruct(errs.Translate(global.Trans)),
			})
		}
		fmt.Println(err.Error())
		return
	}
}
