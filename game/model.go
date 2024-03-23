package game

import (
	"gorm.io/gorm"
)

type UserId string

// UserInfo 玩家信息表
// 存储一些基本信息: uid, 密码, 是否有游戏角色, 角色的名字
type UserInfo struct {
	gorm.Model
	Uid           UserId `gorm:"primaryKey;unique"` // 玩家的uid(账号) 42亿对于我们这个游戏来说足够了
	Password      string // 密码
	HasCharacter  bool   // 是否创建过游戏角色
	CharacterName string `gorm:"not null"` // 角色名字
	Coin          uint64 `gorm:"not null"` // 玩家的金币数量
}

// Friendship 好友关系表
// SrcUid -> DstUid 是好友
// DstUid -> SrcUid 是好友
type Friendship struct {
	gorm.Model
	SrcUid UserId `gorm:"primaryKey"` //
	DstUid UserId `gorm:"primaryKey"` //
}
