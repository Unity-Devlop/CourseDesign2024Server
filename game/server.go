package game

import (
	pb "Server/proto"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedGameServiceServer          // Rpc服务
	Db                                *gorm.DB // 游戏的数据库
}

func (s *Server) UserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	panic("implement me")
}

func (s *Server) UserLogin(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	var response pb.UserLoginResponse
	response.Error = &pb.Error{
		Code: pb.ErrorCode_ERROR,
	}
	var info UserInfo
	result := s.Db.First(&info, "uid = ?", in.Uid)
	if result.Error != nil {
		return &response, nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(in.Password))
	if err != nil {

		response.Error.Msg = fmt.Sprintf("UserLogin: uid %d password error", in.Uid)
		return &response, nil
	}

	response.Error.Code = pb.ErrorCode_OK
	response.Uid = info.Uid
	return &response, nil
}

func (s *Server) UserRegister(ctx context.Context, in *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	var response pb.UserRegisterResponse
	var errorMsg pb.Error
	errorMsg.Code = pb.ErrorCode_ERROR
	response.Error = &errorMsg

	// 判断表是否存在 不存在则自动创建
	if !s.Db.Migrator().HasTable(&UserInfo{}) {
		// 创建表
		err := s.Db.Migrator().CreateTable(&UserInfo{})
		fmt.Printf("UserRegister: create table UserInfo\n")
		if err != nil {
			return &response, err
		}
	}

	var info UserInfo
	// Uid 相同
	result := s.Db.First(&info, "uid = ?", in.Uid)

	// 只有在找不到记录的时候才插入 -> 允许注册
	if result.Error == gorm.ErrRecordNotFound {
		if !PasswordCheck(in.Password) {
			return &response, nil
		}
		bytes, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return &response, err
		}
		// 插入数据库
		info = UserInfo{
			Uid:      in.Uid,
			Password: string(bytes),
		}
		s.Db.Create(&info)

		fmt.Printf("UserRegister: uid %d success\n", in.Uid)
		errorMsg.Code = pb.ErrorCode_OK
		response.Uid = info.Uid
		return &response, nil
	}
	// 其他情况都是注册失败
	return &response, nil

}
