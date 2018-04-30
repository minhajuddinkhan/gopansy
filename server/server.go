package server

import (
	"net/http"
	"time"

	conf "github.com/minhajuddinkhan/gopansy/config"

	"github.com/urfave/negroni"
)

var configuration conf.Configuration

//Server Go Pansy Server
type Server struct {
	Conf   conf.Configuration
	Router *negroni.Negroni
}

//Start Start the server
func (h *Server) Start() error {

	svr := http.Server{
		Addr:         h.Conf.Addr,
		Handler:      h.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return svr.ListenAndServe()
}
