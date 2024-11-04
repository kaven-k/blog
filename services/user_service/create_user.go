package user_service

import (
	"errors"
	"server/global"
	"server/models"
	"server/models/ctype"
	"server/utils/pwd"
)

const Avatar = "/uploads/avatar/default.png"

func (UserService) CreateUser(userName, nickName, password string, role ctype.Role, email string, ip string) error {
	// 判断用户名是否存在
	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name = ?", userName).Error
	if err == nil {
		// 存在用户名
		//global.Log.Error("用户名已存在，请重新输入")
		return errors.New("用户名已存在")
	}
	// 对密码进行加密(hash)
	hashPwd := pwd.HashPwd(password)
	// 头像问题
	// 1.默认头像
	// 2.随机选择头像
	//avater := "F:\\Go\\code\\gvb_study\\server\\uploads\\avatar\\default.png"
	//avater := "/uploads/avatar/default.png"

	// 入库
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     Avatar,
		IP:         ip,
		Addr:       "内网地址",          // 内网地址需要根据ip算
		SignStatus: ctype.SignEmail, // 通过邮箱注册
	}).Error
	if err != nil {
		//global.Log.Error(err)
		//return err
		return err
	}
	//global.Log.Infof("用户%s创建成功！", userName)
	return nil
}
