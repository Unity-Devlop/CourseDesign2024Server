package main

import (
	"Server/game"
	pb "Server/proto"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
)

func main() {
	fmt.Printf("[Goland Server] init.\n")
	//sql.Open("sqlite3", "./game.db")
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("[Goland Server] connected to database.\n")

	// 开启rpc服务
	lis, err := net.Listen("tcp", ":44445")
	if err != nil {
		fmt.Printf("[Goland Server] failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	fmt.Printf("[Goland Server] start.\n")
	pb.RegisterGameServiceServer(s, &game.Server{
		Db: db,
	})
	fmt.Printf("[Goland Server] listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[Goland Server] failed to serve: %v\n", err)
	}
}
