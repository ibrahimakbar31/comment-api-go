package db

import "github.com/jinzhu/gorm"

//DB1 postgres
type DB1 struct {
	*gorm.DB
}
