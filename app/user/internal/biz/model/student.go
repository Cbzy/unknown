package model

// Student 学生表
type Student struct {
	Username string `json:"username" gorm:"column:username;comment:用户名"`
	Nickname string `json:"nickname" gorm:"column:nickname;comment:昵称"`
	Password string `json:"-" gorm:"column:password;comment:密码"`
	Gender   string `json:"gender" gorm:"column:gender;comment:性别"`
	Status   string `json:"status" gorm:"column:status;comment:状态"`
}
