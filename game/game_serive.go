package game

import (
	pb "Server/proto"
	"fmt"
	"gorm.io/gorm"
)

type GameService struct {
	pb.UnimplementedGameServiceServer          // Rpc服务
	Db                                *gorm.DB // 游戏的数据库
	tickInterval                      uint32   // 定时器间隔
}

func NewGameService(db *gorm.DB) *GameService {
	return &GameService{
		Db: db,
	}
}

func (s *GameService) Run(tickInterval uint32) {
	s.tickInterval = tickInterval
	// 判断表是否存在 不存在则自动创建
	if !s.Db.Migrator().HasTable(&UserInfo{}) {
		// 创建表
		err := s.Db.Migrator().CreateTable(&UserInfo{})
		fmt.Printf("CreateTable UserInfo err: %v\n", err)
	}

	if !s.Db.Migrator().HasTable(&Friendship{}) {
		// 创建表
		err := s.Db.Migrator().CreateTable(&Friendship{})
		fmt.Printf("CreateTable Friendship err: %v\n", err)
	}
}
