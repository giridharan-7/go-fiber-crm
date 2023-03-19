package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// it helps golang to interact with database
	DBConn *gorm.DB
)
