//author xinbing
//time 2018/9/11 14:57
//
package service

import (
	"common-utilities/encrypt"
	"common-utilities/utilities"
	"github.com/sirupsen/logrus"
	"ibs_service/saas_background/dbs"
	"ibs_service/saas_background/entity"
	"ibs_service/saas_background/store"
	"ibs_service/saas_background/utils"
)

func Login(userName, pwd string) (*utils.Resp, string) {
	var dbUser entity.SaasUsers
	dbs.SaasGormDB.Client.Where("username = ?", userName).First(&dbUser)
	if dbUser.ID == 0 || encryptUserPwd(pwd) != dbUser.Pwd {
		return utils.Resp{}.Failed("用户名或密码错误！"), ""
	}
	accessToken := encrypt.SHA1(utilities.GetRandomStr(32) + pwd)
	store.CacheUser(accessToken, dbUser.ID)
	return utils.Resp{}.Success("登录成功", dbUser), accessToken
}

func GetUserInfo(userId uint) *utils.Resp {
	var dbUser entity.SaasUsers
	dbs.SaasGormDB.Client.Where("id = ?", userId).First(&dbUser)
	if dbUser.ID == 0 {
		return utils.Resp{}.Failed("要查找的用户不存在！")
	}
	return utils.Resp{}.Success("获取成功！", dbUser)
}

func AddUser(userDTO entity.SaasUsers) *utils.Resp {
	if userDTO.Username == "" || userDTO.Pwd == ""{
		return utils.Resp{}.Failed("用户名或密码不允许为空！")
	}
	count := 0
	dbs.SaasGormDB.Client.Model(&userDTO).Where("username = ?",userDTO.Username).Count(&count)
	if count > 0 {
		return utils.Resp{}.Failed("用户名已存在！")
	}
	err := dbs.SaasGormDB.Client.Create(&entity.SaasUsers{
		Username: userDTO.Username,
		Pwd: encryptUserPwd(userDTO.Pwd),
	}).Error
	if err != nil {
		logrus.WithError(err).WithField("userDTO",userDTO).Errorln("AddUser failed")
		return utils.Resp{}.Failed("添加用户失败！")
	}
	return utils.Resp{}.Success("添加用户成功！", nil)
}

func encryptUserPwd(pwd string) string{
	return encrypt.SHA1("君生我未生我生君已老" + pwd)
}