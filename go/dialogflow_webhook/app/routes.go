// Copyright All rights reserved
/**
 * routes file
 */

package app

import (
    "github.com/jonee/dialogflow_example/go/dialogflow_webhook/handler"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	restiful "github.com/laicosly/restiful/v1"
    "github.com/urfave/negroni"
)

var (
    Router *mux.Router
)

// setups the server routes etc
func NewServer() *negroni.Negroni {
	Router := mux.NewRouter()

	// webhook routes
	route_webhook := Router.PathPrefix("/webhook").Subrouter()
	{
		route_webhook.Handle("/", restiful.Handle(handler.HandleWebhook)).Name("webhook").Methods("POST")
		route_webhook.Handle("/", restiful.Handle(handler.HandleWebhook)).Name("webhook").Methods("GET")
	}

/*
	Router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/index.html", http.StatusMovedPermanently) // 301
	}).Methods("GET")
*/

	n := negroni.Classic()
	// n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(http.Dir(currentPath+"/../www/build/")))
	// n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	// n := negroni.New(negroni.NewRecovery(), &negroni.Logger{Log})

	// Router goes last
	n.UseHandler(context.ClearHandler(Router))

	return n

	// NewServer
}


