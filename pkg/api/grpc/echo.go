package grpc

import (
	"context"

	"go.uber.org/zap"

	"github.com/stefanprodan/podinfo/pkg/api/grpc/echo"
)

type echoServer struct {
	echo.UnimplementedEchoServiceServer
	config *Config
	logger *zap.Logger
}

func (s *echoServer) Echo(ctx context.Context, message *echo.Message) (*echo.Message, error) {
	s.logger.Info("Received message body from client", zap.String("Message", message.Body))
	return &echo.Message{Body: message.Body}, nil
}
