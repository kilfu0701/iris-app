package user_profile

import (
	"time"
)

type UserProfile struct {
	ID        int64     // auto-increment by-default by xorm
	UserId    string    `xorm:"varchar(200)"`
	Address   string    `xorm:"address"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
