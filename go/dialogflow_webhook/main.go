// Copyright All rights reserved
/**
 * main class
 */

package main

import (
    "github.com/jonee/dialogflow_example/go/dialogflow_webhook/app"
    "github.com/jonee/dialogflow_example/go/dialogflow_webhook/global"
    "github.com/jonee/dialogflow_example/go/dialogflow_webhook/model"

    "github.com/unrolled/render"
    // "github.com/urfave/negroni"

	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"

    "net/http"
    // "fmt"
    "log"
    "time"
)

func main() {
	var err error

	t0 := time.Now()
	_, offset := t0.Zone()
	log.Println("offset:", offset)

	global.MapContext = make(map[string]interface{})

	var mongoDialInfo *mgo.DialInfo

	global.MongoDatabase = "dialog"
	global.MapContext["MongoDatabase"] = global.MongoDatabase

	mongoDialInfo = &mgo.DialInfo{
		Addrs: []string{"localhost:27017"},
		// Timeout:  60 * time.Second,
		Database: "admin",
		Username: "admin",
		Password: "password",
	}

	global.MongoSession, err = mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer global.MongoSession.Close()
	global.MapContext["MongoSession"] = global.MongoSession

	// global.R = render.New()
	global.R = render.New(render.Options{StreamingJSON: true}) // try json streaming by default
	global.MapContext["R"] = global.R


	model.ReminderEnsureIndex(global.MapContext)

	n := app.NewServer()

    err = http.ListenAndServeTLS(":8453", "/etc/letsencrypt/live/talktothebot.greatwwwhost.com/fullchain.pem", "/etc/letsencrypt/live/talktothebot.greatwwwhost.com/privkey.pem", n)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
