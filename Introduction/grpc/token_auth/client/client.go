package main

import (
	"context"
	"fmt"

	"GolangStudy/Introduction/grpc/token_auth/proto"

	"google.golang.org/grpc"
)

type customCredential struct{}

func (cc *customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (cc *customCredential) RequireTransportSecurity() bool {
	return false
}
func main() {
	// interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 	start := time.Now()
	// 	md := metadata.New(map[string]string{
	// 		"appid":  "101010",
	// 		"appkey": "i am key",
	// 	})
	// 	ctx = metadata.NewOutgoingContext(context.Background(), md)
	// 	err := invoker(ctx, method, req, reply, cc, opts...)
	// 	fmt.Printf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	// 	return err
	// }
	//stream
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	// 指定客户端interceptor
	opts = append(opts, grpc.WithPerRPCCredentials(&customCredential{}))

	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "bobby"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
