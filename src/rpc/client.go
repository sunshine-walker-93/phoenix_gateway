package rpc

import (
	"context"
	"time"

	pb "github.com/sunshine-walker-93/phoenix_apis/protobuf3.pb/user_info_manage"
	"github.com/sunshine-walker-93/phoenix_gateway/src/config"
	"github.com/sunshine-walker-93/phoenix_gateway/src/log"
	"github.com/sunshine-walker-93/phoenix_gateway/src/util"

	"github.com/micro/plugins/v5/registry/etcd"
	goMicro "go-micro.dev/v5"
	"go-micro.dev/v5/registry"
)

var GrpcClient pb.UserService

func init() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("etcd:2379"),
	)
	// 初始化服务
	srv := goMicro.NewService(
		goMicro.Registry(etcdReg),
	)

	GrpcClient = pb.NewUserService("phoenix_account", srv.Client())
}

// Register 注册接口
func Register(requestID string, name string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	_, err := GrpcClient.Register(ctx, &pb.RegisterRequest{RequestID: requestID, Name: name, Password: password})
	if err != nil {
		log.Warnf("call Register failed, err:%v", err)
		return err
	}

	return nil
}

// Auth 认证接口
func Auth(requestID string, name string, password string) (string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.Auth(ctx, &pb.AuthRequest{RequestID: requestID, Name: name, Password: password})
	if err != nil {
		log.Warnf("call Auth failed, err:%v", err)
		return "", "", err
	}

	return rsp.Nickname, rsp.Image, nil
}

// GetProfile 查询用户的属性信息
func GetProfile(requestID string, name string) (info *util.ProfileInfo, err error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.GetProfile(ctx, &pb.GetProfileRequest{RequestID: requestID, Name: name})
	if err != nil {
		log.Warnf("call GetProfile failed, err:%v", err)
		return nil, err
	}

	return &util.ProfileInfo{Nickname: rsp.Nickname, ImageID: rsp.ImageID}, nil
}

// GetHeadImage 查询用户的头像信息
func GetHeadImage(requestID string, imageID string) (image []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	rsp, err := GrpcClient.GetHeadImage(ctx, &pb.GetHeadImageRequest{RequestID: requestID, ImageID: imageID})
	if err != nil {
		log.Warnf("call GetHeadImage failed, err:%v", err)
		return nil, err
	}

	return rsp.Image, nil
}

// EditProfile 编辑用户的属性信息
func EditProfile(requestID string, name string, nickname string, image []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(config.GetGlobalConfig().AppSetting.DeadlineSecond)*time.Second)
	defer cancel()

	_, err := GrpcClient.EditProfile(ctx, &pb.EditProfileRequest{RequestID: requestID,
		Name: name, Nickname: nickname, Image: image})
	if err != nil {
		log.Warnf("call EditProfile failed, err:%v", err)
		return err
	}

	return nil
}
