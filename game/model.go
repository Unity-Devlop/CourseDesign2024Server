package game

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Uid           uint32 `gorm:"primaryKey;unique"` // 玩家的uid(账号) 42亿对于我们这个游戏来说足够了
	Password      string // 密码
	HasCharacter  bool   // 是否创建过游戏角色
	CharacterName string `gorm:"not null"` // 角色名字
	//Pos           Vector3 `json:"pos" gorm:"type:text"` // 角色位置
}

// Friendship 好友关系表
// SrcUid -> DstUid 是好友
// DstUid -> SrcUid 是好友
type Friendship struct {
	gorm.Model
	SrcUid uint32 `gorm:"primaryKey;unique"` //
	DstUid uint32 `gorm:"primaryKey;unique"` //
}
