// Code generated by Kitex v0.11.0. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/lvkeliang/WHOIM-user-service/RPC/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, username string, password string, email string, callOptions ...callopt.Option) (r bool, err error)
	Login(ctx context.Context, username string, password string, callOptions ...callopt.Option) (r string, err error)
	ValidateToken(ctx context.Context, token string, callOptions ...callopt.Option) (r *user.User, err error)
	GetUserInfo(ctx context.Context, id string, callOptions ...callopt.Option) (r *user.User, err error)
	SetUserOnline(ctx context.Context, id string, deviceID string, serverAddress string, callOptions ...callopt.Option) (r bool, err error)
	SetUserOffline(ctx context.Context, id string, deviceID string, callOptions ...callopt.Option) (r bool, err error)
	GetUserDevices(ctx context.Context, id string, callOptions ...callopt.Option) (r map[string]*user.UserStatus, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, username string, password string, email string, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, username, password, email)
}

func (p *kUserServiceClient) Login(ctx context.Context, username string, password string, callOptions ...callopt.Option) (r string, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, username, password)
}

func (p *kUserServiceClient) ValidateToken(ctx context.Context, token string, callOptions ...callopt.Option) (r *user.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ValidateToken(ctx, token)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, id string, callOptions ...callopt.Option) (r *user.User, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, id)
}

func (p *kUserServiceClient) SetUserOnline(ctx context.Context, id string, deviceID string, serverAddress string, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetUserOnline(ctx, id, deviceID, serverAddress)
}

func (p *kUserServiceClient) SetUserOffline(ctx context.Context, id string, deviceID string, callOptions ...callopt.Option) (r bool, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SetUserOffline(ctx, id, deviceID)
}

func (p *kUserServiceClient) GetUserDevices(ctx context.Context, id string, callOptions ...callopt.Option) (r map[string]*user.UserStatus, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserDevices(ctx, id)
}
