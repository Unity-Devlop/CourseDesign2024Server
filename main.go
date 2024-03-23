package main

import (
	"Server/game"
	pb "Server/proto"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net"
	"os"
)

type dsnConfig struct {
	Dsn string `json:"dsn"`
}

func getMySqlDB() *gorm.DB {
	var config dsnConfig    // MySQL配置
	var conn gorm.Dialector // gorm连接
	file, _ := os.Open("config.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("[Service] failed to close file: %v\n", err)
		}
	}(file)

	// 解析json
	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&config)
	fmt.Printf("[Service] dsn: %s\n", config.Dsn)
	conn = mysql.Open(config.Dsn) // mysql
	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Printf("[Service] connected to database.\n")
	return db
}

func getSqliteDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func StartGameService() {
	fmt.Printf("[ Service] init.\n")
	db := getMySqlDB()
	//db := getSqliteDB()
	// 开启rpc服务
	lis, err := net.Listen("tcp", ":22333")
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			fmt.Printf("[Service] failed to close: %v\n", err)
		}
	}(lis)

	if err != nil {
		fmt.Printf("[Service] failed to listen: %v\n", err)
	}
	server := grpc.NewServer()
	fmt.Printf("[Service] start.\n")

	fmt.Printf("[Service] start game service.\n")
	gameService := game.NewGameService(db)
	gameService.Run(300)
	pb.RegisterGameServiceServer(server, gameService)

	fmt.Printf("[Service] listening at %v\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		fmt.Printf("[Service] failed to serve: %v\n", err)
	}
}

func StartGlobalService() {

}

func main() {
	go StartGameService()
	//Console.ReadLine()
	select {}
}
