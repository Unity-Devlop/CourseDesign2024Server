package game

import (
	"gorm.io/gorm"
)

// User 玩家信息表
type User struct {
	gorm.Model
	Uid  string `gorm:"primaryKey;unique"` // 玩家的uid
	Name string `gorm:"unique"`            // 玩家的昵称
	//Password string `gorm:"not null"` // 玩家的密码
}

// Friendship 好友关系表
// SrcUid -> DstUid 是好友
// DstUid -> SrcUid 是好友
type Friendship struct {
	gorm.Model
	SrcUid string `gorm:"primaryKey"` //
	DstUid string `gorm:"primaryKey"` //
}
