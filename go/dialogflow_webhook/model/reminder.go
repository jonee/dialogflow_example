// Copyright All rights reserved
/**
 * reminder model
 */

package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

    "log"
    "time"
)

var (
)

// Reminder class or struct definition
type Reminder struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`

    SessionId   string          `json:"session_id,omitempty" bson:"session_id,omitempty"`

    Name   string          `json:"name,omitempty" bson:"name,omitempty"`
    Recurrence   string          `json:"recurrence,omitempty" bson:"recurrence,omitempty"`
	DateTime time.Time `json:"date_time,omitempty" bson:"date_time,omitempty"`
    Location   string          `json:"location,omitempty" bson:"location,omitempty"`

    Data   string          `json:"data,omitempty" bson:"data,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

// Class init function
func init() {
	log.Println("Model - Reminder - init")
}

// Reminder class save function
func (u *Reminder) Save(mapContext map[string]interface{}) (*Reminder, error) {
	// ReminderEnsureIndex(mapContext)

	session, _ := mapContext["MongoSession"].(*mgo.Session)
	db, _ := mapContext["MongoDatabase"].(string)

	reminderCol := session.DB(db).C("reminder")

	var err error
	if !u.Id.Valid() {
		u.Id = bson.NewObjectId()
		_, err = reminderCol.UpsertId(u.Id, u)
	} else {
		_, err = reminderCol.UpsertId(u.Id, u)
	}

	return u, err

	// save
}

// ensureIndex function
func ReminderEnsureIndex(mapContext map[string]interface{}) {
	session, _ := mapContext["MongoSession"].(*mgo.Session)
	db, _ := mapContext["MongoDatabase"].(string)

	reminderCol := session.DB(db).C("reminder")

    // index for searching
	keyIndex := mgo.Index{
		Key:        []string{"session_id"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	err := reminderCol.EnsureIndex(keyIndex)
	if err != nil {
		log.Println("ensure key index {session_id} failed for Reminder: " + err.Error())
	}

	// ReminderEnsureIndex
}

/*
func (r *Reminder) ExportArray(mapContext map[string]interface{}) map[string]interface{} {
	// session, _ := ctx.GetSession()
	// config := ctx.GetConfig()

	recArr := global.ExportArray(r)

	recArr["_id"] = recArr["id"]

	return recArr

	// ExportArray
}
*/
