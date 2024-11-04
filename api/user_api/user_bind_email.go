package user_api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"server/global"
	"server/models"
	"server/models/res"
	"server/plugins/email"
	"server/utils/jwts"
	"server/utils/pwd"
	"server/utils/random"
)

type BindEmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password"`
}

func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	// 用户绑定邮箱，第一次输入是 邮箱
	// 后台会给这个邮箱发验证码
	var cr BindEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)
	// 如果用户不传的话，code 传的就是空
	if cr.Code == nil {
		// 为空说明是第一次，后台发验证码
		// 生成4位验证码（这是一个公共的方法 rand 是内置的），生成的四位验证码要和本次的会话保持一致（session）
		code := random.Code(4)
		//fmt.Println(code) // 测试的时候将验证码打印到控制台查看
		// 将生成的验证码存入session
		session.Set("valid_code", code)
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("session错误", c)
			return
		}
		//fmt.Println(err)
		err = email.NewCode().Send(cr.Email, "你的验证码是:"+code)
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}
	// 第二次，用户输入邮箱，验证码，密码 区分两次的关键是有没有code
	code := session.Get("valid_code")
	//fmt.Println(code, *cr.Code)
	// 校验验证码
	if code != *cr.Code {
		res.FailWithMessage("验证码错误", c)
		return
	}
	// 修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMessage("密码强度太低", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.Password)
	// 第一次的邮箱和第二次的邮箱也要做一致性校验，收验证码的是一个邮箱，绑定的是另一个验证码
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashPwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}
	// 完成绑定
	res.OkWithMessage("邮箱绑定成功", c)
	return
}
