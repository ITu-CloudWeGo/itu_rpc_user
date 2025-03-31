package dao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ITu-CloudWeGo/itu_rpc_user/config"
	"github.com/ITu-CloudWeGo/itu_rpc_user/db/model"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

type UserDaoImpl struct {
	db    *gorm.DB
	cache *redis.Client
}

type UserDao interface {
	Insert(data *model.User) error
	FindOne(Uid int64) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByUserName(userName string) (*model.User, error)
	Update(user *model.User) error
	Cache() *redis.Client
}

var (
	instanceUserDao *UserDaoImpl
	onceUserDao     sync.Once
)

func postDsn() string {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		conf.PostgresSQL.Host,
		conf.PostgresSQL.User,
		conf.PostgresSQL.Password,
		conf.PostgresSQL.DBName,
		conf.PostgresSQL.Port,
		conf.PostgresSQL.Sslmode,
	)
	return dsn
}

func initUserDao(dsn string, redisOpt *redis.Options) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}
	cache := redis.NewClient(redisOpt)
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	instanceUserDao = &UserDaoImpl{
		db:    db,
		cache: cache,
	}

}

func GetUserDao() UserDao {
	onceUserDao.Do(func() {
		conf := config.GetConfig()
		dsn := postDsn()
		initUserDao(dsn, &redis.Options{
			Addr:     conf.Redis.Address,
			Username: conf.Redis.Username,
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
		})
	})
	return instanceUserDao
}

func (dao *UserDaoImpl) Insert(data *model.User) error {
	return dao.db.Create(data).Error
}

func (dao *UserDaoImpl) FindOne(uid int64) (*model.User, error) {
	q := &model.User{
		Uid: uid,
	}
	var redisResult model.User
	err := dao.cache.Get(context.Background(), q.GetCacheKey()).Scan(redisResult)
	if err == nil {
		return &redisResult, nil
	}
	err = dao.db.First(q).Error
	if err != nil {
		return nil, err
	}
	return q, nil
}

func (dao *UserDaoImpl) FindByEmail(email string) (*model.User, error) {
	// 尝试从 Redis 缓存中获取用户信息
	cacheKey := fmt.Sprintf("users:email:%s", email)
	val, err := dao.cache.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var cachedUser model.User
		err := json.Unmarshal([]byte(val), &cachedUser)
		if err != nil {
			klog.Error("用户信息反序列化失败，Email:", email, " 错误信息:", err)
			return nil, fmt.Errorf("用户信息反序列化失败: %v", err)
		}
		return &cachedUser, nil
	}
	if err != redis.Nil {
		klog.Error("Redis 查询失败:", err)
		return nil, err
	}

	// 如果缓存中没有，从数据库中查询
	var user model.User
	err = dao.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在，返回 nil, nil
		}
		klog.Error("数据库查询失败:", err)
		return nil, err // 其他错误，返回具体错误信息
	}

	// 将查询到的用户信息存入 Redis 缓存
	userJSON, err := json.Marshal(user)
	if err != nil {
		klog.Error("用户信息序列化失败，Email:", email, " 错误信息:", err)
		return nil, fmt.Errorf("用户信息序列化失败: %v", err)
	}

	err = dao.cache.Set(context.Background(), cacheKey, userJSON, time.Hour).Err()
	if err != nil {
		klog.Error("设置 Redis 缓存失败:", err)
	}

	return &user, nil
}

func (dao *UserDaoImpl) FindByUserName(userName string) (*model.User, error) {
	// 尝试从 Redis 缓存中获取用户信息
	cacheKey := fmt.Sprintf("users:user_name:%s", userName)
	val, err := dao.cache.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var cachedUser model.User
		err := json.Unmarshal([]byte(val), &cachedUser)
		if err != nil {
			klog.Error("用户信息反序列化失败，UserName:", userName, " 错误信息:", err)
			return nil, fmt.Errorf("用户信息反序列化失败: %v", err)
		}
		return &cachedUser, nil
	}
	if err != redis.Nil {
		klog.Error("Redis 查询失败:", err)
		return nil, err
	}

	// 如果缓存中没有，从数据库中查询
	var user model.User
	err = dao.db.Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // 用户不存在，返回 nil, nil
		}
		klog.Error("数据库查询失败:", err)
		return nil, err // 其他错误，返回具体错误信息
	}

	// 将查询到的用户信息存入 Redis 缓存
	userJSON, err := json.Marshal(user)
	if err != nil {
		klog.Error("用户信息序列化失败，UserName:", userName, " 错误信息:", err)
		return nil, fmt.Errorf("用户信息序列化失败: %v", err)
	}

	err = dao.cache.Set(context.Background(), cacheKey, userJSON, time.Hour).Err()
	if err != nil {
		klog.Error("设置 Redis 缓存失败:", err)
	}

	return &user, nil
}

func (dao *UserDaoImpl) Update(user *model.User) error {
	err := dao.db.Save(user).Error
	if err != nil {
		return err
	}
	cacheKey := user.GetCacheKey()

	userJSON, err := json.Marshal(user)
	if err != nil {
		klog.Error("用户信息序列化失败，UserId:", user.Uid, " 错误信息:", err)
		return fmt.Errorf("用户信息序列化失败: %v", err)
	}

	err = dao.cache.Set(context.Background(), cacheKey, userJSON, time.Hour).Err()
	if err != nil {
		klog.Error("设置 Redis 缓存失败，UserId:", user.Uid, " 错误信息:", err)
		return fmt.Errorf("设置 Redis 缓存失败: %v", err)
	}

	return nil
}

func (dao *UserDaoImpl) Cache() *redis.Client { return dao.cache }
