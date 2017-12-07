package net

import (
	"github.com/urfave/negroni"
	"net/http"
)

// todo: get rid of negroni

// InitNegroni is a helper function for starting up http servers. it will create a new Negroni instance,
// use attach a logrus instance with a json formatter, and register the log instance with the appName
// it will also attach a handler, so really all you have to call is Run() on the returned instance
func InitNegroni(appName string, handler http.Handler) (n *negroni.Negroni){
	n = negroni.New(negroni.NewRecovery(), negroni.NewStatic(http.Dir("public")))
	n.UseHandler(handler)
	return
}