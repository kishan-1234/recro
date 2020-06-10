package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id          int    `orm:"column(id);auto"`
	Name        string `orm:"column(name);size(30);null"`
	Email       string `orm:"column(email);size(30);null"`
	PhoneNumber string `orm:"column(phoneNumber);size(10);null"`
	Meta        string `orm:"column(meta);size(200);null"`
	Password    string `orm:"column(password);size(10);null"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func GetAllUsers() (v []User, err error) {
	o := orm.NewOrm()
	v = []User{}
	_, err = o.QueryTable(new(User)).All(&v)
	if err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserById(Id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{}
	if err = o.QueryTable("user").Filter("Id", Id).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserByPhoneNumber(phoneNumber string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{}
	if err = o.QueryTable("user").Filter("phoneNumber", phoneNumber).One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetUserByEmailandPassword(password string, email string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{}
	if err = o.QueryTable("user").Filter("email", email).Filter("password", password).One(v); err == nil {
		return v, nil
	}
	return nil, err
}
