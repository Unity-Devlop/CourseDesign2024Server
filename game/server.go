package game

import (
	pb "Server/proto"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"runtime"
	"time"
)

type Server struct {
	pb.UnimplementedGameServiceServer                                    // Rpc服务
	Db                                *gorm.DB                           // 游戏的数据库
	publicChat                        *BroadcastService[*pb.ChatMessage] // 公共聊天
	bubbleChat                        *BroadcastService[*pb.ChatMessage] // 泡泡聊天
	uid2publicChat                    map[uint32]*chan *pb.ChatMessage   // uid -> 公共聊天
	uid2bubbleChat                    map[uint32]*chan *pb.ChatMessage   // uid -> 泡泡聊天
	tickInterval                      uint32                             // 定时器间隔
}

func NewServer(db *gorm.DB) *Server {
	return &Server{
		Db:             db,
		publicChat:     NewBroadcastService[*pb.ChatMessage](),
		bubbleChat:     NewBroadcastService[*pb.ChatMessage](),
		uid2publicChat: make(map[uint32]*chan *pb.ChatMessage),
		uid2bubbleChat: make(map[uint32]*chan *pb.ChatMessage),
	}
}

func (s *Server) Run(tickInterval uint32) {
	s.tickInterval = tickInterval
	s.publicChat.Run(tickInterval)
	s.bubbleChat.Run(tickInterval)
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
	fmt.Printf("PublicChat: uid %d msg %s\n", in.Uid, in.Msg)
	// 广播消息
	s.publicChat.Send(in)
	fmt.Printf("PublicChat: uid %d msg %s success\n", in.Uid, in.Msg)
	return &pb.ErrorMessage{
		Code: pb.StatusCode_OK,
	}, nil
}

func (s *Server) BubbleChat(ctx context.Context, in *pb.ChatMessage) (*pb.ErrorMessage, error) {
	fmt.Printf("BubbleChat: uid %d msg %s\n", in.Uid, in.Msg)
	// 广播消息
	s.bubbleChat.Send(in)
	fmt.Printf("BubbleChat: uid %d msg %s success\n", in.Uid, in.Msg)
	return &pb.ErrorMessage{
		Code: pb.StatusCode_OK,
	}, nil
}

func (s *Server) StartPublicChat(in *pb.ChatRequest, stream pb.GameService_StartPublicChatServer) error {
	if s.uid2publicChat[in.Uid] != nil {
		fmt.Printf("Start PublicChat: uid %d already in chat\n", in.Uid)
		fmt.Printf("End PublicChat: uid %d\n", in.Uid)
		s.publicChat.UnListen(*s.uid2publicChat[in.Uid])
	}
	fmt.Printf("Start PublicChat: uid %d\n", in.Uid)
	// 注册广播
	var pushChan = s.publicChat.Listen()
	// 记录这个通道
	s.uid2publicChat[in.Uid] = &pushChan
	defer func() {
		s.publicChat.UnListen(pushChan)
		fmt.Printf("End PublicChat: uid %d\n", in.Uid)
	}()
	return s.startChat("PublicChat", in.Uid, pushChan, stream)
}

func (s *Server) StartBubbleChat(in *pb.ChatRequest, stream pb.GameService_StartBubbleChatServer) error {
	//
	if s.uid2bubbleChat[in.Uid] != nil {
		fmt.Printf("Start BubbleChat: uid %d already in chat\n", in.Uid)
		fmt.Printf("End BubbleChat: uid %d\n", in.Uid)
		s.bubbleChat.UnListen(*s.uid2bubbleChat[in.Uid])
	}
	fmt.Printf("Start BubbleChat: uid %d\n", in.Uid)
	// 注册广播
	var pushChan = s.bubbleChat.Listen()
	// 记录这个通道
	s.uid2bubbleChat[in.Uid] = &pushChan

	defer func() {
		s.bubbleChat.UnListen(pushChan)
		fmt.Printf("End BubbleChat: uid %d\n", in.Uid)
	}()
	return s.startChat("BubbleChat", in.Uid, pushChan, stream)
}

func (s *Server) startChat(name string, uid uint32, c chan *pb.ChatMessage, stream pb.GameService_StartBubbleChatServer) error {
	defer func() {
		fmt.Printf("%s chat end: uid %d\n", name, uid)
	}()
	for {
		select {
		case msg, ok := <-c:
			if !ok {
				fmt.Printf("%s chat closed: uid %d\n", name, uid)
				return nil
			}
			if msg.Uid == uid {
				// 不给自己发消息
				runtime.Gosched()
				continue
			}
			err := stream.Send(msg)
			if err != nil {
				fmt.Printf("%s chat error: uid %d failed err:%v \n", name, uid, err)
				return nil
			}
		default:
			// 没有消息 休息一下
			time.Sleep(time.Duration(s.tickInterval) * time.Millisecond)
		}
	}
}
