package grpc

import (
	"context"
	"runtime"
	"strconv"

	"github.com/stefanprodan/podinfo/pkg/api/grpc/info"
	"github.com/stefanprodan/podinfo/pkg/version"
)

type infoServer struct {
	info.UnimplementedInfoServiceServer
	config *Config
}

func (s *infoServer) Info (ctx context.Context, message *info.InfoRequest) (*info.InfoResponse, error){

	return &info.InfoResponse {
		Hostname:     s.config.Hostname,
		Version:      version.VERSION,
		Revision:     version.REVISION,
		Color:        s.config.UIColor,
		Logo:         s.config.UILogo,
		Message:      s.config.UIMessage,
		Goos: 		  runtime.GOOS,
		Goarch: 	  runtime.GOARCH,
		Runtime: 	  runtime.Version(),
		Numgoroutine: strconv.FormatInt(int64(runtime.NumGoroutine()), 10),
		Numcpu: 	  strconv.FormatInt(int64(runtime.NumCPU()), 10),

	}, nil

}


