package adpter

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hollson/ddd_auth2/adpter/http"
	"github.com/hollson/ddd_auth2/domain/service"
	"github.com/hollson/ddd_auth2/infrastructure/conf"
	"github.com/hollson/ddd_auth2/infrastructure/pkg/log"
	"go.uber.org/zap"
)

// QuitSignal
// syscall.SIGQUIT 用户发送QUIT字符(Ctrl+/)触发
// syscall.SIGTERM 结束程序(可以被捕获、阻塞或忽略)
// syscall.SIGINT 用户发送INTR字符(Ctrl+C)触发
func QuitSignal(quitFunc func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	fmt.Printf("server start success pid:%d\n", os.Getpid())
	for s := range c {
		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			quitFunc()
			return
		default:
			return
		}
	}
}

type Server struct {
	conf *conf.AppConfig
	log  *log.Logger
	auth service.AuthSrv
}

func NewSrv(c *conf.AppConfig, log *log.Logger, auth service.AuthSrv) *Server {
	s := &Server{conf: c, log: log, auth: auth}
	s.Init()
	return s
}
func (s *Server) Init() {
	// hcode.Click()
}

func (s *Server) RunApp() {
	http.NewHttp(s.conf, s.auth)
	QuitSignal(func() {
		s.Close()
		log.GetLogger().Info("auth server exit", zap.Any("addr", s.conf.NetConf.ServerAddr))
	})
}

func (s *Server) Close() {

}
