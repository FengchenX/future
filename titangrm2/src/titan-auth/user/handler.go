package user

import (
	"fmt"
	"strconv"
	"time"

	"github.com/emicklei/go-restful"

	"grm-service/log"
	. "grm-service/util"
	"titan-auth/captcha"
	"titan-auth/types"
)

var (
	sessionTTL = 5 * time.Hour
)

// POST http://localhost:8080/captcha
func (s UserSvc) getCaptcha(req *restful.Request, res *restful.Response) {
	info := captchaRequest{}
	err := req.ReadEntity(&info)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	id, capt, err := captcha.GenerateCaptcha(s.DynamicDB.Cli, info.NumCount, info.Height, info.Width)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	if len(id) == 0 || len(capt) == 0 {
		log.Error("Failed to get captcha.")
		ResWriteError(res, nil)
	}
	ResWriteEntity(res, &captchaPic{id, capt})
}

// PUT http://localhost:8080/login
func (s UserSvc) login(req *restful.Request, res *restful.Response) {
	loginInfo := userlogin{}
	err := req.ReadEntity(&loginInfo)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 验证码
	if len(loginInfo.Captcha.Id) == 0 || len(loginInfo.Captcha.Captcha) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid captcha")))
		return
	}
	if ret := captcha.VerifyCaptcha(s.DynamicDB.Cli, loginInfo.Captcha.Id, loginInfo.Captcha.Captcha); !ret {
		ResWriteError(res, fmt.Errorf(TR("Invalid captcha:%s", loginInfo.Captcha.Captcha)))
		return
	}

	// 密码
	user, err := s.DynamicDB.UserLogin(loginInfo.User, loginInfo.Password)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	go s.DynamicDB.UpdateUserTime(user.Id)

	// 创建session
	session, err := s.DynamicDB.CreateUserSession(user.Id, int64(sessionTTL.Seconds()))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	user.Session = session
	ResWriteEntity(res, &user)
}

// PUT http://localhost:8080/logout
func (s UserSvc) logout(req *restful.Request, res *restful.Response) {
	s.DynamicDB.DelUserSession(req.HeaderParameter("auth-session"))
	ResWriteEntity(res, nil)
}

// POST http://localhost:8080/users
func (s UserSvc) userRegistry(req *restful.Request, res *restful.Response) {
	args := userRegistry{}
	err := req.ReadEntity(&args)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	if len(args.User) == 0 || len(args.Email) == 0 || len(args.Password) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid user info")))
		return
	}

	// 判断用户名是否重复
	users, err := s.DynamicDB.GetUserList()
	if err != nil {
		ResWriteError(res, err)
		return
	}
	for _, v := range users {
		if v.User == args.User {
			ResWriteError(res, fmt.Errorf(TR("User is already exists")))
			return
		}
	}

	user := types.User{
		User:     args.User,
		Name:     args.Name,
		Profile:  args.Profile,
		Email:    args.Email,
		Password: args.Password,
	}
	ret, err := s.DynamicDB.AddUser(&user)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, ret)
}

// PUT http://localhost:8080/users/{user-id}/status
func (s UserSvc) userActive(req *restful.Request, res *restful.Response) {
	err := s.DynamicDB.UpdateUserStatus(req.PathParameter("user-id"))
	if err != nil {
		ResWriteError(res, err)
	}
	ResWriteEntity(res, nil)
}

// GET http://localhost:8080/users
func (s UserSvc) userList(req *restful.Request, res *restful.Response) {
	ret, err := s.DynamicDB.GetUserList()
	if err != nil {
		ResWriteError(res, err)
	}

	var users types.UserList
	for _, val := range ret {
		fmt.Println(val)
		if val.Type != types.User_Type_Admin {
			users.Users = append(users.Users, val)
			users.Total++
		}
	}

	page := ParserPageArgs(req)
	// 分页
	offset, _ := strconv.Atoi(page.Offset)
	limit, _ := strconv.Atoi(page.Limit)
	if offset >= 0 && limit > 0 {
		if offset+limit >= len(users.Users)-1 {
			users.Users = users.Users[offset:]
		} else {
			users.Users = users.Users[offset : offset+limit]
		}
	}

	// 排序
	userSort(users.Users, page.Sort, page.Order)
	if len(users.Users) > 0 {
		ResWriteEntity(res, &users)
	} else {
		ResWriteEntity(res, nil)
	}
}
