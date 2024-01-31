package game

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	Name     string
	Uid      uint32 `gorm:"primaryKey;unique"` // 42亿足够了
	Password string
}

type Position struct {
	x float32
	y float32
	z float32
}
