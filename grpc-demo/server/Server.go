package main

import (
	"context"
	"flag"
	pb "github.com/go-programming-tour-book/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)
var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}
func main() {
	server := grpc.NewServer()
	// 将 GreeterServer（其包含需要被调用的服务端接口）注册到 gRPC Server。 的内部注册中心。这样可以在接受到请求时，通过内部的 “服务发现”，发现该服务端接口并转接进行逻辑处理。
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}


// 创建 gRPC Server 对象，你可以理解为它是 Server 端的抽象对象。
type GreeterServer struct{}

// 一元 RPC
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello.world"}, nil
}

// 服务端流式 RPC
func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for n := 0; n <= 6; n++ {
		_ = stream.Send(&pb.HelloReply{Message: "from server -- hello.list"+r.Name})
	}
	return nil
}
// 在 Server 端，主要留意 stream.Send 方法，通过阅读源码，可得知是 protoc 在生成时，根据定义生成了各式各样符合标准的接口方法。最终再统一调度内部的 SendMsg 方法，该方法涉及以下过程:
//
//   消息体（对象）序列化。
//   压缩序列化后的消息体。
//   对正在传输的消息体增加 5 个字节的 header（标志位）。
//   判断压缩 + 序列化后的消息体总字节长度是否大于预设的 maxSendMessageSize（预设值为 math.MaxInt32），若超出则提示错误。
//    写入给流的数据集


// 客户端流式 RPC
func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message:"say.record.Done"})
		}
		if err != nil {
			return err
		}

		log.Printf("resp from client: %v", resp)
	}

	return nil
}
// 你可以发现在这段程序中，我们对每一个 Recv 都进行了处理，
// 当发现 io.EOF (流关闭) 后，需要通过 stream.SendAndClose 方法将最终的响应结果发送给客户端，同时关闭正在另外一侧等待的 Recv。


// 双向流式 RPC
//由客户端以流式的方式发起请求，服务端同样以流式的方式响应请求。
func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "from server: say.route"})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Printf("resp: %v", resp)
	}
}