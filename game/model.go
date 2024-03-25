package game

// User 玩家信息表
type User struct {
	Uid  string `bson:"uid,omitempty"`  // 玩家的uid
	Name string `bson:"name,omitempty"` // 玩家的昵称
	//Password string `gorm:"not null"` // 玩家的密码
}

// Friendship 好友关系表
// SrcUid -> DstUid 是好友
// DstUid -> SrcUid 是好友
type Friendship struct {
	SrcUid string `bson:"src_uid,omitempty"`
	DstUid string `bson:"dst_uid,omitempty"`
}
