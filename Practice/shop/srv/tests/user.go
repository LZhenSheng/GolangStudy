package main

import (
	"GolangStudy/Practice/shop/srv/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}
func TestGetUserList() {
	rsp, err := userClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    1,
		PSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for _, user := range rsp.Data {
		fmt.Println(user.Mobile, user.NickName, user.Password)
		checkRsp, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(checkRsp.Success)
	}
}
func TestCreateUser() {
	for i := 0; i < 10; i++ {
		rsp, err := userClient.CreateUser(context.Background(), &proto.CreateUserInfo{
			NickName: fmt.Sprintf("bobby%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			PassWord: "admin123",
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(rsp.Id)
	}
}
func TestCheckPassword() {
	rsp, err := userClient.CheckPassword(context.Background(), &proto.PasswordCheckInfo{
		Password:          "admin123",
		EncryptedPassword: "$pbkdf2-sha512$mzfWQ3Z51CPZIRwo$5f41d71e7a9c24d665f2ee271d1d49530e0ae58f30ad591f0d328e225a3621b0",
	})
	fmt.Println(err)
	fmt.Println(rsp)
}
func main() {
	Init()
	// TestCreateUser()
	// TestGetUserList()
	TestCheckPassword()
	conn.Close()
}
