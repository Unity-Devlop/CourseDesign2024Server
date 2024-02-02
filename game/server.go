package game

import (
	pb "Server/proto"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type Server struct {
	pb.UnimplementedGameServiceServer                                    // Rpc服务
	Db                                *gorm.DB                           // 游戏的数据库
	publicChat                        *BroadcastService[*pb.ChatMessage] // 公共聊天
	bubbleChat                        *BroadcastService[*pb.ChatMessage] // 泡泡聊天
}

func (s *Server) CreateCharacter(ctx context.Context, in *pb.CreateCharacterRequest) (*pb.CreateCharacterResponse, error) {
	var response pb.CreateCharacterResponse
	var errorMsg pb.ErrorMessage
	response.Error = &errorMsg
	errorMsg.Code = pb.StatusCode_ERROR

	var userInfo UserInfo
	result := s.Db.First(&userInfo, "uid = ?", in.Uid)
	if result.Error == gorm.ErrRecordNotFound {
		errorMsg.Msg = fmt.Sprintf("CreateCharacter: uid %d not found", in.Uid)
		fmt.Println(errorMsg.Msg)
		return &response, nil
	}
	if result.Error != nil {
		fmt.Printf("CreateCharacter: uid %d failed err:%v \n", in.Uid, result.Error)
		errorMsg.Msg = fmt.Sprintf("unknown error: %v", result.Error)
		return &response, nil
	}
	if userInfo.HasCharacter {
		errorMsg.Msg = fmt.Sprintf("CreateCharacter: uid %d already has character", in.Uid)
		fmt.Println(errorMsg.Msg)
		return &response, nil
	}
	// 更新数据库
	userInfo.HasCharacter = true
	userInfo.CharacterName = in.CharacterName
	userInfo.Pos.x = 0
	userInfo.Pos.y = 0
	userInfo.Pos.z = 0

	err := s.Db.Model(&userInfo).Updates(userInfo).Error
	//err := s.Db.Save(&userInfo).Error
	if err != nil {
		fmt.Printf("CreateCharacter: uid %d failed err:%v \n", in.Uid, err)
		errorMsg.Msg = fmt.Sprintf("unknown error: %v", err)
		return &response, nil
	}

	fmt.Printf("CreateCharacter: uid %d success\n", in.Uid)
	errorMsg.Code = pb.StatusCode_OK
	return &response, nil
}

func (s *Server) UserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	var response pb.UserInfoResponse
	var errorMsg pb.ErrorMessage
	response.Error = &errorMsg
	errorMsg.Code = pb.StatusCode_ERROR

	var userInfo UserInfo
	result := s.Db.First(&userInfo, "uid = ?", in.Uid)
	if result.Error == gorm.ErrRecordNotFound {
		errorMsg.Msg = fmt.Sprintf("UserInfo: uid %d not found", in.Uid)
		return &response, nil
	}
	if result.Error != nil {
		fmt.Printf("UserInfo: uid %d failed err:%v \n", in.Uid, result.Error)
		errorMsg.Msg = fmt.Sprintf("unknown error: %v", result.Error)
		return &response, nil
	}
	errorMsg.Code = pb.StatusCode_OK
	return &pb.UserInfoResponse{
		Error:         &errorMsg,
		Uid:           userInfo.Uid,
		HasCharacter:  userInfo.HasCharacter,
		CharacterName: userInfo.CharacterName,
		CharacterPosX: userInfo.Pos.x,
		CharacterPosY: userInfo.Pos.y,
		CharacterPosZ: userInfo.Pos.z,
	}, nil
}

func (s *Server) UserLogin(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	var response pb.UserLoginResponse
	response.Error = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}
	var info UserInfo
	result := s.Db.First(&info, "uid = ?", in.Uid)
	if result.Error != nil {
		fmt.Printf("UserLogin: uid %d failed err:%v \n", in.Uid, result.Error)
		return &response, nil
	}
	err := bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(in.Password))
	if err != nil {
		response.Error.Msg = fmt.Sprintf("UserLogin: uid %d password error", in.Uid)
		return &response, nil
	}
	fmt.Printf("UserLogin: uid %d with password %v success\n", in.Uid, in.Password)
	response.Error.Code = pb.StatusCode_OK
	response.Uid = info.Uid
	return &response, nil
}

func (s *Server) UserRegister(ctx context.Context, in *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	var response pb.UserRegisterResponse
	var errorMsg pb.ErrorMessage
	errorMsg.Code = pb.StatusCode_ERROR
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
		var reason string
		if !PasswordCheck(in.Password, &reason) {
			response.Error.Msg = reason
			return &response, nil
		}
		bytes, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("bcrypt.GenerateFromPassword(%s) err: %v\n", in.Password, err)
			return &response, err
		}
		// 插入数据库
		info = UserInfo{
			Uid:      in.Uid,
			Password: string(bytes),
			Pos: Vector3{
				x: 0,
				y: 0,
				z: 0,
			},
		}
		s.Db.Create(&info)

		fmt.Printf("UserRegister: uid %d success\n", in.Uid)
		errorMsg.Code = pb.StatusCode_OK
		response.Uid = info.Uid
		return &response, nil
	} else if result.Error != nil {
		fmt.Printf("UserRegister: uid %d failed err:%v \n", in.Uid, result.Error)
	}
	// 其他情况都是注册失败
	return &response, nil
}

func (s *Server) PublicChat(ctx context.Context, in *pb.ChatMessage) (*pb.ErrorMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublicChat not implemented")
}
func (s *Server) BubbleChat(ctx context.Context, in *pb.ChatMessage) (*pb.ErrorMessage, error) {
	//fmt.Printf("BubbleChat: uid %d msg %s\n", in.Uid, in.Msg)
	return nil, status.Errorf(codes.Unimplemented, "method BubbleChat not implemented")
}
func (s *Server) StartPublicChat(in *pb.ChatRequest, stream pb.GameService_StartPublicChatServer) error {
	return status.Errorf(codes.Unimplemented, "method StartPublicChat not implemented")
}
func (s *Server) StartBubbleChat(in *pb.ChatRequest, stream pb.GameService_StartBubbleChatServer) error {
	return status.Errorf(codes.Unimplemented, "method StartBubbleChat not implemented")
}
