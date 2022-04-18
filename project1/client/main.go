package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc "xza/proto/test1/proto"
)

func test1(){
	conn,err:=grpc.Dial("localhost:8888",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err)
		return
	}
	client:=hello_grpc.NewHelloGRPCClient(conn)
	req,_:=client.SayHi(context.Background(),&hello_grpc.Req{Message: "我从客户端来"})
	fmt.Println("123")
	fmt.Println(req.GetMessage())
}
func test2(){
	conn,err:=grpc.Dial("localhost:8888",grpc.WithInsecure())
	if err!=nil{
		fmt.Println(err)
		return
	}
	client:=hello_grpc.NewHelloGRPCClient(conn)
	req,_:=client.SayHi(context.Background(),&hello_grpc.Req{Message: "我从客户端来"})
	req2,_:=client.SayHello(context.Background(),&hello_grpc.Req2{Message: "我从客户端来",Name: "good",Age: 20})
	fmt.Println(req.GetMessage())
	fmt.Println(req2.GetMessage(),req2.GetName(),req2.GetAge())
}
func main(){
	test2()
	//grpc.WithTransportCredentials()
	//conn,err:=grpc.Dial("localhost:8888",grpc.WithInsecure())

}

