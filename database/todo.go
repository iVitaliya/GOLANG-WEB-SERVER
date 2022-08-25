package database

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ToDoModel struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string        `bson:"title"`
	Completed bool          `bson:"completed"`
	CreatedAt time.Time     `bson:"createdAt"`
}

type ToDo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
