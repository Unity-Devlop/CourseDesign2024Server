package game

import (
	pb "Server/proto"
	"fmt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedGameServiceServer          // Rpc服务
	Db                                *gorm.DB // 游戏的数据库
}

func (s *Server) PlayerInfo(ctx context.Context, request *pb.PlayerInfoRequest) (*pb.PlayerInfoResponse, error) {
	//TODO implement me
	var info PlayerInfo
	// 查询数据库
	//result := s.Db.First(&info, request.Uid)
	result := s.Db.First(&info, "uid = ?", request.Uid)

	if result.Error != nil {
		fmt.Printf("PlayerInfo Query Error: %v\n", result.Error)
		return &pb.PlayerInfoResponse{Exist: false}, nil
	}
	// 返回数据
	return &pb.PlayerInfoResponse{
		Exist: true,
		Uid:   info.Uid,
		Name:  info.Name,
	}, nil

}

func (s *Server) PlayerLogin(ctx context.Context, request *pb.PlayerLoginRequest) (*pb.PlayerLoginResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) PlayerRegister(ctx context.Context, request *pb.PlayerRegisterRequest) (*pb.PlayerRegisterResponse, error) {
	//TODO implement me
	var response pb.PlayerRegisterResponse
	response.Success = false

	if request.Uid == 0 {
		fmt.Printf("PlayerRegister: uid %d error\n", request.Uid)
		return &response, nil
	}

	// 判断表是否存在 不存在则自动创建
	if !s.Db.Migrator().HasTable(&PlayerInfo{}) {
		// 创建表
		err := s.Db.Migrator().CreateTable(&PlayerInfo{})
		if err != nil {
			return &response, err
		}
	}

	var info PlayerInfo
	// Uid 相同
	result := s.Db.First(&info, "uid = ?", request.Uid)
	// 只有在找不到记录的时候才插入 -> 允许注册
	if result.Error == gorm.ErrRecordNotFound {
		// 插入数据库
		info = PlayerInfo{Uid: request.Uid, Name: request.Name}
		s.Db.Create(&info)

		fmt.Printf("PlayerRegister: uid %d success\n", request.Uid)
		response.Success = true
		response.Uid = info.Uid
		return &response, nil
	}

	fmt.Printf("Already Exist: uid %d\nError:%s", request.Uid, result.Error)
	return &response, nil
}
