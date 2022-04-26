package repository

import (
	"bribrain-initial/model"
	"bribrain-initial/rpc/rka"
	"context"
)

type Services interface {
	Adduser(ctx context.Context, req *rka.User) (*model.Responses, error)
	AddRKA(ctx context.Context, req *rka.RKA) (*model.Responses, error)
	AddUserRKA(ctx context.Context, req *rka.UserRKA) (*model.Responses, error)
	AccumulationRKA(ctx context.Context, req *rka.RequestAccumulation) (*model.Responses, error)
}
