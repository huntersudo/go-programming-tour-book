package main

import (
	"context"
	"flag"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	// 创建与给定目标（服务端）的连接句柄。
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()
    // 建 Greeter 的客户端对象。
	client := pb.NewGreeterClient(conn)
	// 发送 RPC 请求，等待同步响应，得到回调后返回响应结果。
	// 一元 RPC
	//_ = SayHello(client)

	// 服务端流式 RPC
	r:=&pb.HelloRequest{Name:"from client sml"}
	//_=SayList(client,r)

	// 客户端流式 RPC
	//_ =SayRecord(client,r)

	// 双向流式 RPC
	_=SayRoute(client,r)
}

// 一元 RPC
func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "eddycjy"})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

// 服务端流式 RPC
func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: %v", resp)
	}

	return nil
}
// 在 Client 端，主要留意 stream.Recv() 方法，我们可以思考一下，什么情况下会出现 io.EOF ，又在什么情况下会出现错误信息呢？
//  实际上 stream.Recv 方法，是对 ClientStream.RecvMsg 方法的封装，而 RecvMsg 方法会从流中读取完整的 gRPC 消息体，我们可得知：
//
//    RecvMsg 是阻塞等待的。
//    RecvMsg 当流成功/结束（调用了 Close）时，会返回 io.EOF。
//    RecvMsg 当流出现任何错误时，流会被中止，错误信息会包含 RPC 错误码。而在 RecvMsg 中可能出现如下错误，例如：
//        io.EOF、io.ErrUnexpectedEOF
//        transport.ConnectionError
//        google.golang.org/grpc/codes（gRPC 的预定义错误码）
// 需要注意的是，默认的 MaxReceiveMessageSize 值为 1024 *1024* 4，若有特别需求，可以适当调整。

// 客户端流式 RPC
// 在 Server 端的 stream.CloseAndRecv，与 Client 端 stream.SendAndClose 是配套使用的方法。
func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for n := 0; n < 6; n++ {
		_ = stream.Send(r)
	}
	resp, _ := stream.CloseAndRecv()

	log.Printf("resp err from server: %v", resp)
	return nil
}

// 双向流式 RPC
func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp err: %v", resp)
	}

	_ = stream.CloseSend()

	return nil
}