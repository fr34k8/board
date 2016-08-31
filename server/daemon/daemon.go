package daemon

import (
	"os"
	"errors"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/julienschmidt/httprouter"
	"github.com/op/go-logging"
	"github.com/mtti/board/server/context"
	"github.com/mtti/board/server/bindata"
)

type Daemon struct {
	Settings *Settings
	Router *httprouter.Router
	Context *context.Global
}

func New() (*Daemon, error) {

	server := &Daemon{
		Settings: NewSettings(),
	}

	// Logging

	log := logging.MustGetLogger("mtti-board")
	logBackend := logging.AddModuleLevel(logging.NewBackendFormatter(
		logging.NewLogBackend(os.Stderr, "", 0),
		logging.MustStringFormatter("%{color}%{time:2006-01-02T15:04:05.999Z-07:00} %{shortfile} %{level}%{color:reset} %{message}"),
	))
	logBackend.SetLevel(logging.DEBUG, "")
	log.SetBackend(logBackend)

	// Global context

	server.Context = &context.Global{
		Log: log,
	}

	// Routes

	server.Router = httprouter.New()

	server.Router.GET("/", server.ServeEmbeddedAsset("index.html"))
	server.Router.GET("/main.js", server.ServeEmbeddedAsset("main.js"))

	/*
	// Create a new document
	server.Router.POST("/documents/", server.Context.WrapRouteHandlers(
		controllers.HandleDocumentCreate))
	
	// Retrieve latest document version
	server.Router.GET("/documents/:id", server.Context.WrapRouteHandlers(
		controllers.HandleDocumentRetrieve))
	
	// Retrieve list of document versions
	server.Router.GET("/documents/:id/versions/",
		server.Context.WrapRouteHandlers(controllers.HandleDocumentRetrieveVersions))
	
	// Retrieve specific document version
	server.Router.GET("/documents/:id/versions/:version",
		server.Context.WrapRouteHandlers(controllers.HandleDocumentRetrieve))
	*/

	/*
	server.router.PUT("/sites/:site/documents/:id", server.Context.WrapRouteHandlers(
		controllers.HandleDocumentUpdate))
	server.router.POST("/login", server.Context.WrapRouteHandlers(
		controllers.HandleSessionLogin))
	*/
		
	return server, nil

}

func (server *Daemon) Start() error {
	if server.Context.Repository == nil {
		return errors.New("No repository has been configured")
	}
	server.Context.Log.Info("Starting to listen on " + server.Settings.ListenAddress)
	return http.ListenAndServe(server.Settings.ListenAddress, server.Router)
}

func (server *Daemon) ServeStaticFile(name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.ServeFile(w, r, server.Settings.StaticDirectory + "/" + name)
	}
}

func (server *Daemon) ServeEmbeddedAsset(name string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data, err := bindata.Asset(name)
		if err != nil {
			return
		}
		w.Write(data)
	}	
}

func (server *Daemon) Redirect(destination string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		http.Redirect(w, r, destination, http.StatusMovedPermanently)
	}	
}