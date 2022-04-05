package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hollson/ddd_auth2/adpter/http/auth_handles"
	"github.com/hollson/ddd_auth2/adpter/http/routers"
	"github.com/hollson/ddd_auth2/domain/service"
	"github.com/hollson/ddd_auth2/infrastructure/conf"
	"github.com/hollson/ddd_auth2/infrastructure/pkg/log"
	"go.uber.org/zap"
)

func NewHttp(s *conf.AppConfig, auth service.AuthSrv) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	h := auth_handles.NewHandles(s, auth)
	routers.SetRouters(g, h)
	server := &http.Server{
		Addr:           s.NetConf.ServerAddr,
		Handler:        g,
		ReadTimeout:    time.Duration(s.NetConf.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(s.NetConf.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.GetLogger().Info("auth server start success", zap.Any("addr", s.NetConf.ServerAddr))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}
