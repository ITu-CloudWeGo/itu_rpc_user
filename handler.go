package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	dao "github.com/ITu-CloudWeGo/itu_rpc_user/db"
	module "github.com/ITu-CloudWeGo/itu_rpc_user/db/model"
	user_service "github.com/ITu-CloudWeGo/itu_rpc_user/kitex_gen/user_service"
	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Sign implements the UserServiceImpl interface.
func (s *UserServiceImpl) Sign(ctx context.Context, req *user_service.SignRequest) (resp *user_service.SignResponse, err error) {
	// TODO: Your code here...
	userDao := dao.GetUserDao()
	if existtingEmail, err := userDao.FindByEmail(req.Email); err != nil {
		return &user_service.SignResponse{
			Response: &user_service.Response{
				Status: 500,
				Msg:    "注册失败，请稍后重试",
			},
		}, err
	} else if existtingEmail != nil {
		return &user_service.SignResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "邮箱已被注册",
			},
		}, nil
	}
	// 生成随机盐值
	salt, err := generateSalt(16)
	if err != nil {
		klog.Error("生成盐值失败:", err)
		return &user_service.SignResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "注册失败，请稍后重试",
			},
		}, nil
	}
	//加密密码
	hashedPassword, err := hashPasswordWithSalt(req.Password, salt)
	if err != nil {
		klog.Error("密码加密失败:", err)
		return &user_service.SignResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "注册失败，请稍后重试",
			},
		}, nil
	}
	data := &module.User{
		Email:    req.Email,
		UserName: req.Username,
		Passwd:   hashedPassword,
		NickName: req.Nickname,
		Icon:     req.Icon,
		Salt:     salt,
	}
	if err := userDao.Insert(data); err != nil {
		klog.Error("用户注册失败:", err)
		return &user_service.SignResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "注册失败，请稍后重试",
			},
		}, nil
	}
	klog.Info("用户注册成功:", data)
	return &user_service.SignResponse{
		Response: &user_service.Response{
			Status: 200,
			Msg:    "注册成功",
		},
	}, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user_service.LoginRequest) (resp *user_service.LoginResponse, err error) {
	// TODO: Your code here...
	userDao := dao.GetUserDao()
	if userDao == nil {
		klog.Error("userDao 未初始化")
		return &user_service.LoginResponse{
			Response: &user_service.Response{
				Status: 500,
				Msg:    "系统错误请稍后重试",
			},
		}, nil
	}

	var user *module.User

	if req.Email != "" {
		user, err = userDao.FindByEmail(req.Email)
		if err != nil {
			klog.Error("登录失败，用户不存在：", err)
			return &user_service.LoginResponse{
				Response: &user_service.Response{
					Status: 403,
					Msg:    "用户不存在",
				},
			}, nil
		}
		if user == nil {
			klog.Info("用户不存在")
			return &user_service.LoginResponse{
				Response: &user_service.Response{
					Status: 403,
					Msg:    "用户不存在",
				},
			}, nil
		}
	} else if req.Username != "" {
		user, err = userDao.FindByUserName(req.Email)
		if err != nil {
			klog.Error("登录失败，用户不存在：", err)
			return &user_service.LoginResponse{
				Response: &user_service.Response{
					Status: 403,
					Msg:    "用户不存在",
				},
			}, nil
		}
		if user == nil {
			klog.Info("用户不存在")
			return &user_service.LoginResponse{
				Response: &user_service.Response{
					Status: 403,
					Msg:    "用户不存在",
				},
			}, nil
		}
	} else {
		return &user_service.LoginResponse{
			Response: &user_service.Response{
				Status: 400,
				Msg:    "请输入邮箱或用户名",
			},
		}, nil
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(req.Password+user.Salt)); err != nil {
		klog.Error("密码验证失败:", err)
		return &user_service.LoginResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "邮箱或密码不正确",
			},
		}, nil
	}
	klog.Info("用户登录成功:", user)
	return &user_service.LoginResponse{
		Response: &user_service.Response{
			Status: 200,
			Msg:    "登录成功",
		},
	}, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user_service.GetUserRequest) (resp *user_service.GetUserResponse, err error) {
	// TODO: Your code here...
	userDao := dao.GetUserDao()
	user, err := userDao.FindOne(req.Uid)
	if err != nil {
		klog.Error("查询错误：", err)
		return &user_service.GetUserResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "用户不存在",
			},
		}, nil
	}
	if user == nil {
		return &user_service.GetUserResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "用户不存在",
			},
		}, nil
	}
	return &user_service.GetUserResponse{
		Response: &user_service.Response{
			Status: 200,
			Msg:    "获取用户成功",
		},
		Username: user.UserName,
		Nickname: user.NickName,
		Icon:     user.Icon,
	}, nil
}

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *user_service.UpdateUserRequest) (resp *user_service.UpdateUserResponse, err error) {
	// TODO: Your code here...
	userDao := dao.GetUserDao()
	user, err := userDao.FindOne(req.Uid)
	if err != nil {
		klog.Info("用户不存在：", err)
		return &user_service.UpdateUserResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "查询用户失败，请稍后重试",
			},
		}, nil
	}
	if user == nil {
		klog.Info("用户不存在")
		return &user_service.UpdateUserResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "用户不存在",
			},
		}, nil
	}
	// 验证当前密码
	saltedCurrentPassword := req.Password + user.Salt
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(saltedCurrentPassword)); err != nil {
		klog.Error("当前密码错误：", err)
		return &user_service.UpdateUserResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "密码错误",
			},
		}, nil
	}

	//更新邮箱
	if req.NewEmail != "" {
		user.Email = req.NewEmail
	}

	//更新密码
	if req.NewPassword != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			klog.Error("密码加密失败:", err)
			return &user_service.UpdateUserResponse{
				Response: &user_service.Response{
					Status: 403,
					Msg:    "更新失败，请稍后重试",
				},
			}, nil
		}
		user.Passwd = string(hashedPassword)
	}

	//更新用户名
	if req.NewNickname != "" {
		user.NickName = req.NewNickname
	}

	if req.NewIcon != "" {
		user.Icon = req.NewIcon
	}

	//更新数据库和Redis缓存
	user.Updated_at = time.Now()
	if err := userDao.Update(user); err != nil {
		klog.Error("更新用户失败：", err)
		return &user_service.UpdateUserResponse{
			Response: &user_service.Response{
				Status: 403,
				Msg:    "更新用户失败",
			},
		}, nil
	}

	// 同步更新 Redis 缓存
	cacheKey := user.GetCacheKey()
	if err := userDao.Cache().Set(ctx, cacheKey, user, time.Hour).Err(); err != nil {
		klog.Error("更新 Redis 缓存失败：", err)
	}
	klog.Info("更新成功")
	return &user_service.UpdateUserResponse{
		Response: &user_service.Response{
			Status: 200,
			Msg:    "更新成功",
		},
	}, nil
}

func generateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

// 密码加盐并哈希
func hashPasswordWithSalt(password string, salt string) (string, error) {
	// 将盐值与密码结合
	saltedPassword := password + salt
	// 使用 bcrypt 进行哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(saltedPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
