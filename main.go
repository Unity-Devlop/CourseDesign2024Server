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

type Config struct {
	Dsn string `json:"dsn"`
}

func getConnection() gorm.Dialector {
	//sql.Open("sqlite3", "./game.db")
	// 连接数据库
	// 加载配置文件
	return sqlite.Open("game.db")
	//
	//var config Config       // MySQL配置
	//var conn gorm.Dialector // gorm连接
	//conn = nil
	//file, err := os.Open("config.json")
	//
	//if err != nil {
	//	return sqlite.Open("game.db")
	//}
	//defer func(file *os.File) {
	//	err := file.Close()
	//	if err != nil {
	//		fmt.Printf("[Goland Server] failed to close file: %v\n", err)
	//	}
	//}(file)
	//
	//// 解析json
	//decoder := json.NewDecoder(file)
	//err = decoder.Decode(&config)
	//if err != nil {
	//	return sqlite.Open("game.db")
	//}
	//
	//conn = mysql.Open(config.Dsn) // mysql
	//
	//return conn
}

func main() {
	fmt.Printf("[Goland Server] init.\n")
	conn := getConnection()

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("[Goland Server] connected to database.\n")

	// 开启rpc服务
	lis, err := net.Listen("tcp", ":22333")
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			fmt.Printf("[Goland Server] failed to close: %v\n", err)
		}
	}(lis)

	if err != nil {
		fmt.Printf("[Goland Server] failed to listen: %v\n", err)
	}
	s := grpc.NewServer()
	fmt.Printf("[Goland Server] start.\n")

	pb.RegisterGameServiceServer(s,
		&game.Server{
			Db: db,
		},
	)
	fmt.Printf("[Goland Server] listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("[Goland Server] failed to serve: %v\n", err)
	}
}
