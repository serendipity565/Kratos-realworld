package server

import (
	v1 "Kratos-realworld/api/realworld/v1"
	"Kratos-realworld/internal/conf"
	"Kratos-realworld/internal/pkg/middleware/auth"
	"Kratos-realworld/internal/service"
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewSkipRoutersMatcher() selector.MatchFunc {

	skipRouters := make(map[string]struct{})
	skipRouters["/realworld.v1.RealWorld/Login"] = struct{}{}
	skipRouters["/realworld.v1.RealWorld/Register"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := skipRouters[operation]; ok {
			return false
		}
		return true
	}
}

func NewHTTPServer(c *conf.Server, jwtc *conf.JWT, greeter *service.RealWorldService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(auth.JWTAuth(jwtc.Token)).Match(NewSkipRoutersMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterRealWorldHTTPServer(srv, greeter)
	return srv
}
