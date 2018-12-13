package user

import (
	"time"

	//"github.com/go-xorm/xorm"
	"app/models/user_profile"
)

type User struct {
	ID        int64                     // auto-increment by-default by xorm
	Version   string                    `xorm:"varchar(200)"`
	Salt      string
	Username  string
	Password  string                    `xorm:"varchar(200)"`
	Languages string                    `xorm:"varchar(200)"`
	Profile   *user_profile.UserProfile `xorm:"'user_profile_id'"`
	CreatedAt time.Time                 `xorm:"created"`
	UpdatedAt time.Time                 `xorm:"updated"`
}
