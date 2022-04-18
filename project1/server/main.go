package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	hello_grpc "xza/proto/test1/proto"
)

type server struct{
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server)SayHi(ctx context.Context,req *hello_grpc.Req)(res *hello_grpc.Res,err error){
	fmt.Println(req.GetMessage())
	return &hello_grpc.Res{Message: "从服务端返回的grpc内容"},nil
}
func (s *server) SayHello(ctx context.Context,req *hello_grpc.Req2) (res *hello_grpc.Res2,err error){
	fmt.Println(req.GetMessage())
	res = &hello_grpc.Res2{
		Message: "Roger that",
		Name: "random name",
	}
	if req.GetAge()==25{
		res.Age = 25
		fmt.Println("25 return 25")
	}else{
		res.Age = 0
		fmt.Println("not 25 return 0 ")
	}
	return res,nil
}
func main(){
	//addr:= &net.TCPAddr{
	//	IP: []byte(""),
	//	Port: 8888,
	//}
	//l,err:=net.ListenTCP("tcp4",addr)
	l,err:=net.Listen("tcp","0.0.0.0:8888")
	if err!=nil{
		fmt.Println(err)
	}
	s:=grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s,&server{})
	s.Serve(l)
}
