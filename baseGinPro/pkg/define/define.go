package define

import "time"

//对应数据库表（backpackdb->user_prop_everyday）
type UserPropEveryDay struct {
	GameID      int32     `json:"game_id" gorm:"column:game_id"`
	UserID      int32     `gorm:"column:user_id" json:"user_id"`
	PropNum     int64     `gorm:"column:prop_num" json:"prop_num"`
	PropType    int32     `gorm:"column:prop_type" json:"prop_type"`
	ElementType int32     `gorm:"column:element_type" json:"element_type"`
	LogDate     time.Time `gorm:"column:log_date" json:"log_date"`
	LogTime     int64     `json:"log_time" gorm:"column:log_time"`
}
