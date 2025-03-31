package main

import (
	"github.com/ITu-CloudWeGo/itu_rpc_user/config"
	user_service "github.com/ITu-CloudWeGo/itu_rpc_user/kitex_gen/user_service/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	etcdRegistry "github.com/kitex-contrib/registry-etcd"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"time"
)

func main() {
	opts := kitexInit()

	svr := user_service.NewServer(new(UserServiceImpl), opts...)
	err := svr.Run()
	if err != nil {
		klog.Error("Failed to run server:", err.Error())
		panic(err)
	}

}

func kitexInit() (opts []server.Option) {
	conf := config.GetConfig()
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.Kitex.Address)
	if err != nil {
		klog.Error("Failed to resolve TCP address:", err.Error())
		panic(err)
	}
	if addr == nil {
		klog.Error("Invalid TCP address")
		panic("Invalid TCP address")
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.Kitex.Service,
	}))
	// registry_etcd
	r, err := etcdRegistry.NewEtcdRegistry(conf.Registry.RegistryAddress)
	if err != nil {
		klog.Error("Failed to create ETCD registry:", err.Error())
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(config.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.Kitex.LogFileName,
			MaxSize:    conf.Kitex.LogMaxSize,
			MaxBackups: conf.Kitex.LogMaxBackups,
			MaxAge:     conf.Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		if err := asyncWriter.Sync(); err != nil {
			klog.Error("Failed to sync log writer:", err.Error())
		}
	})

	return
}
