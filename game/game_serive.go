package game

import (
	pb "Server/proto"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GameService struct {
	pb.UnimplementedGameServiceServer                 // Rpc服务
	Db                                *mongo.Database // mongodb
	tickInterval                      uint32          // 定时器间隔
	userCollection                    *mongo.Collection
	friendshipCollection              *mongo.Collection
}

func NewGameService(db *mongo.Database) *GameService {
	return &GameService{
		Db: db,
	}
}

func (s *GameService) Run() {
}

func (s *GameService) ContainsUser(ctx context.Context, request *pb.StringMessage) (*pb.ErrorMessage, error) {
	var response = &pb.ErrorMessage{Code: pb.StatusCode_ERROR}

	if s.Db.Collection("user").FindOne(context.TODO(), bson.D{{"uid", request.Content}}).Err() != nil {
		response.Content = fmt.Sprintf("用户[%s]不存在", request.Content)
		return response, nil
	}
	response.Code = pb.StatusCode_OK
	return response, nil
}

func (s *GameService) GetUid(context.Context, *pb.GetUidRequest) (*pb.UidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUid not implemented")
}

func (s *GameService) RegisterUser(ctx context.Context, request *pb.RegisterRequest) (*pb.ErrorMessage, error) {
	var response = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}
	if request.Uid == "" {
		response.Content = "uid不能为空"
		return response, nil
	}

	if request.Name == "" {
		response.Content = "昵称不能为空"
		return response, nil
	}
	userCollection := s.Db.Collection("user")

	var user User
	filter := bson.D{{"uid", request.Uid}}
	if userCollection.FindOne(context.Background(), filter).Decode(&user) == nil {
		response.Content = fmt.Sprintf("用户已经存在 uid:%s", request.Uid)
		return response, nil
	}

	user = User{
		Uid:  request.Uid,
		Name: request.Name,
	}

	if _, err := userCollection.InsertOne(context.TODO(), user); err != nil {
		response.Content = fmt.Sprintf("创建用户失败 uid:%s", request.Uid)
		return response, nil
	}

	response.Code = pb.StatusCode_OK
	fmt.Printf("用户注册成功 uid:%s name:%s\n", request.Uid, request.Name)
	return response, nil
}

func (s *GameService) GetFriendList(ctx context.Context, request *pb.FriendListRequest) (*pb.FriendShipList, error) {
	var errorMsg = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}
	var response = &pb.FriendShipList{
		Error: errorMsg,
	}

	var user User
	if s.Db.Collection("user").FindOne(context.TODO(), bson.D{{"uid", request.Uid}}).Decode(&user) != nil {
		errorMsg.Content = fmt.Sprintf("用户[%s]不存在", request.Uid)
		return response, nil

	}

	errorMsg.Code = pb.StatusCode_OK
	// 查询好友列表
	// 首先查FriendShip表中src_uid为in.Uid的记录
	// 然后查dst_uid为in.Uid的记录
	// 合并两个结果
	var friendships []Friendship
	response.List = make([]*pb.FriendInfo, 0)

	// 查找src_uid为in.Uid的记录
	find, err := s.Db.Collection("friendship").Find(context.TODO(), bson.D{{"src_uid", request.Uid}})
	if err != nil {
		return response, err
	}
	if find.All(context.TODO(), &friendships) != nil {
		return response, err
	}
	for _, friendship := range friendships {
		if friendship.DstUid == request.Uid {
			continue
		}
		response.List = append(response.List, &pb.FriendInfo{
			Uid: friendship.DstUid,
		})
	}
	// 查找dst_uid为in.Uid的记录
	find, err = s.Db.Collection("friendship").Find(context.TODO(), bson.D{{"dst_uid", request.Uid}})
	if err != nil {
		return response, err
	}
	if find.All(context.TODO(), &friendships) != nil {
		return response, err
	}
	for _, friendship := range friendships {
		if friendship.SrcUid == request.Uid {
			continue
		}
		response.List = append(response.List, &pb.FriendInfo{
			Uid: friendship.SrcUid,
		})
	}
	// 不需要去重 因为插入的时候保证 src_uid < dst_uid
	// 查Name
	for _, friendInfo := range response.List {
		var friend User
		err := s.Db.Collection("user").FindOne(context.TODO(), bson.D{{"uid", friendInfo.Uid}}).Decode(&friend)
		if err != nil {
			return response, err
		}
		friendInfo.Name = friend.Name
	}
	return response, nil
}

func (s *GameService) AddFriend(ctx context.Context, request *pb.AddFriendRequest) (*pb.ErrorMessage, error) {
	var response = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}

	if request.SenderUid == request.TargetUid {
		response.Content = fmt.Sprintf("不能添加自己为好友 uid:%s", request.SenderUid)
		return response, nil
	}
	if s.Db.Collection("user").FindOne(context.TODO(), bson.D{{"uid", request.SenderUid}}).Err() != nil {
		response.Content = fmt.Sprintf("发起用户[%s]不存在", request.SenderUid)
		return response, nil
	}
	if s.Db.Collection("user").FindOne(context.TODO(), bson.D{{"uid", request.TargetUid}}).Err() != nil {
		response.Content = fmt.Sprintf("目标用户[%s]不存在", request.TargetUid)
		return response, nil
	}

	var friendship Friendship
	// 为了避免重复存储 保证小的uid在前面
	var smallId = min(request.SenderUid, request.TargetUid)
	var bigId = max(request.SenderUid, request.TargetUid)
	// 看下两个人是否有好友关系
	if s.Db.Collection("friendship").FindOne(context.TODO(), bson.D{{"src_uid", smallId}, {"dst_uid", bigId}}).Decode(&friendship) == nil {
		response.Content = fmt.Sprintf("用户[%s]和用户[%s]已经是好友", request.SenderUid, request.TargetUid)
		return response, nil

	}
	// 添加好友

	friendship = Friendship{
		SrcUid: smallId,
		DstUid: bigId,
	}
	_, err := s.Db.Collection("friendship").InsertOne(context.TODO(), friendship)
	if err != nil {
		response.Content = fmt.Sprintf("用户[%s]和用户[%s]添加好友失败", request.SenderUid, request.TargetUid)
		return response, nil
	}

	response.Code = pb.StatusCode_OK
	fmt.Printf("用户[%s]和用户[%s]添加好友成功\n", request.SenderUid, request.TargetUid)
	return response, nil
}

func (s *GameService) DeleteFriend(ctx context.Context, resutst *pb.DeleteFriendRequest) (*pb.ErrorMessage, error) {
	var response = &pb.ErrorMessage{
		Code: pb.StatusCode_ERROR,
	}
	if resutst.SenderUid == resutst.TargetUid {
		response.Content = fmt.Sprintf("不能删除自己 uid:%s", resutst.SenderUid)
		return response, nil
	}

	var friendship Friendship
	// 为了避免重复存储 保证小的uid在前面
	var smallId = min(resutst.SenderUid, resutst.TargetUid)
	var bigId = max(resutst.SenderUid, resutst.TargetUid)
	// 看下两个人是否有好友关系
	if s.Db.Collection("friendship").FindOne(context.TODO(), bson.D{{"src_uid", smallId}, {"dst_uid", bigId}}).Decode(&friendship) != nil {
		response.Content = fmt.Sprintf("用户[%s]和用户[%s]不是好友", resutst.SenderUid, resutst.TargetUid)
		return response, nil
	}
	// 删除好友
	if s.Db.Collection("friendship").FindOneAndDelete(context.TODO(), bson.D{{"src_uid", smallId}, {"dst_uid", bigId}}).Decode(&friendship) != nil {
		response.Content = fmt.Sprintf("用户[%s]和用户[%s]删除好友失败", resutst.SenderUid, resutst.TargetUid)
		return response, nil
	}
	response.Code = pb.StatusCode_OK
	fmt.Printf("用户[%s]和用户[%s]删除好友成功\n", resutst.SenderUid, resutst.TargetUid)
	return response, nil
}

func (s *GameService) SearchFriend(ctx context.Context, request *pb.SearchFriendRequest) (*pb.SearchFriendResponse, error) {
	var errorMsg = &pb.ErrorMessage{
		Code:    pb.StatusCode_ERROR,
		Content: "未知错误",
	}
	var response = &pb.SearchFriendResponse{
		Error: errorMsg,
	}

	if request.SearcherUid == request.TargetUid {
		errorMsg.Content = fmt.Sprintf("不能搜索自己 uid:%s", request.SearcherUid)
		return response, nil
	}

	var senderUser User
	var targetUser User
	userCollection := s.Db.Collection("user")
	if userCollection.FindOne(context.TODO(), bson.D{{"uid", request.SearcherUid}}).Decode(&senderUser) != nil {
		errorMsg.Content = fmt.Sprintf("进行搜索的玩家不存在 uid:%s", request.SearcherUid)
		return response, nil

	}

	if userCollection.FindOne(context.TODO(), bson.D{{"uid", request.TargetUid}}).Decode(&targetUser) != nil {
		errorMsg.Content = fmt.Sprintf("搜索的目标玩家不存在 uid:%s", request.TargetUid)
		return response, nil
	}

	errorMsg.Code = pb.StatusCode_OK
	response.TargetUid = request.TargetUid
	response.TargetName = targetUser.Name

	var friendship Friendship
	// 为了避免重复存储 保证小的uid在前面
	var smallId = min(request.SearcherUid, request.TargetUid)
	var bigId = max(request.SearcherUid, request.TargetUid)
	// 看下两个人是否有好友关系
	if s.Db.Collection("friendship").FindOne(context.TODO(), bson.D{{"src_uid", smallId}, {"dst_uid", bigId}}).Decode(&friendship) == nil {
		response.IsFriend = true
		return response, nil
	} else {
		response.IsFriend = false
		return response, nil
	}
}
