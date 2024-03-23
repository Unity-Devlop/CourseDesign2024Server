package game

import (
	pb "Server/proto"
	"gorm.io/gorm"
)

type GlobalService struct {
	pb.UnimplementedGameServiceServer          // Rpc服务
	Db                                *gorm.DB // 游戏的数据库
	tickInterval                      uint32   // 定时器间隔
}

func NewGlobalService(db *gorm.DB) *GlobalService {
	return &GlobalService{
		Db: db,
	}
}
