//|------------------------------------------------------------------
//|        __
//|     __/  \
//|  __/  \__/_
//| /  \__/    \
//|/\__/CellGo /_
//|\/_/NetFW__/  \
//|  /\__ _/  \__/
//|  \/_/  \__/_/
//|    /\__/_/
//|    \/_/
//| ------------------------------------------------------------------
//| Cellgo Framework core type
//| All rights reserved: By cellgo.cn CopyRight
//| You are free to use the source code, but in the use of the process,
//| please keep the author information. Respect for the work of others
//| is respect for their own
//|-------------------------------------------------------------------
// Author:Tommy.Jin Dtime:2016-7-15

package cellgo

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	// CellCore is an core instance
	CellCore *Core
)

func init() {
	// create cellgo core
	CellCore = NewCore()
}

type Core struct {
	Handlers *ControllerRegister
	Server   *http.Server
}

func (core *Core) defalutHandler(w http.ResponseWriter, r *http.Request) {
	core.CheckParam(r)
	CellCore.Handlers.workHTTP(w, r)
	return
}

// NewCore returns a new cellgo core.
func NewCore() *Core {
	cr := NewControllerRegister()
	core := &Core{Handlers: cr, Server: &http.Server{}}
	return core
}

// Run cellgo core.
func (core *Core) Run() {
	core.RegisterHttp()
}

//register http service
func (core *Core) RegisterHttp() {

	for _, v := range CellConf.SiteConfig.StaticRouter {
		http.Handle(v, http.FileServer(http.Dir(CellConf.SiteConfig.StaticDir)))
	}
	http.HandleFunc(CellConf.SiteConfig.Dynamic, core.defalutHandler)
	server := http.Server{
		Addr:        CellConf.Listen.HTTPAddr + ":" + fmt.Sprintf("%d", CellConf.Listen.HTTPPort),
		Handler:     nil,
		ReadTimeout: time.Duration(CellConf.Listen.ServerTimeOut) * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

//Check default parameters
//From the configuration file options
func (core *Core) CheckParam(r *http.Request) {

	r.ParseForm() //Analytical parameters, the default is not resolved
	if r.Form["c"] == nil {
		r.Form["c"] = []string{strings.ToLower(CellConf.SiteConfig.DefaultController)} //default controller
	}
	if r.Form["a"] == nil {
		r.Form["a"] = []string{strings.ToLower(CellConf.SiteConfig.DefaultAction)} //default action
	}
}

//register https service
func (core *Core) RegisterHttps() {}

//register websocket service
func (core *Core) RegisterWebsocket() {}

func (core *Core) RegisterController(title string, c ControllerInterface, param []string) *Core {
	core.Handlers.Add(title, c, param)
	return core
}
