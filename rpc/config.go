package rpc

import (
	"fmt"
	"log"
	"os"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var retryPolicy = `{
	"methodConfig": [{
		"waitForReady": true,
		"retryPolicy": {
			"MaxAttempts": 3,
			"InitialBackoff": ".01s",
			"MaxBackoff": ".01s",
			"BackoffMultiplier": 1.0,
			"RetryableStatusCodes": [ "UNAVAILABLE" ]
		}
	}]
}`

func NewRkaServiceRPC() (rpcConn *grpc.ClientConn, err error) {
	rpcAddress := fmt.Sprintf("%v:%v",
		os.Getenv("GRPC_HOST"),
		os.Getenv("GRPC_PORT"),
	)
	rpcConn, err = grpc.Dial(rpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))
	if err != nil {
		log.Fatalf("connection failure : %v", err)
	}
	return rpcConn, err
}
