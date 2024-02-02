package main

import (
	pb "Server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"math/rand"
)

func createClient() (*grpc.ClientConn, pb.GameServiceClient) {
	conn, err := grpc.Dial("localhost:22333",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	c := pb.NewGameServiceClient(conn)
	return conn, c
}

func main() {
	conn, client := createClient()
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	var randomUid = rand.Uint32()
	ctx := context.Background()
	// 一直测试 直到退出
	endChan := make(chan int)
	// 控制台输入 任务
	go func() {
		defer func() {
			endChan <- 1
		}()
		for {
			var input string
			_, err := fmt.Scanln(&input)
			if err != nil {
				panic(err)
			}
			msg := &pb.ChatMessage{Uid: randomUid, Msg: input}
			//response, err := client.BubbleChat(ctx, msg)
			response, err := client.PublicChat(ctx, msg)
			if err != nil {
				fmt.Printf("failed to call BubbleChat: %v\n", err)
			}
			if response.Code != pb.StatusCode_OK {
				fmt.Printf("failed to call BubbleChat: %v\n", response.Code)
			}

		}
	}()

	// 从服务器接收消息
	go func() {
		defer func() {
			endChan <- 1
		}()
		var chatRequest = &pb.ChatRequest{
			Uid: randomUid,
		}
		//chatStream, err := client.StartBubbleChat(ctx, chatRequest)
		chatStream, err := client.StartPublicChat(ctx, chatRequest)
		if err != nil {
			fmt.Printf("failed to call StartBubbleChat: %v\n", err)
			return
		}
		for {
			chatMessage, err := chatStream.Recv()
			if err == io.EOF {
				fmt.Printf("chat stream closed\n")
				return
			}
			if err != nil {
				fmt.Printf("failed to receive chat message: %v\n", err)
				return
			}
			fmt.Printf("receive chat message: %v\n", chatMessage)
		}

	}()

	// 等待goroutine结束
	select {
	case <-endChan:
		fmt.Printf("end\n")
	}
}
