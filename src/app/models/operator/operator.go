package operator

import (
	"time"

	//"github.com/go-xorm/xorm"
)

type Operator struct {
	ID        int64     // auto-increment by-default by xorm
	Version   string    `xorm:"varchar(200)"`
	Salt      string
	Username  string
	Password  string    `xorm:"varchar(200)"`
	Languages string    `xorm:"varchar(200)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
}
