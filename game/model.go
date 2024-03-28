package game

// User 玩家信息表
type User struct {
	Uid  string `bson:"uid,omitempty"`  // 玩家的uid
	Name string `bson:"name,omitempty"` // 玩家的昵称
	//Password string `gorm:"not null"` // 玩家的密码
	TeamId string `bson:"team_id,omitempty"` // 玩家所在的队伍
}

// Friendship 好友关系表
// SrcUid -> DstUid 是好友
// DstUid -> SrcUid 是好友
type Friendship struct {
	SrcUid string `bson:"src_uid,omitempty"`
	DstUid string `bson:"dst_uid,omitempty"`
}

const DefaultTeamId = "Default"

type Team struct {
	Owner string `bson:"owner,omitempty"` // 队长的uid
	Id    string `bson:"uid,omitempty"`
	Name  string `bson:"name,omitempty"`
	Color string `bson:"color,omitempty"` // hex color
}
