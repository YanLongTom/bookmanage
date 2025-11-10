package sql

import (
	"booksmanage/model"
	"errors"
)

var EmptyErr = errors.New("Null data")

func findUser(u *model.User) error {
	var ret []model.User
	result := DB.Where("name = ?", u.Name).Find(&ret)
	if result.Error != nil {
		return result.Error
	}
	if len(ret) > 0 {
		return EmptyErr
	}
	return nil
}

func SaveUser(u *model.User) (error, string) {
	err := findUser(u)
	if err == EmptyErr {
		return nil, "login success"
	}
	if err == nil {
		ret := DB.Create(u)
		if ret.Error != nil {
			return err, "create failed"
		}
		return nil, "create success"
	}
	return err, "error"
}
