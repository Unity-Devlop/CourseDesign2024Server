package game

import (
	"gorm.io/gorm"
)

// UserInfo 玩家信息表
// 存储一些基本信息: uid, 密码, 是否有游戏角色, 角色的名字
type UserInfo struct {
	gorm.Model
	Uid           uint32 `gorm:"primaryKey;unique"` // 玩家的uid(账号) 42亿对于我们这个游戏来说足够了
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
	SrcUid uint32 `gorm:"primaryKey"` //
	DstUid uint32 `gorm:"primaryKey"` //
}

// ShopItem 商店商品表
type ShopItem struct {
	gorm.Model
	ShopId       uint32      `gorm:"primaryKey;unique"` // 商店的id
	Id           uint32      `gorm:"primaryKey;unique"` // 在售商品的id
	Name         string      `gorm:"not null"`          // 在售商品的名字
	Desc         string      `gorm:"not null"`          // 在售商品的描述
	GoodsList    ItemArray   `gorm:"type:json"`         // 商品内容
	GoodsNumList Uint32Array `gorm:"type:json"`         // 商品内容数量
	Price        uint32      `gorm:"not null"`          // 在售商品的价格
}
