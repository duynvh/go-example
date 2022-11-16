package appgrpc

import (
	"flag"
	"fmt"
	"food-delivery-service/common"
	"net"
	"time"

	"github.com/200Lab-Education/go-sdk/logger"
	"google.golang.org/grpc"
)

type grpcServer struct {
	prefix      string
	port        int
	server      *grpc.Server
	logger      logger.Logger
	registerHdl func(*grpc.Server)
}

func NewGRPCServer(prefix string) *grpcServer {
	return &grpcServer{prefix: prefix}
}

func (s *grpcServer) SetRegisterHdl(hdl func(*grpc.Server)) {
	s.registerHdl = hdl
}

func (s *grpcServer) GetPrefix() string {
	return s.prefix
}

func (s *grpcServer) Get() interface{} {
	return s
}

func (s *grpcServer) Name() string {
	return s.prefix
}

func (s *grpcServer) InitFlags() {
	flag.IntVar(&s.port, s.GetPrefix()+"-port", 50051, "Port of gRPC service")
}

func (s *grpcServer) Configure() error {
	s.logger = logger.GetCurrent().GetLogger(s.prefix)

	s.logger.Infoln("Setup gRPC service:", s.prefix)
	s.server = grpc.NewServer()

	return nil
}

func (s *grpcServer) Run() error {

	go func() {
		defer common.Recover()

		time.Sleep(time.Second * 3)

		_ = s.Configure()

		if s.registerHdl != nil {
			s.logger.Infoln("registering services...")
			s.registerHdl(s.server)
		}

		address := fmt.Sprintf("0.0.0.0:%d", s.port)
		lis, err := net.Listen("tcp", address)

		if err != nil {
			s.logger.Errorln("Error %v", err)
		}

		s.server.Serve(lis)
	}()

	return nil
}

func (s *grpcServer) Stop() <-chan bool {
	c := make(chan bool)

	go func() {
		s.server.Stop()
		c <- true
		s.logger.Infoln("Stopped")
	}()
	return c
}
