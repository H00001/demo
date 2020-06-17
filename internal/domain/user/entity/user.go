package entity

type User struct {
	// ID
	ID int `gorm:"column:id;primary_key"`

	// username
	UserName string `gorm:"column:user_name"`

	// pwd
	PassWord string `gorm:"column:pass_word"`
}

func (u User) CheckPwd(pwd string) bool {
	if pwd == "" {
		return false
	}

	if u.PassWord == pwd {
		return true
	}

	return false
}
