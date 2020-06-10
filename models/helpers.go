package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func MysqlTest() error {
	o := orm.NewOrm()
	_, err := o.Raw("SELECT 1").Exec()
	return err
}

func UpdateRowByColumns(m interface{}, updatedBy string, columns ...string) (err error) {
	o := orm.NewOrm()
	err = UpdateRowByColumnsByORM(m, o, updatedBy, columns...)
	return err
}

func UpdateRowByColumnsByORM(m interface{}, o orm.Ormer, updatedBy string, columns ...string) (err error) {
	if _, err = o.Update(m, columns...); err != nil {
		fmt.Println(err)
	}
	// call for insertion ypusertrace table
	return
}
