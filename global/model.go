package global

import "time"

type GVA_MODEL struct {
	ID          uint `gorm:"primarykey"`
	CreatedTime time.Time
	UpdateTime  time.Time
}
