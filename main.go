package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var serverName string = "localhost"
var port int32 = 9090
var help string = "help"

func main() {
	// set default target
	target := fmt.Sprintf("dns:///%v:%v", serverName, port)
	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	err := handle(os.Args[0], target)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}

func handle(name, target string) error {
	if target == help {
		fmt.Printf("usage: %s [target]\n\n", name)
		fmt.Println("list grpc server services for a given target (default `localhost:9090`)")
		os.Exit(0)
	}

	insec_opt := grpc.WithTransportCredentials(insecure.NewCredentials())
	con, err := grpc.NewClient(target, insec_opt)
	if err != nil {
		return fmt.Errorf("failed to create client with target `%v`: %w", target, err)

	}

	refClient := grpcreflect.NewClientAuto(context.Background(), con)
	services, err := refClient.ListServices()
	if err != nil {
		return fmt.Errorf("failed to list services for server `%v`: %w", target, err)
	}

	fmt.Println("|------------|")
	fmt.Println("|- Services -|")
	fmt.Println("|------------|")
	for _, s := range services {
		fmt.Println(s)
	}
	return nil
}
