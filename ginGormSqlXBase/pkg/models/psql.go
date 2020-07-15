package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

var (
	MyDB1, MyDB *gorm.DB
	MyDBSqlx    *sqlx.DB
)
