package server

import "git.u-linke.com/ulink/commons/library/worm"

// IServer 包含Service ，Service包含Component ,Component包括Module
type IServer interface {
	Init()
	Run()
	Exit()
	Ticker()
	LoadService()
}

// Server struct
type Server struct {
	etcdClient *etcd.Client
	wormClient *worm.Client
	mqClient   *mq.Client
}

func NewServer() *Server {

}

func (srv *Server) Init() {
	//TODO:服务器初始化
}

func (srv *Server) Run() {
	//TODO::服务器运行
}

func (srv *Server) Exit() {
	//TODO：服务器退出
}

func (srv *Server) LoadService() {
	//todo::加载服务
}

func (srv *Server) Ticker() {
	//todo:: Ticker 处理
}
