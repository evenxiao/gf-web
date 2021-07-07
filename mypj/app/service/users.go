package service

import (
	"mypj/app/dao"
	"mypj/app/model"
)

func GetUserInfo() (list []*model.Users, err error) {

	//var a interface{}
	//var a  map[string]interface{}
	//
	//a = make(map[string]interface{})
    // a["id"] = 111;
    // a["name"] = "awennnnnn";
	list, err = dao.Users.All("user_id = 10")
	return list, err
}
