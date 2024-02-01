package game

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	Uid           uint32  `gorm:"primaryKey;unique"` // 玩家的uid(账号) 42亿对于我们这个游戏来说足够了
	Password      string  // 密码
	HasCharacter  bool    // 是否创建过游戏角色
	CharacterName string  `gorm:"not null"`             // 角色名字
	Pos           Vector3 `json:"pos" gorm:"type:text"` // 角色位置
}

type Vector3 struct {
	x float32
	y float32
	z float32
}

func (vec *Vector3) Scan(value interface{}) error {
	b, _ := value.([]byte)
	return json.Unmarshal(b, &vec)
}

func (vec Vector3) Value() (driver.Value, error) {
	return json.Marshal(vec)
}
