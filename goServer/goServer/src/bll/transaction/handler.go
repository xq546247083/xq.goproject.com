package transaction

import (
	"xq.goproject.com/Vendor/github.com/jinzhu/gorm"
	"xq.goproject.com/goServer/goServer/src/dal"
)

//Handle 事务处理
func Handle(dbAction func(*gorm.DB) error) {
	DBTemp := dal.DB.Begin()
	defer func() {
		if err := recover(); err != nil {
			DBTemp.Callback()
		}
	}()

	//调用事务方法
	if dbAction != nil {
		if err := dbAction(DBTemp); err != nil {
			DBTemp.Callback()
			return
		}
	}

	DBTemp.Commit()
}
