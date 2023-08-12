package todo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title" bson:"title" binding:"required"`
	ActiveAt string             `json:"activeAt" bson:"activeAt"`
	Status   string             `json:"status" bson:"status"`
}

type TodoRes struct {
	Title    string `json:"title" bson:"title" binding:"required"`
	ActiveAt string `json:"activeAt" bson:"activeAt"`
}

type TodoRequset struct {
	Title    string `json:"title" binding:"required"`
	ActiveAt string `json:"activeAt" binding:"required"`
}

type TodoURI struct {
	Id string `uri:"id" binding:"required"`
}

type TodoStatus struct {
	Status string `uri:"status" binding:"required"`
}
