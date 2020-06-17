package dto

import "test/internal/domain/user/entity"

type User struct {
	// 用户名称
	UserName string `json:"user_name"`

	// 密码
	PassWord string `json:"pass_word"`
}

func NewUser(user entity.User) User {
	dto := User{}

	dto.UserName = user.UserName
	dto.PassWord = user.PassWord
	return dto
}
