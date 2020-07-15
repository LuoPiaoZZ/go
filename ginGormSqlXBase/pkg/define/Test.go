package define

import "github.com/jmoiron/sqlx/types"

type Test struct {
	Test types.JSONText `gorm:"test" json:"test" form:"test"`
}
