package forms

type PassWordLoginForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required"` //自定义validator实现手机号码格式验证
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
}
