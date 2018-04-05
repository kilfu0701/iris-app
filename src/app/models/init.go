package models

import (
	"time"

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

func InsertDummyData(orm *xorm.Engine) error {
	u := &user.User{Username: "user1"}
	has, err := orm.Get(u)
	if err != nil {
		return err
	}

	if !has {
		newUser := &user.User{
			Username: "user1",
			Salt: "hash---",
			Password: "hashed",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if _, err := orm.Insert(newUser); err != nil {
			return err
		}
	}

	op := &operator.Operator{Username: "operator1"}
	has, err = orm.Get(op)
	if err != nil {
		return err
	}

	if !has {
		newOp := &operator.Operator{
			Username: "operator1",
			Salt: "hash---",
			Password: "testonly",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if _, err := orm.Insert(newOp); err != nil {
			return err
		}
	}

	return nil
}
