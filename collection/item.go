package collection

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Item ...
type Item struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	ImageURL string        `json:"image_url" bson:"image_url"`
	PageURL  string        `json:"page_url" bson:"page_url"`
	Text     string        `json:"text" bson:"text"`
	Created  time.Time     `json:"created" bson:"created"`
}
