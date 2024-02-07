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

func (s *Server) CreateCharacter(ctx context.Context, in *pb.CreateCharacterRequest) (*pb.CreateCharacterResponse, error) {
	var response pb.CreateCharacterResponse
	var errorMsg pb.ErrorMessage
	response.Error = &errorMsg
	response.Uid = in.Uid
	errorMsg.Code = pb.StatusCode_ERROR

	var userInfo UserInfo
	result := s.Db.First(&userInfo, "uid = ?", in.Uid)
	if result.Error != nil {
		errorMsg.Msg = fmt.Sprintf("用户[%d]创建角色失败,err:%v", in.Uid, result.Error)
		fmt.Println(errorMsg.Msg)
		return &response, nil
	}

	if userInfo.HasCharacter {
		errorMsg.Msg = fmt.Sprintf("用户:[%d]已经创建过角色", in.Uid)
		fmt.Println(errorMsg.Msg)
		return &response, nil
	}
	var reason string
	if !CharacterNameCheck(in.CharacterName, &reason) {
		response.Error.Msg = reason
		return &response, nil
	}

	// 更新数据库
	userInfo.HasCharacter = true
	userInfo.CharacterName = in.CharacterName

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
		fmt.Printf(" PublicChat: uid %d already in chat now close the pre one\n", in.Uid)
		s.publicChat.UnListen(*s.uid2publicChat[in.Uid])
	}
	fmt.Printf("Start PublicChat Channel: uid %d\n", in.Uid)
	// 注册广播
	var pushChan = s.publicChat.Listen()
	// 记录这个通道
	s.uid2publicChat[in.Uid] = &pushChan
	defer func() {
		s.publicChat.UnListen(pushChan)
		fmt.Printf("End PublicChat Channel: uid %d\n", in.Uid)
	}()
	return s.startChat("PublicChat", in.Uid, pushChan, stream)
}

func (s *Server) StartBubbleChat(in *pb.ChatRequest, stream pb.GameService_StartBubbleChatServer) error {
	//
	if s.uid2bubbleChat[in.Uid] != nil {
		fmt.Printf("Start BubbleChat: uid %d already in chat now close the pre one\n", in.Uid)
		s.bubbleChat.UnListen(*s.uid2bubbleChat[in.Uid])
	}
	fmt.Printf("Start BubbleChat Channel: uid %d\n", in.Uid)
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

func (s *Server) StopPublicChat(ctx context.Context, in *pb.ChatRequest) (*pb.ErrorMessage, error) {
	if s.uid2publicChat[in.Uid] != nil {
		fmt.Printf("Stop PublicChat: uid %d\n", in.Uid)
		s.publicChat.UnListen(*s.uid2publicChat[in.Uid])
		delete(s.uid2publicChat, in.Uid)
		return &pb.ErrorMessage{
			Code: pb.StatusCode_OK,
		}, nil
	}
	return &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
		Msg:  fmt.Sprintf("StopPublicChat: uid %d not in chat", in.Uid),
	}, nil
}

func (s *Server) StopBubbleChat(ctx context.Context, in *pb.ChatRequest) (*pb.ErrorMessage, error) {
	if s.uid2bubbleChat[in.Uid] != nil {
		fmt.Printf("Stop BubbleChat: uid %d\n", in.Uid)
		s.bubbleChat.UnListen(*s.uid2bubbleChat[in.Uid])
		delete(s.uid2bubbleChat, in.Uid)
		return &pb.ErrorMessage{
			Code: pb.StatusCode_OK,
		}, nil
	}
	return &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
		Msg:  fmt.Sprintf("StopBubbleChat: uid %d not in chat", in.Uid),
	}, nil
}

func (s *Server) SearchFriend(ctx context.Context, in *pb.SearchFriendRequest) (*pb.SearchFriendResponse, error) {

	var errorMsg = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}
	var response = &pb.SearchFriendResponse{
		Error: errorMsg,
	}

	if in.SearcherUid == in.TargetUid {
		// 不能搜索自己
		errorMsg.Msg = fmt.Sprintf("不能搜索自己 uid:%d", in.SearcherUid)
		return response, nil
	}

	var srcUser UserInfo
	var dstUser UserInfo

	if s.Db.First(&srcUser, "uid = ?", in.SearcherUid).Error != nil {
		// 进行搜索的玩家不存在
		errorMsg.Msg = fmt.Sprintf("进行搜索的玩家不存在 uid:%d", in.SearcherUid)
		return response, nil
	}

	if !srcUser.HasCharacter {
		// 进行搜索的玩家没有角色
		errorMsg.Msg = fmt.Sprintf("进行搜索的玩家没有角色 uid:%d", in.SearcherUid)
		return response, nil
	}

	if s.Db.First(&dstUser, "uid = ?", in.TargetUid).Error != nil {
		// 搜索的目标玩家不存在
		errorMsg.Msg = fmt.Sprintf("搜索的目标玩家不存在 uid:%d", in.TargetUid)
		return response, nil
	}

	if !dstUser.HasCharacter {
		// 搜索的目标玩家没有角色
		errorMsg.Msg = fmt.Sprintf("搜索的目标玩家没有角色 uid:%d", in.TargetUid)
		return response, nil
	}

	// 到这里说明两个玩家都存在 填充返回值
	errorMsg.Code = pb.StatusCode_OK

	response.SearcherUid = srcUser.Uid
	response.Name = dstUser.CharacterName
	response.TargetUid = dstUser.Uid

	var friendship Friendship
	// 为了避免重复存储 保证小的uid在前面
	var smallId = min(in.SearcherUid, in.TargetUid)
	var bigId = max(in.SearcherUid, in.TargetUid)
	// 看下两个人是否有好友关系
	if s.Db.First(&friendship, "src_uid = ? AND dst_uid = ?", smallId, bigId).Error != nil {
		response.IsFriend = false
	} else {
		response.IsFriend = true
	}

	return response, nil
}

func (s *Server) AddFriend(ctx context.Context, in *pb.AddFriendRequest) (*pb.ErrorMessage, error) {
	var response = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}

	if in.SrcUid == in.TargetUid {
		response.Msg = fmt.Sprintf("不能添加自己为好友 uid:%d", in.SrcUid)
		return response, nil
	}

	// 检查两个用户是否存在
	if s.Db.First(&UserInfo{}, "uid = ?", in.SrcUid).Error != nil {
		response.Msg = fmt.Sprintf("添加好友,发起用户[%d]不存在", in.SrcUid)
		return response, nil
	}
	if s.Db.First(&UserInfo{}, "uid = ?", in.TargetUid).Error != nil {
		response.Msg = fmt.Sprintf("添加好友,目标用户[%d]不存在", in.TargetUid)
		return response, nil
	}

	// 判断是否已经是好友
	var friendship Friendship
	// 为了避免重复存储 保证小的uid在前面
	var smallId = min(in.SrcUid, in.TargetUid)
	var bigId = max(in.SrcUid, in.TargetUid)
	// 看下两个人是否有好友关系
	err := s.Db.First(&friendship, "src_uid = ? AND dst_uid = ?", smallId, bigId).Error
	if err == nil {
		response.Msg = fmt.Sprintf("添加好友,用户[%d]和用户[%d]已经是好友", in.SrcUid, in.TargetUid)
		return response, nil
	}

	// 添加好友
	friendship = Friendship{
		SrcUid: smallId,
		DstUid: bigId,
	}

	if s.Db.Create(&friendship).Error != nil {
		response.Msg = fmt.Sprintf("添加好友,用户[%d]和用户[%d]添加好友失败", in.SrcUid, in.TargetUid)
		return response, nil
	}

	response.Code = pb.StatusCode_OK
	return response, nil
}

func (s *Server) FriendList(ctx context.Context, in *pb.FriendListRequest) (*pb.FriendListResponse, error) {

	var errorMsg = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}
	var response = &pb.FriendListResponse{
		Error: errorMsg,
	}

	if s.Db.First(&UserInfo{}, "uid = ?", in.Uid).Error != nil {
		errorMsg.Msg = fmt.Sprintf("用户[%d]不存在", in.Uid)
		return response, nil
	}

	errorMsg.Code = pb.StatusCode_OK
	response.Uid = in.Uid

	// 查询好友列表
	// 首先查FriendShip表中src_uid为in.Uid的记录
	// 然后查dst_uid为in.Uid的记录
	// 合并两个结果
	var friendships []Friendship
	var friendUidList []uint32
	var friendNameList []string

	// 查找src_uid为in.Uid的记录
	s.Db.Model(&Friendship{}).Find(&friendships, "src_uid = ?", in.Uid)
	for _, friendship := range friendships {
		if friendship.DstUid == in.Uid {
			continue
		}
		friendUidList = append(friendUidList, friendship.DstUid)
	}
	// 查找dst_uid为in.Uid的记录
	s.Db.Model(&Friendship{}).Find(&friendships, "dst_uid = ?", in.Uid)
	for _, friendship := range friendships {
		if friendship.DstUid == in.Uid {
			continue
		}
		friendUidList = append(friendUidList, friendship.DstUid)
	}
	// 不需要去重 因为插入的时候保证 src_uid < dst_uid

	// 查询好友的名字
	for _, uid := range friendUidList {
		if uid == in.Uid {
			continue
		}
		var userInfo UserInfo
		s.Db.First(&userInfo, "uid = ?", uid)
		friendNameList = append(friendNameList, userInfo.CharacterName)
	}

	response.FriendUidList = friendUidList
	response.FriendNameList = friendNameList

	return response, nil
}

func (s *Server) RemoveFriend(ctx context.Context, in *pb.RemoveFriendRequest) (*pb.ErrorMessage, error) {
	var response = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}

	if in.SrcUid == in.TargetUid {
		response.Msg = fmt.Sprintf("不能删除自己 uid:%d", in.SrcUid)
		return response, nil
	}

	var friendship Friendship
	// 为了避免重复存储 保证小的uid在前面
	var smallId = min(in.SrcUid, in.TargetUid)
	var bigId = max(in.SrcUid, in.TargetUid)

	// 看下两个人是否有好友关系
	if s.Db.First(&friendship, "src_uid = ? AND dst_uid = ?", smallId, bigId).Error != nil {
		response.Msg = fmt.Sprintf("用户[%d]和用户[%d]不是好友", in.SrcUid, in.TargetUid)
		return response, nil
	}
	// 删除好友
	if s.Db.Model(&Friendship{}).Delete(&friendship, "src_uid = ? AND dst_uid = ?", smallId, bigId).Error != nil {
		response.Msg = fmt.Sprintf("用户[%d]和用户[%d]删除好友失败", in.SrcUid, in.TargetUid)
		return response, nil
	}

	response.Code = pb.StatusCode_OK

	return response, nil
}
