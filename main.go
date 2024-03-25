package main

import (
	"Server/game"
	pb "Server/proto"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
	"time"
)

//type dsnConfig struct {
//	Dsn string `json:"dsn"`
//}

//func getMySqlDB() *gorm.DB {
//	var config dsnConfig    // MySQL配置
//	var conn gorm.Dialector // gorm连接
//	file, _ := os.Open("config.json")
//	defer func(file *os.File) {
//		err := file.Close()
//		if err != nil {
//			fmt.Printf("[Service] failed to close file: %v\n", err)
//		}
//	}(file)
//
//	// 解析json
//	decoder := json.NewDecoder(file)
//	_ = decoder.Decode(&config)
//	fmt.Printf("[Service] dsn: %s\n", config.Dsn)
//	conn = mysql.Open(config.Dsn) // mysql
//	db, err := gorm.Open(conn, &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	fmt.Printf("[Service] connected to database.\n")
//	return db
//}
//
//func getSqliteDB() *gorm.DB {
//	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//	return db
//}

func getMongoDB() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	return client.Database("course2024")
}

func StartGameService() {
	fmt.Printf("[ Service] init.\n")
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
	db := getMongoDB()
	gameService := game.NewGameService(db)
	gameService.Run()
	pb.RegisterGameServiceServer(server, gameService)

	fmt.Printf("[Service] listening at %v\n", lis.Addr())
	if err := server.Serve(lis); err != nil {
		fmt.Printf("[Service] failed to serve: %v\n", err)
	}

	// 关闭数据库连接
	err = db.Client().Disconnect(context.TODO())
	if err != nil {
		fmt.Printf("[Service] failed to disconnect: %v\n", err)
	}
}

func StartGlobalService() {

}

func main() {
	go StartGameService()
	//Console.ReadLine()
	select {}
}
