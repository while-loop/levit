package main

import (
	"github.com/while-loop/levit/users/proto"
	"google.golang.org/grpc"
	"github.com/while-loop/levit/liblevit/log"
	"context"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c:= users.NewUsersClient(conn)

	resp, err :=c.Get(context.TODO(),&users.GetRequest{
		Ids:[]uint64{1},
	})

	log.Infof("%v", resp)
}
