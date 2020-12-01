package psql

import (
	"baseGinPro/pkg/define"
	"fmt"
)

//postgresql的增删查改测试
func CRUD() {
	//增
	req := define.UserPropEveryDay{
		GameID:   1,
		UserID:   1,
		PropNum:  1,
		PropType: 1,
	}
	db := LPDB.Debug().Table(define.TableUserPropEveryDay)
	db.Save(req)

	//查
	var find define.UserPropEveryDay
	db.Where("id=?", 1).Find(&find)
	fmt.Printf("%+v\n", find)

	//改
	req1 := define.UserPropEveryDay{
		GameID:   2,
		UserID:   2,
		PropNum:  2,
		PropType: 2,
	}
	db.Where("game_id=?", 1).Update(req1)
	db.Where("game_id=?", 2).Find(&find)
	fmt.Printf("%+v\n", find)

	//删
	db.Where("id=?", 1).Delete(define.UserPropEveryDay{})
	db.Where("id=?", 1).Find(&find)
	fmt.Printf("%+v\n", find)
}
