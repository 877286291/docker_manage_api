package models

import "time"

type User struct {
	id        int
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	AvatarUrl string    `json:"avatar_url"` //头像地址
	Role      string    `json:"role"`       //用户角色
	Group     string    `json:"group"`      //用户组
	Status    bool      `json:"status"`     //用户状态
	CreatedAt time.Time `json:"created_at"` //创建时间
	UpdatedAt time.Time `json:"updated_at"` //更新时间
}
