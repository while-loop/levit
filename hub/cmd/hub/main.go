package main

import (
	"context"
	"flag"
	"log"

	"fmt"

	"math/rand"
	"time"

	"github.com/while-loop/levit/hub/proto"
	"google.golang.org/grpc"
)

var (
	raddr = flag.String("raddr", "localhost:8080", "remote address of hub server")
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())
	uid := rand.Uint64()

	conn, err := grpc.Dial(*raddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := proto.NewHubClient(conn)

	client, err := c.Connect(context.Background())
	if err != nil {
		log.Fatalf("unable to call hub connect: %v", err)
	}

	go func() {
		var line string
		for {
			fmt.Scanln(&line)
			if err := client.Send(&proto.HubMessage{
				Event: &proto.HubMessage_EventMessage{
					EventMessage: &proto.EventMessage{
						Message: line,
					},
				},
				Uid: uid,
			}); err != nil {
				log.Fatal("unable to send to server ", err)
			}
			line = ""
		}
	}()

	for {
		msg, err := client.Recv()
		if err != nil {
			log.Fatal("unable to recv from server ", err)
		}

		fmt.Println(msg)
	}
}
