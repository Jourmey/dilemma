package main

import (
	"dilemma/internal/config"
	"dilemma/internal/handler"
	"dilemma/internal/svc"
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/dilemma.json", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	g := service.NewServiceGroup()
	staticFile := NewStaticFile(c.Staticfile)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	//server.Use(rest.ToMiddleware(TracingHandler))
	handler.RegisterHandlers(server, ctx)

	g.Add(staticFile)
	g.Add(server)

	defer g.Stop()
	g.Start()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
}

type staticFile struct {
	config config.Staticfile
}

func NewStaticFile(config config.Staticfile) *staticFile {
	s := new(staticFile)
	s.config = config
	return s
}

func (s staticFile) Start() {
	m := http.NewServeMux()
	m.Handle("/", http.FileServer(http.Dir(s.config.Root)))
	_ = http.ListenAndServe(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port), m)
}

func (s staticFile) Stop() {}
