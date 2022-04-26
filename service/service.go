package service

import (
	"bribrain-initial/model"
	"bribrain-initial/repository"
	"bribrain-initial/rpc"
	"bribrain-initial/rpc/rka"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type RkaService struct {
}

func NewServiceRka() repository.Services {
	return RkaService{}
}

func (RkaService) Adduser(ctx context.Context, req *rka.User) (*model.Responses, error) {
	rpcConn, err := rpc.NewRkaServiceRPC()
	if err != nil {
		logrus.Error(fmt.Sprintf("cannot connect to server Rpc : %v", err))
		return nil, err
	}

	defer rpcConn.Close()

	rpcServiceRKA := rka.NewRkaServiceClient(rpcConn)
	res, err := rpcServiceRKA.Adduser(ctx, req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if res.Code != 200 {
		logrus.Warning("call rpc func failed : " + res.Message)
		return &model.Responses{
			Code:    int(res.Code),
			Message: res.Message,
		}, nil
	}

	return &model.Responses{
		Code:    int(res.Code),
		Message: res.Message,
		Data:    res.Data,
	}, nil
}

func (RkaService) AddRKA(ctx context.Context, req *rka.RKA) (*model.Responses, error) {
	rpcConn, err := rpc.NewRkaServiceRPC()
	if err != nil {
		logrus.Error(fmt.Sprintf("cannot connect to server Rpc : %v", err))
		return nil, err
	}

	defer rpcConn.Close()

	rpcServiceRKA := rka.NewRkaServiceClient(rpcConn)
	res, err := rpcServiceRKA.AddRKA(ctx, req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if res.Code != 200 {
		logrus.Warning("call rpc func failed : " + res.Message)
		return &model.Responses{
			Code:    int(res.Code),
			Message: res.Message,
		}, nil
	}

	return &model.Responses{
		Code:    int(res.Code),
		Message: res.Message,
		Data:    res.Data,
	}, nil
}

func (RkaService) AddUserRKA(ctx context.Context, req *rka.RequestUserRKA) (*model.Responses, error) {
	rpcConn, err := rpc.NewRkaServiceRPC()
	if err != nil {
		logrus.Error(fmt.Sprintf("cannot connect to server Rpc : %v", err))
		return nil, err
	}

	defer rpcConn.Close()

	rpcServiceRKA := rka.NewRkaServiceClient(rpcConn)
	res, err := rpcServiceRKA.AddUserRKA(ctx, req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if res.Code != 200 {
		logrus.Warning("call rpc func failed : " + res.Message)
		return &model.Responses{
			Code:    int(res.Code),
			Message: res.Message,
		}, nil
	}

	return &model.Responses{
		Code:    int(res.Code),
		Message: res.Message,
		Data:    res.Data,
	}, nil
}

func (RkaService) AccumulationRKA(ctx context.Context, req *rka.RequestAccumulation) (*model.Responses, error) {
	rpcConn, err := rpc.NewRkaServiceRPC()
	if err != nil {
		logrus.Error(fmt.Sprintf("cannot connect to server Rpc : %v", err))
		return nil, err
	}

	defer rpcConn.Close()

	rpcServiceRKA := rka.NewRkaServiceClient(rpcConn)
	res, err := rpcServiceRKA.AccumulationRKA(ctx, req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	if res.Code != 200 {
		logrus.Warning("call rpc func failed : " + res.Message)
		return &model.Responses{
			Code:    int(res.Code),
			Message: res.Message,
		}, nil
	}

	return &model.Responses{
		Code:    int(res.Code),
		Message: res.Message,
		Data:    res.ProgressVisited,
	}, nil
}
