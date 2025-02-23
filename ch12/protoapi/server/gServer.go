package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/testaquatic/Mastering_GO/ch12/protoapi/protoapi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	min  = 0
	max  = 100
	port = ":8080"
)

type RandomServer struct {
	protoapi.UnimplementedRandomServer
}

func (RandomServer) GetDate(ctx context.Context, _ *protoapi.RequestDateTime) (*protoapi.DateTime, error) {
	currentTime := time.Now()
	response := &protoapi.DateTime{
		Value: currentTime.String(),
	}

	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *protoapi.RandomParams) (*protoapi.RandomInt, error) {
	rng := rand.New(rand.NewSource(r.GetSeed()))
	place := r.GetPlace()

	temp := random(rng, min, max)
	for {
		place--
		if place <= 0 {
			break
		}
		temp = random(rng, min, max)
	}

	response := &protoapi.RandomInt{
		Value: int64(temp),
	}

	return response, nil
}

func random(rng *rand.Rand, min, max int) int {
	return rng.Intn(max-min) + min
}

func (RandomServer) GetRandomPass(_ context.Context, r *protoapi.RequestPass) (*protoapi.RandomPass, error) {
	rng := rand.New(rand.NewSource(r.GetSeed()))
	temp := getString(rng, r.GetLength())

	response := &protoapi.RandomPass{
		Password: temp,
	}

	return response, nil

}

func getString(rng *rand.Rand, len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(rng, 0, 94)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			return temp
		}
		i++
	}
}
func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Using default port:", port)
	} else {
		port = flag.Arg(0)
	}

	server := grpc.NewServer()

	var randomServer RandomServer
	protoapi.RegisterRandomServer(server, randomServer)
	reflection.Register(server)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Serving requests...")
	if err = server.Serve(listen); err != nil {
		fmt.Println(err)
	}
}
