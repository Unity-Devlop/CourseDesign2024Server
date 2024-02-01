package game

import (
	pb "Server/proto"
	"context"
	"google.golang.org/grpc"
	"testing"
)

func TestServer_BubbleChat(t *testing.T) {
	conn, err := grpc.Dial("localhost:22333", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			t.Fatalf("failed to close conn: %v", err)
		}
	}(conn)
	c := pb.NewGameServiceClient(conn)
	response, err := c.BubbleChat(context.Background(), &pb.ChatMessage{Uid: 1, Msg: "hello"})
	if err != nil {
		t.Fatalf("failed to call BubbleChat: %v", err)
	}
	if response.Code != pb.StatusCode_OK {
		t.Fatalf("failed to call BubbleChat: %v", response.Code)
	}
}
