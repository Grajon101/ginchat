package models

import (
	"fmt"
	"ginchat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartBeatTime time.Time
	LoginOutTime  time.Time
	IsLogiut      bool
	DeviceInfo    string
}

func (table *UserBasic) TabelName() string {
	return "user_basic"
}

// 获取用户集合
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

// 创建用户
func CreateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Create(user)
}

// 删除用户
func DeleteUser(user *UserBasic) *gorm.DB {
	return utils.DB.Delete(user)
}

// 更新用户
func UpdateUser(user *UserBasic) *gorm.DB {
	return utils.DB.Model(user).Updates(
		UserBasic{
			Name:     user.Name,
			PassWord: user.PassWord,
			Phone:    user.Phone,
			Email:    user.Email,
		},
	)
}

// 根据用户名获取用户
func GetUserByName(name string) *UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return &user
}

// 根据用户名和密码获取用户
func GetUserByNameAndPwd(name, password string) *UserBasic {
	user := UserBasic{}
	//
	utils.DB.Where(&UserBasic{Name: name, PassWord: password}).First(&user)
	//更新token加密
	str := utils.MD5Encode(fmt.Sprintf("%d", time.Now().Unix()))
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", str)
	return &user
}

// 根据电话获取用户
func GetUserByPhone(phone string) *UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone = ?", phone).First(&user)
	return &user
}

// 根据邮箱获取用户
func GetUserByEmail(email string) *UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ?", email).First(&user)
	return &user
}
