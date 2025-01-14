package middleware

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var IsProdMod = false

func GrpcInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = status.Errorf(codes.Internal, "panic: %v", r)
			logrus.Warnf("gRPC Handler Err:%v", err)
		}
	}()

	if !IsProdMod {
		reqJson, _ := json.Marshal(req)
		if len(reqJson) < 512 {
			fmt.Printf("%s Start...\n=====> Req：%s \n", info.FullMethod, reqJson)
		}
	}

	// if !strings.HasSuffix(info.FullMethod, "Authorization/Check") {
	ctx, err = validHead(ctx)
	if err != nil {
		return
	}

	resp, err = handler(ctx, req)
	// logs.ConsoleLogs.Info("%s End...Err:%v \n", info.FullMethod, err)
	return
}

func validHead(ctx context.Context) (newCtx context.Context, err error) {
	// md 的值类似于: map[:authority:[192.168.40.123:50051] authorization:[Bearer some-secret-token] content-type:[application/grpc] user-agent:[grpc-go/1.20.1]]
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		err = status.Errorf(codes.InvalidArgument, "missing metadata")
		return
	}

	reqUuid, ok := md["requuid"]
	if !ok {
		err = status.Errorf(codes.Unauthenticated, `missing "requuid" header`)
		return
	}
	if len(reqUuid) == 0 || reqUuid[0] == "" {
		err = status.Errorf(codes.Unauthenticated, "missing requuid")
		return
	}
	newCtx = context.WithValue(ctx, "requuid", reqUuid[0])
	return
}

// func validJWT(ctx context.Context) (err error) {

// 	// md 的值类似于: map[:authority:[192.168.40.123:50051] authorization:[Bearer some-secret-token] content-type:[application/grpc] user-agent:[grpc-go/1.20.1]]
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return status.Errorf(codes.InvalidArgument, "missing metadata")
// 	}

// 	// 1. 判断是否存在 authorization 请求头
// 	authorization, ok := md["authorization"]
// 	if !ok {
// 		return status.Errorf(codes.Unauthenticated, `missing "Authorization" header`)
// 	}

// 	// 2. 如果存在 authorization 请求头的话，则 md["authorization"] 是一个 []string
// 	// if !strings.HasPrefix(authorization[0], Prefix) {
// 	// 	return status.Errorf(codes.Unauthenticated, `missing "Bearer " prefix in "Authorization" header`)
// 	// }

// 	_, err = ValidJWT(authorization[0])
// 	if err != nil {
// 		return status.Errorf(codes.Unauthenticated, err.Error())
// 	}
// 	return
// }
