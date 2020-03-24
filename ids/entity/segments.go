package entity

import "time"

type Segments struct {
	BizTag      string `gorm:"type:varchar(32);not null;primary_key"`
	MaxId       int64  `gorm:"type:bigint(20);not null"`
	Step        int64  `gorm:"type:bigint(20);not null"`
	Description string `gorm:"type:varchar(1024);not null"`
	UpdateTime  time.Time
}
