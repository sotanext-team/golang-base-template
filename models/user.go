package models

type User struct {
	Model

	Username string `json:"username" gorm:"size:100;index:user_name_idx,unique"`
	Password string `json:"password" gorm:"size:100"`
	Email    string `json:"email" gorm:"size:100;index:user_email_idx,unique"`
	Token    string `json:"token" gorm:"size:100"`

	Shops     []*Shop     `json:"shop" gorm:"many2many:user_shops"`
	UserMetas []*UserMeta `json:"userMetas"`

	Roles []string `json:"roles" gorm:"-"`
}

type UserCheck struct {
	Model `mapstructure:",squash"`

	Username string `json:"username" `
	Email    string `json:"email" `

	Shops []*Shop  `json:"shop" gorm:"many2many:user_shops"`
	Roles []string `json:"roles" gorm:"-"`
}

func (UserCheck) TableName() string {
	return "users"
}

type UserToken struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
