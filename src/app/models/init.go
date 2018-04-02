package models

import (
	"github.com/go-xorm/xorm"

	"app/models/user"
	"app/models/operator"
)

func InitTables(orm *xorm.Engine) error {
	if err := orm.Sync2(new(user.User)); err != nil {
		return err
	}

	if err := orm.Sync2(new(operator.Operator)); err != nil {
		return err
	}

	return nil
}
