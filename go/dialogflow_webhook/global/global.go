// Copyright All rights reserved
/**
 * global class
 */

package global

import (
	"github.com/unrolled/render"

	"gopkg.in/mgo.v2"
)

var (
	R *render.Render
	MongoSession *mgo.Session
    MongoDatabase string
    MapContext map[string]interface{}
)

const DIALOG_DATE_TIME_FORMAT = "2006-01-02T15:04:05Z"
const OUTPUT_DATE_TIME_FORMAT = "Mon Jan 02 2006 15:04:05"

