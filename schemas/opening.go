package schemas

import (
	"gorm.io/gorm"
)

// parece model de laravel
type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
