// Copyright All rights reserved
/**
 * handler class
 */

package handler

import (
    "github.com/jonee/dialogflow_example/go/dialogflow_webhook/global"
    "github.com/jonee/dialogflow_example/go/dialogflow_webhook/model"

	apiai "github.com/mlabouardy/dialogflow-go-client/models"

    "gopkg.in/mgo.v2/bson"

	"encoding/json"
	"io/ioutil"
	"net/http"

    "errors"
    "log"
    "time"
)


// HandleWebhook function
//
// curl -X POST -d '{}' http://$server:$port/webhook
func HandleWebhook(w http.ResponseWriter, req *http.Request) error {

	var body []byte

	// grab json values from body
	if req.Body != nil {
		body, _ = ioutil.ReadAll(req.Body)
	}
	if req.Body == nil {
		global.R.JSON(w, http.StatusNotAcceptable, "parameters not provided")
		return errors.New("")
	}

	var t apiai.QueryResponse
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		global.R.JSON(w, http.StatusNotAcceptable, "parameters not provided: "+err.Error())
		return errors.New("")
	}

    stringBody := string(body)

log.Println("-- req --")
log.Println(stringBody)

    sessionId := t.SessionID

    action := t.Result.Action

	ret := make(map[string]interface{})

	reminderCol := global.MongoSession.DB(global.MongoDatabase).C("reminder")

    if action == "reminders.add" {
        var rem model.Reminder
        rem.SessionId = sessionId
        rem.Name = t.Result.Parameters["name"]
        rem.Recurrence, _ = t.Result.Parameters["recurrence"]

        // todo time like "05:00:00"
        stringDate := t.Result.Parameters["date-time"]

        rem.DateTime, _ = time.Parse(global.DIALOG_DATE_TIME_FORMAT, stringDate)

        rem.Location, _ = t.Result.Parameters["location"]
        rem.Data = stringBody
        rem.CreatedAt = time.Now()

        _, err = rem.Save(global.MapContext)
        if err != nil {
            log.Println(err)

	        ret["speech"] = "Sorry not able to save that :("
	        ret["displayText"] = "Sorry not able to save that :("

        } else {
	        ret["speech"] = t.Result.Fulfillment.Speech + " " + rem.Name
	        ret["displayText"] = t.Result.Fulfillment.Speech + " " + rem.Name
        }

    } else if action == "reminders.get" {
    	var rec model.Reminder

    	query := reminderCol.Find(bson.M{"session_id":sessionId, "date_time": bson.M{"$gt": time.Now()}}).Sort("date_time")
	    iter := query.Iter()
	    if iter.Err() == nil {
		    remindersText := ""
		    for iter.Next(&rec) {
                remindersText = remindersText + ", " + rec.Name + " " + rec.DateTime.Format(global.OUTPUT_DATE_TIME_FORMAT)
		    }

            if remindersText != "" {
	            ret["speech"] = remindersText[1:]
	            ret["displayText"] = remindersText[1:]

            } else {
	            ret["speech"] = "No current reminder set"
	            ret["displayText"] = "No current reminder set"
            }

	    } else {
	        ret["speech"] = "No current reminder set"
	        ret["displayText"] = "No current reminder set"
	    }

    } else if action == "reminders.remove" {
    	err = reminderCol.Remove(bson.M{"session_id": sessionId})

	    if err == nil {
	        ret["speech"] = t.Result.Fulfillment.Speech
	        ret["displayText"] = t.Result.Fulfillment.Speech

        } else {
	        ret["speech"] = "I am not sure if I was able to remove the reminders :("
	        ret["displayText"] = "I am not sure if I was able to remove the reminders :("
        }
    }

	global.R.JSON(w, http.StatusOK, ret)
	return nil

	// HandleWebhook
}
