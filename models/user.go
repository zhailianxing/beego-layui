package models

import "github.com/astaxie/beego/orm"
import _ "github.com/go-sql-driver/mysql"

type User struct {
	Id       int64
	UserName string
	Pwd      string
	Sex      int
	Address  string
	Phone    string
	UserType int // 超级管理员，系统用户等
	Status   int
}

func Get(username, pwd string) (error, *User) {

	o := orm.NewOrm()

	obj := User{}

	err := o.QueryTable("user").Filter("user_name", username).Filter("pwd", pwd).One(&obj)
	if err != nil {
		return err, nil
	}

	return nil, &obj
}

//
//// 修改
//func Update(city string, num int) error {
//
//	o := orm.NewOrm()
//
//	_, err := o.QueryTable("fei_yan").Filter("city", city).Update(orm.Params{
//		"num": num,
//	})
//	return err
//}
//
//// 插入
//func Add(city string, num int) error {
//
//	o := orm.NewOrm()
//
//	data := &FeiYan{City: city, Num: int64(num)}
//
//	_, err := o.Insert(data)
//
//	return err
//
//}
//
//// 删除
//func Del(id int) error {
//
//	o := orm.NewOrm()
//
//	//data := &FeiYan{Id: int64(id)}
//	//_, err := o.Delete(data)
//
//	data := FeiYan{Id: int64(id)}
//	_, err := o.Delete(&data)
//
//	return err
//
//}
