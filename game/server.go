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
	result := s.Db.First(&info, request.Uid)
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
	var info PlayerInfo
	result := s.Db.First(&info, request.Uid)
	if result.Error != nil {
		fmt.Printf("PlayerRegister Query Error: %v\n", result.Error)
		return &pb.PlayerRegisterResponse{Success: false}, nil
	}
	// 插入数据库
	info = PlayerInfo{Uid: request.Uid, Name: request.Name}
	s.Db.Create(&info)

	fmt.Printf("PlayerRegister: uid %d success\n", request.Uid)

	return &pb.PlayerRegisterResponse{
		Success: true,
		Uid:     request.Uid,
	}, nil
}
