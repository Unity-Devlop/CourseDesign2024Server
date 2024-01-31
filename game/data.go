package game

import "gorm.io/gorm"

type PlayerInfo struct {
	gorm.Model
	Name     string
	Uid      uint32 `gorm:"primaryKey"` // 42亿足够了
	position Position
}

type Position struct {
	x float32
	y float32
	z float32
}
